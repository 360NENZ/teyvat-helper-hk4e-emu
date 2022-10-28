package net

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

type Session struct {
	id, token uint32
	kcp       *KCP
	remote    *net.UDPAddr
	sendFn    SendFunc
}

type SendFunc func(*net.UDPAddr, []byte) error

var NullSendFunc SendFunc = func(*net.UDPAddr, []byte) error { return nil }

func newSession(addr *net.UDPAddr, id uint32) *Session {
	session := &Session{
		id:     id,
		remote: addr,
		sendFn: NullSendFunc,
	}
	session.kcp = NewKCP(id, session.output)
	session.kcp.SetMtu(1200)
	session.kcp.WndSize(256, 256)
	session.kcp.NoDelay(1, 10, 2, 1)
	return session
}

func (s *Session) output(buf []byte, size int) {
	if err := s.sendFn(s.remote, buf[:size]); err != nil {
		log.Error().Err(err).Msg("Failed to send packet")
	}
}

func (s *Session) ID() uint32           { return s.id }
func (s *Session) Remote() *net.UDPAddr { return s.remote }

func (s *Session) SetToken(token uint32) {
	s.token = token
	s.kcp.SetToken(token)
}

func (s *Session) SetSendFunc(fn SendFunc) { s.sendFn = fn }

func (s *Session) OnPacket(buf []byte, ch chan<- *Packet) error {
	if ret := s.kcp.Input(buf, true, true); ret < 0 {
		return fmt.Errorf("kcp input error %d", ret)
	}
	n := s.kcp.PeekSize()
	for n > 0 {
		packet := newPacket(s)
		packet.data = packet.data[:n]
		_ = s.kcp.Recv(packet.data)
		ch <- packet
		n = s.kcp.PeekSize()
	}
	s.kcp.Update()
	return nil
}

func (s *Session) Send(buf []byte, flush ...bool) error {
	if ret := s.kcp.Send(buf); ret < 0 {
		return fmt.Errorf("kcp send error %d", ret)
	}
	s.kcp.Update()
	return nil
}
