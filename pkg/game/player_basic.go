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

func (p *PlayerBasic) SetPlayerBornData(ctx *Context, id uint32, name string) error {
	p.Lock()
	defer p.Unlock()
	p.Level = 1
	p.Nickname = name
	p.OpenStateMap = map[uint32]uint32{
		OpenStateType_OPEN_STATE_AVATAR_PROMOTE:           1,
		OpenStateType_OPEN_STATE_AVATAR_TALENT:            1,
		OpenStateType_OPEN_STATE_WEAPON_PROMOTE:           1,
		OpenStateType_OPEN_STATE_WEAPON_AWAKEN:            1,
		OpenStateType_OPEN_STATE_WEAPON_UPGRADE:           1,
		OpenStateType_OPEN_STATE_RELIQUARY_UPGRADE:        1,
		OpenStateType_OPEN_STATE_RELIQUARY_PROMOTE:        1,
		OpenStateType_OPEN_STATE_GUIDE_TALENT:             1,
		OpenStateType_OPEN_STATE_GUIDE_RELIC_PROM:         1,
		OpenStateType_OPEN_STATE_RESIN:                    1,
		OpenStateType_OPEN_STATE_PHOTOGRAPH:               1,
		OpenStateType_OPEN_STATE_DERIVATIVE_MALL:          1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_RECOMMANDED:    1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_GENESISCRYSTAL: 1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_GIFTPACKAGE:    1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_PAIMON:         1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_BLACKSMITH:     1,
		OpenStateType_OPEN_STATE_SHOP_TYPE_VIRTUAL_SHOP:   1,
		OpenStateType_OPEN_STATE_BATTLE_PASS:              1,
	}
	if p.player.ID == 1 {
		p.OpenStateMap[OpenStateType_OPEN_STATE_PAIMON] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_PAIMON_NAVIGATION] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_COOK] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_FRESHMAN_GUIDE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_COMBINE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_GACHA] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_FORGE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_EXPEDITION] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_TOWER] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_LIMIT_REGION_FRESHMEAT] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_LIMIT_REGION_GLOBAL] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_MULTIPLAYER] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_AVATAR_FASHION] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_PERSONAL_LINE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_CITY_REPUATION_MENGDE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_CITY_REPUATION_LIYUE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_CITY_REPUATION_UI_HINT] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_CITY_REPUATION_INAZUMA] = 1
		p.OpenStateMap[OpenStateType_OPEN_ADVENTURE_MANUAL] = 1
		p.OpenStateMap[OpenStateType_OPEN_ADVENTURE_MANUAL_CITY_MENGDE] = 1
		p.OpenStateMap[OpenStateType_OPEN_ADVENTURE_MANUAL_CITY_LIYUE] = 1
		p.OpenStateMap[OpenStateType_OPEN_ADVENTURE_MANUAL_MONSTER] = 1
		p.OpenStateMap[OpenStateType_OPEN_ADVENTURE_MANUAL_BOSS_DUNGEON] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_ACTIVITY_ENTRY_OPEN] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_MIRACLE_RING] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_COOP_LINE] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_FISH] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_HOME] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_HOME_AVATAR] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_HOME_EDIT] = 1
		p.OpenStateMap[OpenStateType_OPEN_STATE_TELEPORT_HUD] = 1
	}
	p.HeadImageAvatarId = id
	return nil
}

func (p *PlayerBasic) SetNickname(ctx *Context, name string) error {
	p.Lock()
	defer p.Unlock()
	p.Nickname = name
	return nil
}

func (p *PlayerBasic) SetOpenState(ctx *Context, key, value uint32) error {
	p.Lock()
	defer p.Unlock()
	if p.OpenStateMap == nil {
		p.OpenStateMap = map[uint32]uint32{}
	}
	p.OpenStateMap[key] = value
	return nil
}
