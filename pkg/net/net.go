package net

import (
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
