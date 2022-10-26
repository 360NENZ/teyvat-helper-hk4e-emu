package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleDefault(c *gin.Context) {
	c.String(http.StatusOK, "ok")
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
		log.Printf("[WARN] ignore unknown request, safely: %s", c.Request.URL.Path)
	}
}
