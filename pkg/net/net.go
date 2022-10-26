package net

import (
	"net"
	"sync"
)

var bufPool = &sync.Pool{
	New: func() interface{} { return make([]byte, 256*1024) },
}

type Packet struct {
	sess *Session
	data []byte
}

func newPacket(sess *Session) *Packet {
	return &Packet{sess: sess, data: bufPool.Get().([]byte)}
}

func (p *Packet) Release() { bufPool.Put(p.data) }

func (p *Packet) Session() *Session { return p.sess }
func (p *Packet) Payload() []byte   { return p.data }

func Listen(addr string) (*KCPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	return ListenKCP("udp", udpAddr)
}
