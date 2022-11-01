package http

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/ec2b"
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
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.New()
	s.router.Use(s.logger.HandlerFunc())
	s.router.Use(gin.Recovery())
	s.store = store.NewStore(s.config)
	return s
}

func (s *Server) Start() (err error) {
	s.initRouter()
	s.server = &http.Server{Addr: s.config.HTTPServer.Addr, Handler: s.router}
	if tls := s.config.HTTPServer.TLS; !tls.Enable {
		err = s.server.ListenAndServe()
	} else {
		err = s.server.ListenAndServeTLS(tls.CertFile, tls.KeyFile)
	}
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) LoadSecret() error {
	_, err := os.Stat("data/secret.bin")
	if errors.Is(err, os.ErrNotExist) {
		s.secret.Shared = ec2b.NewEc2b()
		if err = os.WriteFile("data/secret.bin", s.secret.Shared.Bytes(), 0644); err != nil {
			return err
		}
		if err = os.WriteFile("data/secret.seed", []byte(strconv.FormatUint(s.secret.Shared.Seed(), 10)), 0644); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	p, err := os.ReadFile("data/secret.bin")
	if err != nil {
		return err
	}
	s.secret.Shared, err = ec2b.LoadKey(p)
	rest, _ := os.ReadFile("data/secret.pem")
	var block *pem.Block
	for {
		block, rest = pem.Decode(rest)
		switch block.Type {
		case "DISPATCH SERVER RSA PRIVATE KEY":
			s.secret.Server.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return err
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 2":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client["2"] = &PublicKey{k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 3":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client["3"] = &PublicKey{k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 4":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client["4"] = &PublicKey{k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 5":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client["5"] = &PublicKey{k}
			}
		}
		if len(rest) == 0 {
			break
		}
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}
	return s.server.Close()
}
