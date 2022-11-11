package game

import "github.com/teyvat-helper/hk4e-proto/pb"

func (s *Server) SendAvatarAddNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarDelNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetUpAvatarTeamReq(ctx *Context, req *pb.SetUpAvatarTeamReq) error {
	panic("not implement")
}

func (s *Server) SendSetUpAvatarTeamRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleChooseCurAvatarTeamReq(ctx *Context, req *pb.ChooseCurAvatarTeamReq) error {
	panic("not implement")
}

func (s *Server) SendChooseCurAvatarTeamRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleChangeAvatarReq(ctx *Context, req *pb.ChangeAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendChangeAvatarRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarPromoteReq(ctx *Context, req *pb.AvatarPromoteReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarPromoteRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSpringUseReq(ctx *Context, req *pb.SpringUseReq) error {
	panic("not implement")
}

func (s *Server) SendSpringUseRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleRefreshBackgroundAvatarReq(ctx *Context, req *pb.RefreshBackgroundAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendRefreshBackgroundAvatarRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarTeamUpdateNotify(ctx *Context) error {
	panic("not implement")
}

// send AvatarDataNotify
//
//	flow:
//		SEND ··> AvatarDataNotify
func (s *Server) SendAvatarDataNotify(ctx *Context) (err error) {
	var notify pb.AvatarDataNotify
	avatar := ctx.Session().GetPlayer().Avatar()
	for _, v := range avatar.GetAvatarList() {
		avatar := &pb.AvatarInfo{
			AvatarType:   v.GetAvatarType(),
			AvatarId:     v.GetAvatarId(),
			Guid:         v.GetGuid(),
			LifeState:    v.GetLifeState(),
			SkillDepotId: v.GetSkillDepotId(),
		}
		avatar.SkillLevelMap = v.GetDepotMap()[avatar.GetSkillDepotId()].GetSkillLevelMap()
		for _, v := range v.GetEquipList() {
			avatar.EquipGuidList = append(avatar.EquipGuidList, v.GetGuid())
		}
		avatar.WearingFlycloakId = v.GetFlycloakId()
		avatar.BornTime = v.GetBornTime()
		avatar.PendingPromoteRewardList = []uint32{1, 3, 5}
		avatar.PropMap = PropMap{
			PropType_PROP_EXP:                    0,
			PropType_PROP_BREAK_LEVEL:            0,
			PropType_PROP_SATIATION_VAL:          0,
			PropType_PROP_SATIATION_PENALTY_TIME: 0,
			PropType_PROP_LEVEL:                  1,
		}.ToPropMap()
		avatar.FightPropMap = FightPropMap{
			FightPropType_FIGHT_PROP_BASE_HP:           911.791,
			FightPropType_FIGHT_PROP_BASE_ATTACK:       41.053,
			FightPropType_FIGHT_PROP_ATTACK_PERCENT:    0,
			FightPropType_FIGHT_PROP_BASE_DEFENSE:      57.225,
			FightPropType_FIGHT_PROP_CRITICAL:          0.05,
			FightPropType_FIGHT_PROP_ANTI_CRITICAL:     0,
			FightPropType_FIGHT_PROP_CRITICAL_HURT:     0.5,
			FightPropType_FIGHT_PROP_CHARGE_EFFICIENCY: 1,
			FightPropType_FIGHT_PROP_HEAL_ADD:          0,
			FightPropType_FIGHT_PROP_HEALED_ADD:        0,
			FightPropType_FIGHT_PROP_ELEMENT_MASTERY:   0,
			FightPropType_FIGHT_PROP_PHYSICAL_SUB_HURT: 0,
			FightPropType_FIGHT_PROP_PHYSICAL_ADD_HURT: 0,
			FightPropType_FIGHT_PROP_FIRE_ADD_HURT:     0,
			FightPropType_FIGHT_PROP_ELEC_ADD_HURT:     0,
			FightPropType_FIGHT_PROP_WATER_ADD_HURT:    0,
			FightPropType_FIGHT_PROP_GRASS_ADD_HURT:    0,
			FightPropType_FIGHT_PROP_WIND_ADD_HURT:     0,
			FightPropType_FIGHT_PROP_ROCK_ADD_HURT:     0,
			FightPropType_FIGHT_PROP_ICE_ADD_HURT:      0,
			FightPropType_FIGHT_PROP_FIRE_SUB_HURT:     0,
			FightPropType_FIGHT_PROP_ELEC_SUB_HURT:     0,
			FightPropType_FIGHT_PROP_WATER_SUB_HURT:    0,
			FightPropType_FIGHT_PROP_GRASS_SUB_HURT:    0,
			FightPropType_FIGHT_PROP_WIND_SUB_HURT:     0,
			FightPropType_FIGHT_PROP_ROCK_SUB_HURT:     0,
			FightPropType_FIGHT_PROP_ICE_SUB_HURT:      0,
			FightPropType_FIGHT_PROP_CUR_HP:            911.791,
			FightPropType_FIGHT_PROP_MAX_HP:            911.791,
			FightPropType_FIGHT_PROP_CUR_ATTACK:        41.053,
			FightPropType_FIGHT_PROP_CUR_DEFENSE:       57.225,
			FightPropType_FIGHT_PROP_CUR_SPEED:         0,
		}.ToFightPropMap()
		avatar.FetterInfo = &pb.AvatarFetterInfo{ExpLevel: 1}
		notify.AvatarList = append(notify.AvatarList, avatar)
	}
	notify.AvatarTeamMap = make(map[uint32]*pb.AvatarTeam)
	for k, v := range avatar.GetTeamMap() {
		notify.AvatarTeamMap[k] = &pb.AvatarTeam{
			AvatarGuidList: v.GetAvatarGuidList(),
			TeamName:       v.GetTeamName(),
		}
	}
	notify.CurAvatarTeamId = avatar.GetCurTeamId()
	notify.ChooseAvatarGuid = avatar.GetChooseAvatarGuid()
	notify.OwnedFlycloakList = avatar.GetOwnedFlycloakList()
	notify.OwnedCostumeList = avatar.GetOwnedCostumeIdList()
	return s.Send(ctx, &notify)
}

func (s *Server) HandleAvatarUpgradeReq(ctx *Context, req *pb.AvatarUpgradeReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarUpgradeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarDieAnimationEndReq(ctx *Context, req *pb.AvatarDieAnimationEndReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarDieAnimationEndRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarChangeElementTypeReq(ctx *Context, req *pb.AvatarChangeElementTypeReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarChangeElementTypeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarFetterDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarExpeditionDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarExpeditionAllDataReq(ctx *Context, req *pb.AvatarExpeditionAllDataReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarExpeditionAllDataRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarExpeditionStartReq(ctx *Context, req *pb.AvatarExpeditionStartReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarExpeditionStartRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarExpeditionCallBackReq(ctx *Context, req *pb.AvatarExpeditionCallBackReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarExpeditionCallBackRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarExpeditionGetRewardReq(ctx *Context, req *pb.AvatarExpeditionGetRewardReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarExpeditionGetRewardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleChangeMpTeamAvatarReq(ctx *Context, req *pb.ChangeMpTeamAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendChangeMpTeamAvatarRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleChangeTeamNameReq(ctx *Context, req *pb.ChangeTeamNameReq) error {
	panic("not implement")
}

func (s *Server) SendChangeTeamNameRsp(ctx *Context) error {
	panic("not implement")
}

// send SceneTeamUpdateNotify
//
//	flow:
//		SEND ··> SceneTeamUpdateNotify
func (s *Server) SendSceneTeamUpdateNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	avatar := player.Avatar().GetAvatarList()[0]
	var notify pb.SceneTeamUpdateNotify
	var abilityEmbryoList []*pb.AbilityEmbryo
	if avatar.GetAvatarId() == 10000005 {
		abilityEmbryoList = []*pb.AbilityEmbryo{
			{AbilityId: 1, AbilityNameHash: 2957764605, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 2, AbilityNameHash: 1410219662, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 3, AbilityNameHash: 1474894886, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 4, AbilityNameHash: 937205334, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 5, AbilityNameHash: 1771196189, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 6, AbilityNameHash: 2306062007, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 7, AbilityNameHash: 3105629177, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 8, AbilityNameHash: 3771526669, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 9, AbilityNameHash: 100636247, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 10, AbilityNameHash: 1564404322, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 11, AbilityNameHash: 497711942, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 12, AbilityNameHash: 3531639848, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 13, AbilityNameHash: 4255783285, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 14, AbilityNameHash: 1042696700, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 15, AbilityNameHash: 825255509, AbilityOverrideNameHash: 1178079449},
		}
	} else if avatar.GetAvatarId() == 10000007 {
		abilityEmbryoList = []*pb.AbilityEmbryo{
			{AbilityId: 1, AbilityNameHash: 4291357363, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 2, AbilityNameHash: 1410219662, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 3, AbilityNameHash: 1474894886, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 4, AbilityNameHash: 3832178184, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 5, AbilityNameHash: 1771196189, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 6, AbilityNameHash: 2306062007, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 7, AbilityNameHash: 3105629177, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 8, AbilityNameHash: 3771526669, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 9, AbilityNameHash: 100636247, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 10, AbilityNameHash: 1564404322, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 11, AbilityNameHash: 497711942, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 12, AbilityNameHash: 3531639848, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 13, AbilityNameHash: 4255783285, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 14, AbilityNameHash: 3374327026, AbilityOverrideNameHash: 1178079449},
			{AbilityId: 15, AbilityNameHash: 825255509, AbilityOverrideNameHash: 1178079449},
		}
	}
	notify.SceneTeamAvatarList = []*pb.SceneTeamAvatar{{
		SceneId:             3,
		PlayerUid:           uint32(player.ID),
		IsPlayerCurAvatar:   false,
		AvatarGuid:          avatar.GetGuid(),
		EntityId:            uint32(pb.ProtEntityType_PROT_ENTITY_AVATAR)<<24 | 1,
		AvatarAbilityInfo:   &pb.AbilitySyncStateInfo{},
		WeaponGuid:          uint64(player.ID)<<32 | 2,
		WeaponEntityId:      uint32(pb.ProtEntityType_PROT_ENTITY_WEAPON)<<24 | 1,
		WeaponAbilityInfo:   &pb.AbilitySyncStateInfo{},
		SceneEntityInfo:     getAvatarEntityInfo(player, avatar),
		AbilityControlBlock: &pb.AbilityControlBlock{AbilityEmbryoList: abilityEmbryoList},
	}}
	return s.Send(ctx, &notify)
}

func (s *Server) HandleFocusAvatarReq(ctx *Context, req *pb.FocusAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendFocusAvatarRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarSatiationDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarWearFlycloakReq(ctx *Context, req *pb.AvatarWearFlycloakReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarWearFlycloakRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarFlycloakChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarGainFlycloakNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarEquipAffixStartNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarFetterLevelRewardReq(ctx *Context, req *pb.AvatarFetterLevelRewardReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarFetterLevelRewardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAddNoGachaAvatarCardNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarPromoteGetRewardReq(ctx *Context, req *pb.AvatarPromoteGetRewardReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarPromoteGetRewardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarChangeCostumeReq(ctx *Context, req *pb.AvatarChangeCostumeReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarChangeCostumeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarChangeCostumeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarGainCostumeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAvatarChangeAnimHashReq(ctx *Context, req *pb.AvatarChangeAnimHashReq) error {
	panic("not implement")
}

func (s *Server) SendAvatarChangeAnimHashRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePersistentDungeonSwitchAvatarReq(ctx *Context, req *pb.PersistentDungeonSwitchAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendPersistentDungeonSwitchAvatarRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAddBackupAvatarTeamReq(ctx *Context, req *pb.AddBackupAvatarTeamReq) error {
	panic("not implement")
}

func (s *Server) SendAddBackupAvatarTeamRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleDelBackupAvatarTeamReq(ctx *Context, req *pb.DelBackupAvatarTeamReq) error {
	panic("not implement")
}

func (s *Server) SendDelBackupAvatarTeamRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarTeamAllDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarRenameInfoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleItemRenameAvatarReq(ctx *Context, req *pb.ItemRenameAvatarReq) error {
	panic("not implement")
}

func (s *Server) SendItemRenameAvatarRsp(ctx *Context) error {
	panic("not implement")
}
