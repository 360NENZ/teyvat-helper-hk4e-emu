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
//		*SEND ··> AvatarDataNotify
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
		avatar.PropMap = map[uint32]*pb.PropValue{
			1001: {Type: 1001, Value: &pb.PropValue_Ival{Ival: 0}},
			1002: {Type: 1002, Value: &pb.PropValue_Ival{Ival: 0}},
			1003: {Type: 1003, Value: &pb.PropValue_Ival{Ival: 0}},
			1004: {Type: 1004, Value: &pb.PropValue_Ival{Ival: 0}},
			4001: {Type: 4001, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		}
		avatar.FightPropMap = map[uint32]float32{
			1:    911.791,
			4:    41.053,
			6:    0,
			7:    57.225,
			20:   0.05,
			21:   0,
			22:   0.5,
			23:   1,
			26:   0,
			27:   0,
			28:   0,
			29:   0,
			30:   0,
			40:   0,
			41:   0,
			42:   0,
			43:   0,
			44:   0,
			45:   0,
			46:   0,
			50:   0,
			51:   0,
			52:   0,
			53:   0,
			54:   0,
			55:   0,
			56:   0,
			1010: 911.791,
			2000: 911.791,
			2001: 41.053,
			2002: 57.225,
			2003: 0,
		}
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
//		*SEND ··> SceneTeamUpdateNotify
func (s *Server) SendSceneTeamUpdateNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	avatar := player.Avatar().GetAvatarList()[0]
	var notify pb.SceneTeamUpdateNotify
	var abilityEmbryoList []*pb.AbilityEmbryo
	if avatar.GetAvatarId() == 10000005 {
		abilityEmbryoList = []*pb.AbilityEmbryo{
			{AbilityId: 1, AbilityNameHash: 2957764605, AbilityOverrideNameHash: 1178079449},  // Avatar_PlayerBoy_NormalAttack_DamageHandler
			{AbilityId: 2, AbilityNameHash: 1410219662, AbilityOverrideNameHash: 1178079449},  // Avatar_Player_FlyingBomber
			{AbilityId: 3, AbilityNameHash: 1474894886, AbilityOverrideNameHash: 1178079449},  // Avatar_Player_CamCtrl
			{AbilityId: 4, AbilityNameHash: 937205334, AbilityOverrideNameHash: 1178079449},   // Avatar_PlayerBoy_FallingAnthem
			{AbilityId: 5, AbilityNameHash: 1771196189, AbilityOverrideNameHash: 1178079449},  // *GrapplingHookSkill_Ability
			{AbilityId: 6, AbilityNameHash: 2306062007, AbilityOverrideNameHash: 1178079449},  // *Avatar_DefaultAbility_VisionReplaceDieInvincible
			{AbilityId: 7, AbilityNameHash: 3105629177, AbilityOverrideNameHash: 1178079449},  // *Avatar_DefaultAbility_AvartarInShaderChange
			{AbilityId: 8, AbilityNameHash: 3771526669, AbilityOverrideNameHash: 1178079449},  // *Avatar_SprintBS_Invincible
			{AbilityId: 9, AbilityNameHash: 100636247, AbilityOverrideNameHash: 1178079449},   // *Avatar_Freeze_Duration_Reducer
			{AbilityId: 10, AbilityNameHash: 1564404322, AbilityOverrideNameHash: 1178079449}, // *Avatar_Attack_ReviveEnergy
			{AbilityId: 11, AbilityNameHash: 497711942, AbilityOverrideNameHash: 1178079449},  // *Avatar_Component_Initializer
			{AbilityId: 12, AbilityNameHash: 3531639848, AbilityOverrideNameHash: 1178079449}, // *Avatar_HDMesh_Controller
			{AbilityId: 13, AbilityNameHash: 4255783285, AbilityOverrideNameHash: 1178079449}, // *Avatar_Trampoline_Jump_Controller
			{AbilityId: 14, AbilityNameHash: 1042696700, AbilityOverrideNameHash: 1178079449}, // Avatar_PlayerBoy_ExtraAttack_Common
			{AbilityId: 15, AbilityNameHash: 825255509, AbilityOverrideNameHash: 1178079449},  // *Avatar_FallAnthem_Achievement_Listener
		}
	} else if avatar.GetAvatarId() == 10000007 {
		abilityEmbryoList = []*pb.AbilityEmbryo{
			{AbilityId: 1, AbilityNameHash: 4291357363, AbilityOverrideNameHash: 1178079449},  // Avatar_PlayerGirl_NormalAttack_DamageHandler
			{AbilityId: 2, AbilityNameHash: 1410219662, AbilityOverrideNameHash: 1178079449},  // Avatar_Player_FlyingBomber
			{AbilityId: 3, AbilityNameHash: 1474894886, AbilityOverrideNameHash: 1178079449},  // Avatar_Player_CamCtrl
			{AbilityId: 4, AbilityNameHash: 3832178184, AbilityOverrideNameHash: 1178079449},  // Avatar_PlayerGirl_FallingAnthem
			{AbilityId: 5, AbilityNameHash: 1771196189, AbilityOverrideNameHash: 1178079449},  // *GrapplingHookSkill_Ability
			{AbilityId: 6, AbilityNameHash: 2306062007, AbilityOverrideNameHash: 1178079449},  // *Avatar_DefaultAbility_VisionReplaceDieInvincible
			{AbilityId: 7, AbilityNameHash: 3105629177, AbilityOverrideNameHash: 1178079449},  // *Avatar_DefaultAbility_AvartarInShaderChange
			{AbilityId: 8, AbilityNameHash: 3771526669, AbilityOverrideNameHash: 1178079449},  // *Avatar_SprintBS_Invincible
			{AbilityId: 9, AbilityNameHash: 100636247, AbilityOverrideNameHash: 1178079449},   // *Avatar_Freeze_Duration_Reducer
			{AbilityId: 10, AbilityNameHash: 1564404322, AbilityOverrideNameHash: 1178079449}, // *Avatar_Attack_ReviveEnergy
			{AbilityId: 11, AbilityNameHash: 497711942, AbilityOverrideNameHash: 1178079449},  // *Avatar_Component_Initializer
			{AbilityId: 12, AbilityNameHash: 3531639848, AbilityOverrideNameHash: 1178079449}, // *Avatar_HDMesh_Controller
			{AbilityId: 13, AbilityNameHash: 4255783285, AbilityOverrideNameHash: 1178079449}, // *Avatar_Trampoline_Jump_Controller
			{AbilityId: 14, AbilityNameHash: 3374327026, AbilityOverrideNameHash: 1178079449}, // Avatar_PlayerGirl_ExtraAttack_Common
			{AbilityId: 15, AbilityNameHash: 825255509, AbilityOverrideNameHash: 1178079449},  // *Avatar_FallAnthem_Achievement_Listener
		}
	}
	notify.SceneTeamAvatarList = []*pb.SceneTeamAvatar{{
		SceneId:             3,
		PlayerUid:           uint32(player.ID),
		IsPlayerCurAvatar:   false,
		AvatarGuid:          avatar.GetGuid(),
		EntityId:            1<<24 | 1,
		AvatarAbilityInfo:   &pb.AbilitySyncStateInfo{},
		WeaponGuid:          uint64(player.ID)<<32 | 2, // TODO: WeaponGUID
		WeaponEntityId:      6<<24 | 1,
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
