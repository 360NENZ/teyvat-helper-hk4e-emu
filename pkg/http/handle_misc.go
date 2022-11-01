package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *Server) handleDefault(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, nil))
}

func (s *Server) handlePing(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}


func (s *Server) handleServerStatus(c *gin.Context) {
	c.String(http.StatusOK, "{\"retcode\":0,\"status\":{\"playerCount\":0,\"maxPlayer\":-1,\"version\":\"3.2.0\"}}")
}

func (s *Server) handleLogUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

func (s *Server) handleWebStaticJSON(c *gin.Context) {
	if c.Request.URL.Path == "/admin/mi18n/plat_oversea/m2020030410/m2020030410-version.json" {
		c.JSON(http.StatusOK, gin.H{"version": 65})
	} else {
		log.Warn().Str("path", c.Request.URL.Path).Msg("ignore unknown request, safely")
	}
}
