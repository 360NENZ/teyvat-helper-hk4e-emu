package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

type PlayerBasic struct {
	sync.RWMutex
	*pb.PlayerBasicCompBin
	player *Player
}

func (p *PlayerBasic) GetNextGuidSeqId() uint64 {
	p.Lock()
	defer p.Unlock()
	p.GuidSeqId++
	return uint64(p.player.ID)<<32 | uint64(p.GetGuidSeqId())
}
