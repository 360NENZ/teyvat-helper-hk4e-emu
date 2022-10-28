package net

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	ErrSessionTokenMismatch = fmt.Errorf("session token mismatch")
	ErrSessionNotFound      = fmt.Errorf("session not found")
	ErrInvalidPacket        = fmt.Errorf("invalid packet")
)

type KCPConn struct {
	sync.Mutex
	conn     *net.UDPConn
	rand     *rand.Rand
	sessions map[uint32]*Session
	packetCh chan *Packet
}

func newKCPConn(conn *net.UDPConn) *KCPConn {
	l := &KCPConn{
		conn:     conn,
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		sessions: make(map[uint32]*Session),
		packetCh: make(chan *Packet, 1024),
	}
	go l.start()
	return l
}

func (l *KCPConn) Packet() chan *Packet { return l.packetCh }

func (l *KCPConn) start() {
	buf := make([]byte, mtuLimit)
	for {
		n, addr, err := l.conn.ReadFromUDP(buf)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read from UDP")
			break
		}
		l.handlePacket(addr, buf, n)
	}
}

func (l *KCPConn) handlePacket(addr *net.UDPAddr, buf []byte, n int) {
	if n == 20 {
		if err := l.handleCtrlPacket(addr, buf[:n]); err != nil {
			log.Error().Err(err).Msg("Failed to handle control packet")
		}
	} else if n >= 28 {
		if err := l.handleDataPacket(addr, buf[:n]); err != nil {
			log.Error().Err(err).Msg("Failed to handle data packet")
		}
	}
}

func (l *KCPConn) nextSessionID() (id uint32) {
	for {
		id = l.rand.Uint32()
		if _, ok := l.sessions[id]; !ok {
			return id
		}
	}
}

func (l *KCPConn) nextSessionToken() uint32 { return l.rand.Uint32() }

func (l *KCPConn) createSession(addr *net.UDPAddr, token uint32) (*Session, error) {
	session := newSession(addr, l.nextSessionID())
	session.SetToken(token)
	session.SetSendFunc(l.sendDataPacket)
	l.sessions[session.id] = session
	log.Debug().Uint32("session_id", session.id).Str("remote_addr", session.remote.String()).Msg("Session created")
	return session, nil
}

func (l *KCPConn) deleteSession(addr *net.UDPAddr, id, token uint32) (*Session, error) {
	session, ok := l.sessions[id]
	if !ok {
		return nil, ErrSessionNotFound
	}
	if session == nil {
		delete(l.sessions, id)
		return nil, ErrSessionNotFound
	}
	if session.remote.String() != addr.String() && session.token != token {
		return nil, ErrSessionTokenMismatch
	}
	delete(l.sessions, id)
	log.Debug().Uint32("session_id", session.id).Str("remote_addr", session.remote.String()).Msg("Session deleted")
	return session, nil
}

func (l *KCPConn) updateSession(addr *net.UDPAddr, id, token uint32) (*Session, error) {
	session, ok := l.sessions[id]
	if !ok {
		return nil, ErrSessionNotFound
	}
	if session == nil {
		delete(l.sessions, id)
		return nil, ErrSessionNotFound
	}
	log.Debug().Uint32("session_id", session.id).Str("remote_addr", session.remote.String()).Msg("Session updated")
	return session, nil
}

func (l *KCPConn) getSession(addr *net.UDPAddr, id, token uint32) (*Session, error) {
	session, ok := l.sessions[id]
	if !ok {
		return nil, ErrSessionNotFound
	}
	if session == nil {
		delete(l.sessions, id)
		return nil, ErrSessionNotFound
	}
	if session.token != token {
		return nil, ErrSessionTokenMismatch
	}
	return session, nil
}

func DialKCP(network string, laddr, raddr *net.UDPAddr) (*KCPConn, error) {
	return nil, nil
}

func ListenKCP(network string, laddr *net.UDPAddr) (*KCPConn, error) {
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}
	l := newKCPConn(conn)
	go l.start()
	return l, nil
}
