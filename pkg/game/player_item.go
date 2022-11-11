package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-proto/pb"
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
	packStore := p.GetPackStore()
	if packStore == nil {
		packStore = &pb.ItemStoreBin{}
		p.PackStore = packStore
	}
	itemList := packStore.GetItemList()
	if itemList == nil {
		itemList = []*pb.ItemBin{}
	}
	packStore.ItemList = append(itemList, &item)
	return &item, nil
}
