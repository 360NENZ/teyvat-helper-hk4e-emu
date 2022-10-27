package game

import "github.com/teyvat-helper/hk4e-proto/pb"

// send SceneTeamUpdateNotify
//
//	flow:
//		*SEND ··> SceneTeamUpdateNotify
func (s *Server) SendSceneTeamUpdateNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	avatar := player.Avatar().GetAvatarList()[0]
	var notify pb.SceneTeamUpdateNotify
	notify.SceneTeamAvatarList = []*pb.SceneTeamAvatar{{
		SceneId:           3,
		PlayerUid:         uint32(player.ID),
		IsPlayerCurAvatar: false,
		AvatarGuid:        avatar.GetGuid(),
		EntityId:          1<<24 | 1,
		AvatarAbilityInfo: &pb.AbilitySyncStateInfo{},
		WeaponGuid:        uint64(player.ID)<<32 | 2, // TODO: WeaponGUID
		WeaponEntityId:    6<<24 | 1,
		WeaponAbilityInfo: &pb.AbilitySyncStateInfo{},
		SceneEntityInfo:   getAvatarEntityInfo(player, avatar),
		AbilityControlBlock: &pb.AbilityControlBlock{AbilityEmbryoList: []*pb.AbilityEmbryo{
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
		}},
	}}
	return s.Send(ctx, &notify)
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
