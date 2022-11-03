package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *Server) handleDefault(c *gin.Context) {
	c.JSON(http.StatusOK, "!>.<")
}

func (s *Server) handleOptionsPing(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func (s *Server) handlePing(c *gin.Context) {
	c.String(http.StatusOK, "ok")
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
