package game

import (
	"sync"
	"time"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

type PlayerAvatar struct {
	sync.RWMutex
	*pb.PlayerAvatarCompBin
	player *Player
}

func (p *PlayerAvatar) SetPlayerBornData(ctx *Context, id uint32) error {
	p.Lock()
	defer p.Unlock()
	avatar, err := p.AddAvatar(id)
	if err != nil {
		return err
	}
	avatarGuid := avatar.GetGuid()
	p.TeamMap = map[uint32]*pb.AvatarTeamBin{
		1: {AvatarGuidList: []uint64{avatarGuid}},
		2: {},
		3: {},
		4: {},
	}
	p.CurTeamId = 1
	p.ChooseAvatarGuid = avatarGuid
	p.OwnedFlycloakList = []uint32{140001}
	p.OwnedCostumeIdList = []uint32{200302, 202101, 204101, 204501}
	return nil
}

func (p *PlayerAvatar) AddAvatar(id uint32) (*pb.AvatarBin, error) {
	var avatar pb.AvatarBin
	avatar.AvatarType = 1
	avatar.AvatarId = id
	avatar.Guid = p.player.Basic().GetNextGuidSeqId()
	avatar.Level = 1
	avatar.LifeState = 1
	if id == 10000005 {
		avatar.SkillDepotId = 501
		avatar.DepotMap = map[uint32]*pb.AvatarSkillDepotBin{
			501: {SkillLevelMap: map[uint32]uint32{100540: 1}},
		}
	} else if id == 10000007 {
		avatar.SkillDepotId = 701
		avatar.DepotMap = map[uint32]*pb.AvatarSkillDepotBin{
			701: {SkillLevelMap: map[uint32]uint32{100550: 1}},
		}
	}
	item, err := p.player.Item().AddItem(11101)
	if err != nil {
		return nil, err
	}
	avatar.EquipList = []*pb.ItemBin{item}
	avatar.FlycloakId = 140001
	avatar.BornTime = uint32(time.Now().Unix())
	p.AvatarList = append(p.AvatarList, &avatar)
	return &avatar, nil
}
