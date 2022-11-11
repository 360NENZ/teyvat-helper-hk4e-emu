package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

type PlayerSocial struct {
	sync.RWMutex
	*pb.PlayerSocialCompBin
	player *Player
}

func (p *PlayerSocial) SetPlayerBornData(ctx *Context, id uint32) error {
	p.Lock()
	defer p.Unlock()
	p.Birthday = &pb.Birthday{}
	p.NameCardId = 210001
	return nil
}

func (p *PlayerSocial) SetBirthday(ctx *Context, birthday *pb.Birthday) error {
	p.Lock()
	defer p.Unlock()
	p.Birthday = birthday
	return nil
}

func (p *PlayerSocial) SetSignature(ctx *Context, signature string) error {
	p.Lock()
	defer p.Unlock()
	p.Signature = signature
	return nil
}
