package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

type PlayerScene struct {
	sync.RWMutex
	*pb.PlayerSceneCompBin
	player *Player
}

func (p *PlayerScene) SetPlayerBornData(ctx *Context, id uint32) error {
	p.Lock()
	defer p.Unlock()
	p.World = &pb.WorldBin{
		SceneMap: map[uint32]*pb.SceneBin{},
		Level:    1,
	}
	p.CurSceneOwnerUid = id
	p.MyPrevSceneId = 3
	p.MyPrevPos = &pb.VectorBin{X: 2747.562, Y: 194.633, Z: -1719.386}
	p.MyPrevRot = &pb.VectorBin{Y: 307}
	p.MyCurSceneId = 3
	return nil
}

func (p *PlayerScene) SetSceneId(id uint32) error {
	p.Lock()
	defer p.Unlock()
	p.MyPrevSceneId = p.MyCurSceneId
	p.MyCurSceneId = id
	return nil
}

func (p *PlayerScene) GetPos() *pb.Vector {
	p.RLock()
	defer p.RUnlock()
	pos := p.GetMyPrevPos()
	return &pb.Vector{X: pos.GetX(), Y: pos.GetY(), Z: pos.GetZ()}
}

func (p *PlayerScene) SetPos(pos *pb.Vector) error {
	p.Lock()
	defer p.Unlock()
	p.MyPrevPos = &pb.VectorBin{X: pos.GetX(), Y: pos.GetY(), Z: pos.GetZ()}
	return nil
}

func (p *PlayerScene) GetRot() *pb.Vector {
	p.RLock()
	defer p.RUnlock()
	rot := p.GetMyPrevRot()
	return &pb.Vector{X: rot.GetX(), Y: rot.GetY(), Z: rot.GetZ()}
}

func (p *PlayerScene) SetRot(rot *pb.Vector) error {
	p.Lock()
	defer p.Unlock()
	p.MyPrevRot = &pb.VectorBin{X: rot.GetX(), Y: rot.GetY(), Z: rot.GetZ()}
	return nil
}
