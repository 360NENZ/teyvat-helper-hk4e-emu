package http

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

var (
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidLoginToken = errors.New("invalid login token")
	ErrInvalidComboToken = errors.New("invalid combo token")
)

type sdkResponse struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Retcode int32  `json:"retcode"`
}

var sdkResponseMessage = map[int32]string{
	0:    "OK",
	-101: "系统错误",
	-102: "密码格式错误，密码格式为8-30位，并且由数字、大小写字母、英文特殊符号两种以上组合",
	-103: "参数错误",
	-104: "缺少配置",
	-106: "协议加载失败",
	-107: "渠道错误",
	-111: "???",
	-115: "请前往官网/商店下载最新版本",
	-202: "账号或密码错误",
	-210: "为了您的账号安全，请重新登录。",
}

func newSDKResponse(retcode int32, data any) *sdkResponse {
	return &sdkResponse{
		Data:    data,
		Message: sdkResponseMessage[retcode],
		Retcode: retcode,
	}
}

func (s *Server) handleSDKRiskyCheck(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, gin.H{
		"id":      "",
		"action":  "ACTION_NONE",
		"geetest": nil,
	}))
}

type ID int64

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(id), 10) + `"`), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*id = ID(i)
	return nil
}

type sdkShieldLoginRequestData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

