package game

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/http"
	"github.com/teyvat-helper/hk4e-emu/pkg/net"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

type Server struct {
	config   *config.Config
	random   *rand.Rand
	secret   *Secret
	filter   *PacketFilter
	conn     *net.KCPConn
	store    *store.Store
	sessions map[uint32]*PlayerSession
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{}
	s.config = cfg
	s.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.secret = NewSecret()
	s.filter = NewPacketFilter(s.config)
	s.store = store.NewStore(s.config)
	s.sessions = make(map[uint32]*PlayerSession)
	return s
}

func (s *Server) createPlayerSession(sess *net.Session) (*PlayerSession, error) {
	session := NewPlayerSession(sess, s.secret, s)
	s.sessions[session.ID()] = session
	log.Debug().Uint32("session_id", session.ID()).Str("remote_addr", session.Remote().String()).Msg("PlayerSession created")
	return session, nil
}

func (s *Server) deletePlayerSession(ctx context.Context, session *PlayerSession) (*PlayerSession, error) {
	if _, err := session.UpdatePlayer(ctx); err != nil {
		return session, err
	}
	delete(s.sessions, session.ID())
	log.Debug().Uint32("session_id", session.ID()).Str("remote_addr", session.Remote().String()).Msg("PlayerSession deleted")
	return session, nil
}

func (s *Server) run() error {
	var err error
	for packet := range s.conn.Packet() {
		session, ok := s.sessions[packet.Session().ID()]
		if !ok {
			session, err = s.createPlayerSession(packet.Session())
			if err != nil {
				fmt.Println("[GAME] Failed to create player session, error:", err)
				continue
			}
		}
		go session.handlePacket(packet, s.onPacket)
	}
	return nil
}

func (s *Server) Store() *store.Store { return s.store }

func (s *Server) Start() error {
	conn, err := net.Listen(s.config.GameServer.Addr)
	if err != nil {
		return err
	}
	s.conn = conn
	log.Info().Str("listen_addr", s.config.GameServer.Addr).Msg("GAME server is starting")
	return s.run()
}

func (s *Server) LoadSecret() error {
	b, err := os.ReadFile("data/secret.seed")
	if err != nil {
		return err
	}
	seed, err := strconv.ParseUint(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return err
	}
	s.secret.Shared.SetSeed(seed)
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
				s.secret.Client[2] = &http.PublicKey{PublicKey: k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 3":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client[3] = &http.PublicKey{PublicKey: k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 4":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client[4] = &http.PublicKey{PublicKey: k}
			}
		case "DISPATCH CLIENT RSA PUBLIC KEY 5":
			k, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			} else if k, ok := k.(*rsa.PublicKey); !ok {
				return errors.New("invalid public key")
			} else {
				s.secret.Client[5] = &http.PublicKey{PublicKey: k}
			}
		}
		if len(rest) == 0 {
			break
		}
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	for _, session := range s.sessions {
		if _, err := s.deletePlayerSession(ctx, session); err != nil {
			log.Error().Err(err).Msg("Failed to update player")
		}
	}
	return nil
}
