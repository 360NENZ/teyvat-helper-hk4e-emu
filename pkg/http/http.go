package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

type Server struct {
	config *config.Config
	logger *Logger
	secret *Secret
	router *gin.Engine
	server *http.Server
	store  *store.Store
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{}
	s.config = cfg
	s.logger = NewLogger()
	s.secret = NewSecret()
	s.router = gin.New()
	s.router.Use(s.logger.HandlerFunc())
	s.router.Use(gin.Recovery())
	s.initRouter()
	s.store = store.NewStore(s.config)
	return s
}

func (s *Server) Start() error {
	s.server = &http.Server{Addr: s.config.HTTPServer.Addr, Handler: s.router}
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}
	return s.server.Close()
}