type sdkAccount struct {
	UID               ID     `json:"uid"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	IsEmailVerify     string `json:"is_email_verify"`
	RealName          string `json:"realname"`
	IdentityCard      string `json:"identity_card"`
	Token             string `json:"token"`
	SafeMobile        string `json:"safe_mobile"`
	FacebookName      string `json:"facebook_name"`
	GoogleName        string `json:"google_name"`
	TwitterName       string `json:"twitter_name"`
	GameCenterName    string `json:"game_center_name"`
	AppleName         string `json:"apple_name"`
	SonyName          string `json:"sony_name"`
	TapName           string `json:"tap_name"`
	Country           string `json:"country"`
	ReactivateTicket  string `json:"reactivate_ticket"`
	AreaCode          string `json:"area_code"`
	DeviceGrantTicket string `json:"device_grant_ticket"`
	SteamName         string `json:"steam_name"`
}

type sdkShieldLoginResponseData struct {
	Account             *sdkAccount `json:"account"`
	DeviceGrantRequired bool        `json:"device_grant_required"`
	SafeMobileRequired  bool        `json:"safe_mobile_required"`
	RealPersonRequired  bool        `json:"realperson_required"`
	ReactivateRequired  bool        `json:"reactivate_required"`
	RealNameOperation   string      `json:"realname_operation"`
}

func (s *Server) handleSDKShieldLogin(c *gin.Context) {
	var req sdkShieldLoginRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	if req.Account == "" {
		log.Error().Err(fmt.Errorf("account is empty")).Msg("Bad request")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	account, err := s.serviceShieldLogin(c, req.Account, "")
	if err == sql.ErrNoRows {
		account, err = s.serviceCreateAccount(c, req.Account, "")
	}
	if err != nil {
		log.Error().Err(err).Msg("Failed to shield login")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	var resp sdkShieldLoginResponseData
	resp.Account = &sdkAccount{
		UID:           ID(account.ID),
		Name:          account.Username,
		Email:         account.Email,
		IsEmailVerify: "0",
		Token:         account.LoginToken,
		Country:       "US",
		AreaCode:      "**",
	}
	resp.RealNameOperation = "None"
	c.JSON(http.StatusOK, newSDKResponse(0, &resp))
}

func (s *Server) serviceShieldLogin(ctx context.Context, username, password string) (record *store.Account, err error) {
	if !strings.Contains(username, "@") {
		record, err = s.store.Account().GetAccountByUsername(ctx, username)
	} else {
		record, err = s.store.Account().GetAccountByEmail(ctx, username)
	}
	if err != nil {
		return nil, err
	}
	if record.Password != "" && record.Password != password {
		return nil, ErrInvalidPassword
	}
	record.Password = ""
	loginToken := make([]byte, 24)
	rand.Read(loginToken)
	record.LoginToken = base64.RawStdEncoding.EncodeToString(loginToken)
	return record, s.store.Account().UpdateAccountLoginToken(ctx, record.ID, record.LoginToken)
}

func (s *Server) serviceCreateAccount(ctx context.Context, username, password string) (record *store.Account, err error) {
	if !strings.Contains(username, "@") {
		record = &store.Account{Username: username, Email: username + "@" + s.config.BaseDomain}
		if err = s.store.Account().CreateAccount(ctx, record); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("email is not supported")
	}
	loginToken := make([]byte, 24)
	rand.Read(loginToken)
	record.LoginToken = base64.RawStdEncoding.EncodeToString(loginToken)
	return record, s.store.Account().UpdateAccountLoginToken(ctx, record.ID, record.LoginToken)
}

type sdkShieldVerifyRequestData struct {
	UID   ID     `json:"uid"`
	Token string `json:"token"`
}

type sdkShieldVerifyResponseData struct {
	Account             *sdkAccount `json:"account"`
	DeviceGrantRequired bool        `json:"device_grant_required"`
	SafeMobileRequired  bool        `json:"safe_mobile_required"`
	RealPersonRequired  bool        `json:"realperson_required"`
	RealNameOperation   string      `json:"realname_operation"`
}

func (s *Server) handleSDKShieldVerify(c *gin.Context) {
	var req sdkShieldVerifyRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-210, nil))
		return
	}
	account, err := s.serviceShieldVerify(c, int64(req.UID), req.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to shield verify")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-210, nil))
		return
	}
	var resp sdkShieldVerifyResponseData
	resp.Account = &sdkAccount{
		UID:           req.UID,
		Email:         account.Email,
		Name:          account.Username,
		IsEmailVerify: "0",
		Token:         account.LoginToken,
		Country:       "US",
		AreaCode:      "**",
	}
	resp.RealNameOperation = "None"
	c.JSON(http.StatusOK, newSDKResponse(0, &resp))
}

func (s *Server) serviceShieldVerify(ctx context.Context, id int64, token string) (record *store.Account, err error) {
	record, err = s.store.Account().GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if record.LoginToken == "" || record.LoginToken != token {
		return nil, ErrInvalidLoginToken
	}
	return record, nil
}

type sdkComboLoginRequestData struct {
	AppID     int32  `json:"app_id"`
	ChannelID int32  `json:"channel_id"`
	Data      string `json:"data"`
	Device    string `json:"device"`
	Sign      string `json:"sign"`
}

type sdkComboLoginData struct {
	UID   ID     `json:"uid"`
	Guest bool   `json:"guest"`
	Token string `json:"token"`
}

type sdkComboLoginResponseData struct {
	ComboID       string `json:"combo_id"`
	OpenID        ID     `json:"open_id"`
	AccountType   int32  `json:"account_type"`
	ComboToken    string `json:"combo_token"`
	Data          string `json:"data"`
	FatigueRemind any    `json:"fatigue_remind"`
	Heartbeat     bool   `json:"heartbeat"`
}

func (s *Server) handleSDKComboLogin(c *gin.Context) {
	var req sdkComboLoginRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	var data sdkComboLoginData
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal data")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	account, err := s.serviceComboLogin(c, int64(data.UID), data.Token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to combo login")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-202, nil))
		return
	}
	var resp sdkComboLoginResponseData
	resp.ComboID = "0"
	resp.OpenID = ID(account.ID)
	resp.ComboToken = account.ComboToken
	if data.Guest {
		resp.Data = "{\"guest\":true}"
	} else {
		resp.Data = "{\"guest\":false}"
	}
	resp.Heartbeat = false
	resp.AccountType = 1
	resp.FatigueRemind = nil
	c.JSON(http.StatusOK, newSDKResponse(0, &resp))
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
