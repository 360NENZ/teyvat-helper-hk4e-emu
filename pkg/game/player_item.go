package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-emu/pkg/pb"
)

type PlayerItem struct {
	sync.RWMutex
	*pb.PlayerItemCompBin
	player *Player
}

func (p *PlayerItem) AddItem(id uint32) (*pb.ItemBin, error) {
	var item pb.ItemBin
	item.ItemId = id
	item.Guid = p.player.Basic().GetNextGuidSeqId()
	item.Detail = &pb.ItemBin_Equip{Equip: &pb.EquipBin{
		IsLocked: false,
		Detail:   &pb.EquipBin_Weapon{Weapon: &pb.WeaponBin{Level: 1}},
	}}
	return &item, nil
}
