package game

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-emu/pkg/net"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
)

type Server struct {
	config   *config.Config
	random   *rand.Rand
	secret   *Secret
	conn     *net.KCPConn
	store    *store.Store
	sessions map[uint32]*PlayerSession
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{}
	s.config = cfg
	s.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.secret = NewSecret()
	s.store = store.NewStore(s.config)
	s.sessions = make(map[uint32]*PlayerSession)
	return s
}

func (s *Server) createPlayerSession(sess *net.Session) (*PlayerSession, error) {
	session := NewPlayerSession(sess, s.secret, s)
	s.sessions[session.ID()] = session
	return session, nil
}

func (s *Server) run() {
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
}

func (s *Server) Store() *store.Store { return s.store }

func (s *Server) Start() error {
	conn, err := net.Listen(s.config.GameServer.Addr)
	if err != nil {
		return err
	}
	s.conn = conn
	go s.run()
	select {}
}

func (s *Server) UpdateSeed() {
	b, _ := os.ReadFile("data/secret.seed")
	seed, _ := strconv.ParseUint(string(b), 10, 64)
	s.secret.Shared.SetSeed(seed)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}
