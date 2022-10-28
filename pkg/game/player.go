package game

import (
	"errors"

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
	scene  *PlayerScene
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
	sceneBin := p.binary.GetSceneBin()
	if sceneBin == nil {
		sceneBin = &pb.PlayerSceneCompBin{}
		p.binary.SceneBin = sceneBin
	}
	p.scene = &PlayerScene{PlayerSceneCompBin: sceneBin, player: p}
}

func (p *Player) Basic() *PlayerBasic   { return p.basic }
func (p *Player) Avatar() *PlayerAvatar { return p.avatar }
func (p *Player) Item() *PlayerItem     { return p.item }
func (p *Player) Social() *PlayerSocial { return p.social }
func (p *Player) Scene() *PlayerScene   { return p.scene }

func (p *Player) SetPlayerBornData(ctx *Context, id uint32, name string) error {
	if id != 10000005 && id != 10000007 {
		return errors.New("avatar id is not allowed")
	}
	if err := p.basic.SetPlayerBornData(ctx, id, name); err != nil {
		return err
	}
	if err := p.avatar.SetPlayerBornData(ctx, id); err != nil {
		return err
	}
	if err := p.social.SetPlayerBornData(ctx, id); err != nil {
		return err
	}
	if err := p.scene.SetPlayerBornData(ctx, id); err != nil {
		return err
	}
	return nil
}
