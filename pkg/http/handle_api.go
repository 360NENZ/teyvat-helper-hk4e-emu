package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg"
	"github.com/teyvat-helper/hk4e-emu/pkg/sdk"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

func (s *Server) handleAPIPublicKey(c *gin.Context) {
	c.String(http.StatusOK, s.secret.Server.PublicKeyPEM)
}

func (s *Server) handleAPIStatus(c *gin.Context) {
	// TODO: not stable yet
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
		"buildVersion": pkg.BuildVersion,
		"protoVersion": pkg.ProtoVersion,
		"maxPlayer":    -1,
		"playerCount":  0, // TODO: get active player sessions from store
	}))
}

func (s *Server) handleAPIStatusLegacy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"retcode": 0,
		"status": gin.H{
			"buildVersion": pkg.BuildVersion,
			"version":      pkg.ProtoVersion,
			"maxPlayer":    -1,
			"playerCount":  0, // TODO: get active player sessions from store
		},
	})
}

func (s *Server) handleAPITokenCheck(c *gin.Context) {
	var req sdk.TokenCheckRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind JSON")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-210, nil))
		return
	}
	_, err := s.serviceCheckComboToken(c, int64(req.OpenID), req.ComboToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check combo token")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-210, nil))
		return
	}
	var resp sdk.TokenCheckResponseData
	resp.AccountType = 1
	resp.IPInfo.CountryCode = "us"
	c.JSON(http.StatusOK, sdk.NewResponse(0, &resp))
}

func (s *Server) serviceCheckComboToken(ctx context.Context, id int64, token string) (*store.Account, error) {
	record, err := s.store.Account().GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	if record.ComboToken == "" || record.ComboToken != token {
		return nil, sdk.ErrInvalidComboToken
	}
	return record, nil
}
