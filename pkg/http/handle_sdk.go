package http

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/sdk"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) handleSDKRiskyCheck(c *gin.Context) {
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
		"id":      "",
		"action":  "ACTION_NONE",
		"geetest": nil,
	}))
}

func (s *Server) handleSDKShieldLogin(c *gin.Context) {
	var req sdk.ShieldLoginRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	if req.Account == "" {
		log.Error().Err(fmt.Errorf("account is empty")).Msg("Bad request")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	account, err := s.serviceShieldLogin(c, req.Account, req.Password, req.IsCrypto)
	if err == sql.ErrNoRows && s.config.AutoSignUp {
		account, err = s.serviceCreateAccount(c, req.Account, req.Password, req.IsCrypto)
	}
	if err != nil {
		log.Error().Err(err).Msg("Failed to shield login")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	var resp sdk.ShieldLoginResponseData
	resp.Account = &sdk.Account{
		UID:           sdk.ID(account.ID),
		Name:          account.Username,
		Email:         account.Email,
		IsEmailVerify: "0",
		Token:         account.LoginToken,
		Country:       "US",
		AreaCode:      "**",
	}
	resp.RealNameOperation = "None"
	c.JSON(http.StatusOK, sdk.NewResponse(0, &resp))
}

func (s *Server) serviceShieldLogin(ctx context.Context, username, password string, isCrypto bool) (record *store.Account, err error) {
	if !strings.Contains(username, "@") {
		record, err = s.store.Account().GetAccountByUsername(ctx, username)
	} else {
		record, err = s.store.Account().GetAccountByEmail(ctx, username)
	}
	if err != nil {
		return nil, err
	}
	if s.config.PassSignIn {
		if isCrypto {
			p, err := s.secret.Server.DecryptBase64(password)
			if err != nil {
				return nil, err
			}
			password = string(p)
		}
		if bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(password)) != nil {
			return nil, sdk.ErrInvalidPassword
		}
	}
	loginToken := make([]byte, 24)
	rand.Read(loginToken)
	record.LoginToken = base64.RawStdEncoding.EncodeToString(loginToken)
	return record, s.store.Account().UpdateAccountLoginToken(ctx, record.ID, record.LoginToken)
}

func (s *Server) serviceCreateAccount(ctx context.Context, username, password string, isCrypto bool) (record *store.Account, err error) {
	if !strings.Contains(username, "@") {
		record = &store.Account{Username: username, Email: username + "@" + s.config.BaseDomain}
		if err = s.store.Account().CreateAccount(ctx, record); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("email is not supported")
	}
	if s.config.PassSignIn {
		if isCrypto {
			p, err := s.secret.Server.DecryptBase64(password)
			if err != nil {
				return nil, err
			}
			password = string(p)
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		if err := s.store.Account().UpdateAccountPassword(ctx, record.ID, string(hash)); err != nil {
			return nil, err
		}
	}
	loginToken := make([]byte, 24)
	rand.Read(loginToken)
	record.LoginToken = base64.RawStdEncoding.EncodeToString(loginToken)
	return record, s.store.Account().UpdateAccountLoginToken(ctx, record.ID, record.LoginToken)
}

func (s *Server) handleSDKShieldVerify(c *gin.Context) {
	var req sdk.ShieldVerifyRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-210, nil))
		return
	}
	account, err := s.serviceShieldVerify(c, int64(req.UID), req.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to shield verify")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-210, nil))
		return
	}
	var resp sdk.ShieldVerifyResponseData
	resp.Account = &sdk.Account{
		UID:           req.UID,
		Email:         account.Email,
		Name:          account.Username,
		IsEmailVerify: "0",
		Token:         account.LoginToken,
		Country:       "US",
		AreaCode:      "**",
	}
	resp.RealNameOperation = "None"
	c.JSON(http.StatusOK, sdk.NewResponse(0, &resp))
}

func (s *Server) serviceShieldVerify(ctx context.Context, id int64, token string) (record *store.Account, err error) {
	record, err = s.store.Account().GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if record.LoginToken == "" || record.LoginToken != token {
		return nil, sdk.ErrInvalidLoginToken
	}
	return record, nil
}

func (s *Server) handleSDKComboLogin(c *gin.Context) {
	var req sdk.ComboLoginRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	var data sdk.ComboLoginData
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal data")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	account, err := s.serviceComboLogin(c, int64(data.UID), data.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to combo login")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-202, nil))
		return
	}
	var resp sdk.ComboLoginResponseData
	resp.ComboID = "0"
	resp.OpenID = sdk.ID(account.ID)
	resp.ComboToken = account.ComboToken
	if data.Guest {
		resp.Data = "{\"guest\":true}"
	} else {
		resp.Data = "{\"guest\":false}"
	}
	resp.Heartbeat = false
	resp.AccountType = 1
	resp.FatigueRemind = nil
	c.JSON(http.StatusOK, sdk.NewResponse(0, &resp))
}

func (s *Server) serviceComboLogin(ctx context.Context, id int64, token string) (record *store.Account, err error) {
	if record, err = s.serviceShieldVerify(ctx, id, token); err != nil {
		return nil, err
	}
	comboToken := make([]byte, 20)
	rand.Read(comboToken)
	record.ComboToken = hex.EncodeToString(comboToken)
	return record, s.store.Account().UpdateAccountComboToken(ctx, record.ID, record.ComboToken)
}
