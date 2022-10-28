package game

import (
	"github.com/teyvat-helper/hk4e-proto/pb"
)

func getAvatarEntityInfo(player *Player, avatar *pb.AvatarBin) *pb.SceneEntityInfo {
	entity := &pb.SceneEntityInfo{
		EntityType: pb.ProtEntityType_PROT_ENTITY_AVATAR,
		EntityId:   uint32(pb.ProtEntityType_PROT_ENTITY_AVATAR)<<24 | 1,
		MotionInfo: &pb.MotionInfo{
			Pos:   player.Scene().GetPos(),
			Rot:   player.Scene().GetRot(),
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
			BornPos:         &pb.Vector{},
			ClientExtraInfo: &pb.EntityClientExtraInfo{SkillAnchorPosition: &pb.Vector{}},
		},
	}
	entity.LifeState = avatar.LifeState
	sceneAvatar := &pb.SceneAvatarInfo{
		Uid:               uint32(player.ID),
		PeerId:            1,
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
					EntityId:            uint32(pb.ProtEntityType_PROT_ENTITY_WEAPON)<<24 | 1,
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
	entity.PropList = PropMap{
		PropType_PROP_EXP:                    0,
		PropType_PROP_BREAK_LEVEL:            0,
		PropType_PROP_SATIATION_VAL:          0,
		PropType_PROP_SATIATION_PENALTY_TIME: 0,
		PropType_PROP_LEVEL:                  1,
	}.ToPropList()
	entity.FightPropList = FightPropMap{
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
	}.ToFightPropList()
	return entity
}
