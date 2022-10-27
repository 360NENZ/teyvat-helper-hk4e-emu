package game

import (
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
	"github.com/teyvat-helper/hk4e-proto/pb"
)

type Player struct {
	*store.Player
	binary *pb.PlayerDataBin

	basic  *PlayerBasic
	avatar *PlayerAvatar
	item   *PlayerItem
	social *PlayerSocial
}

func NewPlayer(player *store.Player) *Player {
	return &Player{Player: player}
}

func (p *Player) SetBinary(binary *pb.PlayerDataBin) {
	p.binary = binary
	basicBin := p.binary.GetBasicBin()
	if basicBin == nil {
		basicBin = &pb.PlayerBasicCompBin{}
		p.binary.BasicBin = basicBin
	}
	p.basic = &PlayerBasic{PlayerBasicCompBin: basicBin, player: p}
	avatarBin := p.binary.GetAvatarBin()
	if avatarBin == nil {
		avatarBin = &pb.PlayerAvatarCompBin{}
		p.binary.AvatarBin = avatarBin
	}
	p.avatar = &PlayerAvatar{PlayerAvatarCompBin: avatarBin, player: p}
	itemBin := p.binary.GetItemBin()
	if itemBin == nil {
		itemBin = &pb.PlayerItemCompBin{}
		p.binary.ItemBin = itemBin
	}
	p.item = &PlayerItem{PlayerItemCompBin: itemBin, player: p}
	socialBin := p.binary.GetSocialBin()
	if socialBin == nil {
		socialBin = &pb.PlayerSocialCompBin{}
		p.binary.SocialBin = socialBin
	}
	p.social = &PlayerSocial{PlayerSocialCompBin: socialBin, player: p}
}

func (p *Player) Basic() *PlayerBasic   { return p.basic }
func (p *Player) Avatar() *PlayerAvatar { return p.avatar }
func (p *Player) Item() *PlayerItem     { return p.item }
func (p *Player) Social() *PlayerSocial { return p.social }
