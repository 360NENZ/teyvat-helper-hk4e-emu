package game

import (
	"github.com/teyvat-helper/hk4e-proto/pb"
)

func getAvatarEntityInfo(player *Player, avatar *pb.AvatarBin) *pb.SceneEntityInfo {
	entity := &pb.SceneEntityInfo{
		EntityType: pb.ProtEntityType_PROT_ENTITY_AVATAR,
		EntityId:   1<<24 | 1,
		MotionInfo: &pb.MotionInfo{
			Pos:   &pb.Vector{X: 2747.562, Y: 194.633, Z: -1719.386},
			Rot:   &pb.Vector{Y: 307},
			Speed: &pb.Vector{},
		},
		AnimatorParaList: []*pb.AnimatorParameterValueInfoPair{{}},
		EntityClientData: &pb.EntityClientData{},
		EntityAuthorityInfo: &pb.EntityAuthorityInfo{
			AbilityInfo:         &pb.AbilitySyncStateInfo{},
			RendererChangedInfo: &pb.EntityRendererChangedInfo{},
			AiInfo: &pb.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  &pb.Vector{},
			},
			BornPos: &pb.Vector{},
			// TODO: Unk2700_KDGMOPELHNE
		},
	}
	entity.LifeState = avatar.LifeState
	sceneAvatar := &pb.SceneAvatarInfo{
		Uid:               uint32(player.ID),
		PeerId:            uint32(player.ID),
		Guid:              avatar.GetGuid(), // AvatarGUID
		AvatarId:          avatar.GetAvatarId(),
		SkillDepotId:      avatar.GetSkillDepotId(),
		SkillLevelMap:     avatar.GetDepotMap()[avatar.GetSkillDepotId()].GetSkillLevelMap(),
		WearingFlycloakId: avatar.GetFlycloakId(),
		BornTime:          avatar.GetBornTime(),
	}
	// sceneAvatar.ExcelInfo = &pb.AvatarExcelInfo{
	// 	PrefabPathHash:           15226614807408239869,
	// 	PrefabPathRemoteHash:     12536773066255941804,
	// 	ControllerPathHash:       17650043375074422160,
	// 	ControllerPathRemoteHash: 10948695505325057441,
	// 	CombatConfigHash:         17235148253593515976,
	// }
	for _, v := range avatar.GetEquipList() {
		sceneAvatar.EquipIdList = append(sceneAvatar.EquipIdList, v.GetItemId())
		switch item := v.GetDetail().(type) {
		case *pb.ItemBin_Equip:
			switch equip := item.Equip.GetDetail().(type) {
			case *pb.EquipBin_Reliquary:
			case *pb.EquipBin_Weapon:
				sceneAvatar.Weapon = &pb.SceneWeaponInfo{
					EntityId:            6<<24 | 1,
					GadgetId:            v.GetItemId() + 50000000,
					ItemId:              v.GetItemId(),
					Guid:                v.GetGuid(), // WeaponGUID
					Level:               equip.Weapon.GetLevel(),
					AbilityInfo:         &pb.AbilitySyncStateInfo{},
					RendererChangedInfo: &pb.EntityRendererChangedInfo{},
				}
			}
		}
	}
	entity.Entity = &pb.SceneEntityInfo_Avatar{Avatar: sceneAvatar}
	entity.PropList = []*pb.PropPair{
		{Type: 1001, PropValue: &pb.PropValue{Type: 1001, Value: &pb.PropValue_Ival{Ival: 0}}},
		{Type: 1002, PropValue: &pb.PropValue{Type: 1002, Value: &pb.PropValue_Ival{Ival: 0}}},
		{Type: 1003, PropValue: &pb.PropValue{Type: 1003, Value: &pb.PropValue_Ival{Ival: 0}}},
		{Type: 1004, PropValue: &pb.PropValue{Type: 1004, Value: &pb.PropValue_Ival{Ival: 0}}},
		{Type: 4001, PropValue: &pb.PropValue{Type: 4001, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1}},
	}
	entity.FightPropList = []*pb.FightPropPair{
		{PropType: 1, PropValue: 911.791},
		{PropType: 4, PropValue: 41.053},
		{PropType: 6, PropValue: 0},
		{PropType: 7, PropValue: 57.225},
		{PropType: 20, PropValue: 0.05},
		{PropType: 21, PropValue: 0},
		{PropType: 22, PropValue: 0.5},
		{PropType: 23, PropValue: 1},
		{PropType: 26, PropValue: 0},
		{PropType: 27, PropValue: 0},
		{PropType: 28, PropValue: 0},
		{PropType: 29, PropValue: 0},
		{PropType: 30, PropValue: 0},
		{PropType: 40, PropValue: 0},
		{PropType: 41, PropValue: 0},
		{PropType: 42, PropValue: 0},
		{PropType: 43, PropValue: 0},
		{PropType: 44, PropValue: 0},
		{PropType: 45, PropValue: 0},
		{PropType: 46, PropValue: 0},
		{PropType: 50, PropValue: 0},
		{PropType: 51, PropValue: 0},
		{PropType: 52, PropValue: 0},
		{PropType: 53, PropValue: 0},
		{PropType: 54, PropValue: 0},
		{PropType: 55, PropValue: 0},
		{PropType: 56, PropValue: 0},
		{PropType: 1010, PropValue: 911.791},
		{PropType: 2000, PropValue: 911.791},
		{PropType: 2001, PropValue: 41.053},
		{PropType: 2002, PropValue: 57.225},
		{PropType: 2003, PropValue: 0},
	}
	return entity
}
