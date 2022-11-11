package game

import "github.com/teyvat-helper/hk4e-proto/pb"

type FightPropMap map[FightPropType]float32

func (x FightPropMap) ToFightPropList() []*pb.FightPropPair {
	out := make([]*pb.FightPropPair, 0, len(x))
	for k, v := range x {
		out = append(out, &pb.FightPropPair{PropType: k, PropValue: v})
	}
	return out
}
func (x FightPropMap) ToFightPropMap() map[uint32]float32 {
	out := make(map[uint32]float32, len(x))
	for k, v := range x {
		out[uint32(k)] = v
	}
	return out
}

type OpenStateMap map[OpenStateType]uint32

func (x OpenStateMap) ToOpenStateMap() map[uint32]uint32 {
	out := make(map[uint32]uint32, len(x))
	for k, v := range x {
		out[uint32(k)] = v
	}
	return out
}

type PropMap map[PropType]int64

func (x PropMap) ToPropList() []*pb.PropPair {
	out := make([]*pb.PropPair, 0, len(x))
	for k, v := range x {
		out = append(out, &pb.PropPair{Type: k, PropValue: &pb.PropValue{Type: k, Value: &pb.PropValue_Ival{Ival: v}, Val: v}})
	}
	return out
}

func (x PropMap) ToPropMap() map[uint32]*pb.PropValue {
	out := make(map[uint32]*pb.PropValue, len(x))
	for k, v := range x {
		out[uint32(k)] = &pb.PropValue{Type: k, Value: &pb.PropValue_Ival{Ival: v}, Val: v}
	}
	return out
}

type FightPropType = uint32

const (
	FightPropType_FIGHT_PROP_NONE                                       FightPropType = 0
	FightPropType_FIGHT_PROP_BASE_HP                                    FightPropType = 1
	FightPropType_FIGHT_PROP_HP                                         FightPropType = 2
	FightPropType_FIGHT_PROP_HP_PERCENT                                 FightPropType = 3
	FightPropType_FIGHT_PROP_BASE_ATTACK                                FightPropType = 4
	FightPropType_FIGHT_PROP_ATTACK                                     FightPropType = 5
	FightPropType_FIGHT_PROP_ATTACK_PERCENT                             FightPropType = 6
	FightPropType_FIGHT_PROP_BASE_DEFENSE                               FightPropType = 7
	FightPropType_FIGHT_PROP_DEFENSE                                    FightPropType = 8
	FightPropType_FIGHT_PROP_DEFENSE_PERCENT                            FightPropType = 9
	FightPropType_FIGHT_PROP_BASE_SPEED                                 FightPropType = 10
	FightPropType_FIGHT_PROP_SPEED_PERCENT                              FightPropType = 11
	FightPropType_FIGHT_PROP_HP_MP_PERCENT                              FightPropType = 12
	FightPropType_FIGHT_PROP_ATTACK_MP_PERCENT                          FightPropType = 13
	FightPropType_FIGHT_PROP_CRITICAL                                   FightPropType = 20
	FightPropType_FIGHT_PROP_ANTI_CRITICAL                              FightPropType = 21
	FightPropType_FIGHT_PROP_CRITICAL_HURT                              FightPropType = 22
	FightPropType_FIGHT_PROP_CHARGE_EFFICIENCY                          FightPropType = 23
	FightPropType_FIGHT_PROP_ADD_HURT                                   FightPropType = 24
	FightPropType_FIGHT_PROP_SUB_HURT                                   FightPropType = 25
	FightPropType_FIGHT_PROP_HEAL_ADD                                   FightPropType = 26
	FightPropType_FIGHT_PROP_HEALED_ADD                                 FightPropType = 27
	FightPropType_FIGHT_PROP_ELEMENT_MASTERY                            FightPropType = 28
	FightPropType_FIGHT_PROP_PHYSICAL_SUB_HURT                          FightPropType = 29
	FightPropType_FIGHT_PROP_PHYSICAL_ADD_HURT                          FightPropType = 30
	FightPropType_FIGHT_PROP_DEFENCE_IGNORE_RATIO                       FightPropType = 31
	FightPropType_FIGHT_PROP_DEFENCE_IGNORE_DELTA                       FightPropType = 32
	FightPropType_FIGHT_PROP_FIRE_ADD_HURT                              FightPropType = 40
	FightPropType_FIGHT_PROP_ELEC_ADD_HURT                              FightPropType = 41
	FightPropType_FIGHT_PROP_WATER_ADD_HURT                             FightPropType = 42
	FightPropType_FIGHT_PROP_GRASS_ADD_HURT                             FightPropType = 43
	FightPropType_FIGHT_PROP_WIND_ADD_HURT                              FightPropType = 44
	FightPropType_FIGHT_PROP_ROCK_ADD_HURT                              FightPropType = 45
	FightPropType_FIGHT_PROP_ICE_ADD_HURT                               FightPropType = 46
	FightPropType_FIGHT_PROP_HIT_HEAD_ADD_HURT                          FightPropType = 47
	FightPropType_FIGHT_PROP_FIRE_SUB_HURT                              FightPropType = 50
	FightPropType_FIGHT_PROP_ELEC_SUB_HURT                              FightPropType = 51
	FightPropType_FIGHT_PROP_WATER_SUB_HURT                             FightPropType = 52
	FightPropType_FIGHT_PROP_GRASS_SUB_HURT                             FightPropType = 53
	FightPropType_FIGHT_PROP_WIND_SUB_HURT                              FightPropType = 54
	FightPropType_FIGHT_PROP_ROCK_SUB_HURT                              FightPropType = 55
	FightPropType_FIGHT_PROP_ICE_SUB_HURT                               FightPropType = 56
	FightPropType_FIGHT_PROP_EFFECT_HIT                                 FightPropType = 60
	FightPropType_FIGHT_PROP_EFFECT_RESIST                              FightPropType = 61
	FightPropType_FIGHT_PROP_FREEZE_RESIST                              FightPropType = 62
	FightPropType_FIGHT_PROP_DIZZY_RESIST                               FightPropType = 64
	FightPropType_FIGHT_PROP_FREEZE_SHORTEN                             FightPropType = 65
	FightPropType_FIGHT_PROP_DIZZY_SHORTEN                              FightPropType = 67
	FightPropType_FIGHT_PROP_MAX_FIRE_ENERGY                            FightPropType = 70
	FightPropType_FIGHT_PROP_MAX_ELEC_ENERGY                            FightPropType = 71
	FightPropType_FIGHT_PROP_MAX_WATER_ENERGY                           FightPropType = 72
	FightPropType_FIGHT_PROP_MAX_GRASS_ENERGY                           FightPropType = 73
	FightPropType_FIGHT_PROP_MAX_WIND_ENERGY                            FightPropType = 74
	FightPropType_FIGHT_PROP_MAX_ICE_ENERGY                             FightPropType = 75
	FightPropType_FIGHT_PROP_MAX_ROCK_ENERGY                            FightPropType = 76
	FightPropType_FIGHT_PROP_SKILL_CD_MINUS_RATIO                       FightPropType = 80
	FightPropType_FIGHT_PROP_SHIELD_COST_MINUS_RATIO                    FightPropType = 81
	FightPropType_FIGHT_PROP_CUR_FIRE_ENERGY                            FightPropType = 1000
	FightPropType_FIGHT_PROP_CUR_ELEC_ENERGY                            FightPropType = 1001
	FightPropType_FIGHT_PROP_CUR_WATER_ENERGY                           FightPropType = 1002
	FightPropType_FIGHT_PROP_CUR_GRASS_ENERGY                           FightPropType = 1003
	FightPropType_FIGHT_PROP_CUR_WIND_ENERGY                            FightPropType = 1004
	FightPropType_FIGHT_PROP_CUR_ICE_ENERGY                             FightPropType = 1005
	FightPropType_FIGHT_PROP_CUR_ROCK_ENERGY                            FightPropType = 1006
	FightPropType_FIGHT_PROP_CUR_HP                                     FightPropType = 1010
	FightPropType_FIGHT_PROP_MAX_HP                                     FightPropType = 2000
	FightPropType_FIGHT_PROP_CUR_ATTACK                                 FightPropType = 2001
	FightPropType_FIGHT_PROP_CUR_DEFENSE                                FightPropType = 2002
	FightPropType_FIGHT_PROP_CUR_SPEED                                  FightPropType = 2003
	FightPropType_FIGHT_PROP_NONEXTRA_ATTACK                            FightPropType = 3000
	FightPropType_FIGHT_PROP_NONEXTRA_DEFENSE                           FightPropType = 3001
	FightPropType_FIGHT_PROP_NONEXTRA_CRITICAL                          FightPropType = 3002
	FightPropType_FIGHT_PROP_NONEXTRA_ANTI_CRITICAL                     FightPropType = 3003
	FightPropType_FIGHT_PROP_NONEXTRA_CRITICAL_HURT                     FightPropType = 3004
	FightPropType_FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY                 FightPropType = 3005
	FightPropType_FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY                   FightPropType = 3006
	FightPropType_FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT                 FightPropType = 3007
	FightPropType_FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT                     FightPropType = 3008
	FightPropType_FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT                     FightPropType = 3009
	FightPropType_FIGHT_PROP_NONEXTRA_WATER_ADD_HURT                    FightPropType = 3010
	FightPropType_FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT                    FightPropType = 3011
	FightPropType_FIGHT_PROP_NONEXTRA_WIND_ADD_HURT                     FightPropType = 3012
	FightPropType_FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT                     FightPropType = 3013
	FightPropType_FIGHT_PROP_NONEXTRA_ICE_ADD_HURT                      FightPropType = 3014
	FightPropType_FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT                     FightPropType = 3015
	FightPropType_FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT                     FightPropType = 3016
	FightPropType_FIGHT_PROP_NONEXTRA_WATER_SUB_HURT                    FightPropType = 3017
	FightPropType_FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT                    FightPropType = 3018
	FightPropType_FIGHT_PROP_NONEXTRA_WIND_SUB_HURT                     FightPropType = 3019
	FightPropType_FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT                     FightPropType = 3020
	FightPropType_FIGHT_PROP_NONEXTRA_ICE_SUB_HURT                      FightPropType = 3021
	FightPropType_FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO              FightPropType = 3022
	FightPropType_FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO           FightPropType = 3023
	FightPropType_FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT                 FightPropType = 3024
	FightPropType_FIGHT_PROP_BASE_ELEM_REACT_CRITICAL                   FightPropType = 3045
	FightPropType_FIGHT_PROP_BASE_ELEM_REACT_CRITICAL_HURT              FightPropType = 3046
	FightPropType_FIGHT_PROP_ELEM_REACT_CRITICAL                        FightPropType = 3025
	FightPropType_FIGHT_PROP_ELEM_REACT_CRITICAL_HURT                   FightPropType = 3026
	FightPropType_FIGHT_PROP_ELEM_REACT_EXPLODE_CRITICAL                FightPropType = 3027
	FightPropType_FIGHT_PROP_ELEM_REACT_EXPLODE_CRITICAL_HURT           FightPropType = 3028
	FightPropType_FIGHT_PROP_ELEM_REACT_SWIRL_CRITICAL                  FightPropType = 3029
	FightPropType_FIGHT_PROP_ELEM_REACT_SWIRL_CRITICAL_HURT             FightPropType = 3030
	FightPropType_FIGHT_PROP_ELEM_REACT_ELECTRIC_CRITICAL               FightPropType = 3031
	FightPropType_FIGHT_PROP_ELEM_REACT_ELECTRIC_CRITICAL_HURT          FightPropType = 3032
	FightPropType_FIGHT_PROP_ELEM_REACT_SCONDUCT_CRITICAL               FightPropType = 3033
	FightPropType_FIGHT_PROP_ELEM_REACT_SCONDUCT_CRITICAL_HURT          FightPropType = 3034
	FightPropType_FIGHT_PROP_ELEM_REACT_BURN_CRITICAL                   FightPropType = 3035
	FightPropType_FIGHT_PROP_ELEM_REACT_BURN_CRITICAL_HURT              FightPropType = 3036
	FightPropType_FIGHT_PROP_ELEM_REACT_FROZENBROKEN_CRITICAL           FightPropType = 3037
	FightPropType_FIGHT_PROP_ELEM_REACT_FROZENBROKEN_CRITICAL_HURT      FightPropType = 3038
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_CRITICAL               FightPropType = 3039
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_CRITICAL_HURT          FightPropType = 3040
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_FIRE_CRITICAL          FightPropType = 3041
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_FIRE_CRITICAL_HURT     FightPropType = 3042
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_ELECTRIC_CRITICAL      FightPropType = 3043
	FightPropType_FIGHT_PROP_ELEM_REACT_OVERGROW_ELECTRIC_CRITICAL_HURT FightPropType = 3044
)

type OpenStateType = uint32

const (
	OpenStateType_OPEN_STATE_NONE                                    OpenStateType = 0
	OpenStateType_OPEN_STATE_PAIMON                                  OpenStateType = 1
	OpenStateType_OPEN_STATE_PAIMON_NAVIGATION                       OpenStateType = 2
	OpenStateType_OPEN_STATE_AVATAR_PROMOTE                          OpenStateType = 3
	OpenStateType_OPEN_STATE_AVATAR_TALENT                           OpenStateType = 4
	OpenStateType_OPEN_STATE_WEAPON_PROMOTE                          OpenStateType = 5
	OpenStateType_OPEN_STATE_WEAPON_AWAKEN                           OpenStateType = 6
	OpenStateType_OPEN_STATE_QUEST_REMIND                            OpenStateType = 7
	OpenStateType_OPEN_STATE_GAME_GUIDE                              OpenStateType = 8
	OpenStateType_OPEN_STATE_COOK                                    OpenStateType = 9
	OpenStateType_OPEN_STATE_WEAPON_UPGRADE                          OpenStateType = 10
	OpenStateType_OPEN_STATE_RELIQUARY_UPGRADE                       OpenStateType = 11
	OpenStateType_OPEN_STATE_RELIQUARY_PROMOTE                       OpenStateType = 12
	OpenStateType_OPEN_STATE_WEAPON_PROMOTE_GUIDE                    OpenStateType = 13
	OpenStateType_OPEN_STATE_WEAPON_CHANGE_GUIDE                     OpenStateType = 14
	OpenStateType_OPEN_STATE_PLAYER_LVUP_GUIDE                       OpenStateType = 15
	OpenStateType_OPEN_STATE_FRESHMAN_GUIDE                          OpenStateType = 16
	OpenStateType_OPEN_STATE_SKIP_FRESHMAN_GUIDE                     OpenStateType = 17
	OpenStateType_OPEN_STATE_GUIDE_MOVE_CAMERA                       OpenStateType = 18
	OpenStateType_OPEN_STATE_GUIDE_SCALE_CAMERA                      OpenStateType = 19
	OpenStateType_OPEN_STATE_GUIDE_KEYBOARD                          OpenStateType = 20
	OpenStateType_OPEN_STATE_GUIDE_MOVE                              OpenStateType = 21
	OpenStateType_OPEN_STATE_GUIDE_JUMP                              OpenStateType = 22
	OpenStateType_OPEN_STATE_GUIDE_SPRINT                            OpenStateType = 23
	OpenStateType_OPEN_STATE_GUIDE_MAP                               OpenStateType = 24
	OpenStateType_OPEN_STATE_GUIDE_ATTACK                            OpenStateType = 25
	OpenStateType_OPEN_STATE_GUIDE_FLY                               OpenStateType = 26
	OpenStateType_OPEN_STATE_GUIDE_TALENT                            OpenStateType = 27
	OpenStateType_OPEN_STATE_GUIDE_RELIC                             OpenStateType = 28
	OpenStateType_OPEN_STATE_GUIDE_RELIC_PROM                        OpenStateType = 29
	OpenStateType_OPEN_STATE_COMBINE                                 OpenStateType = 30
	OpenStateType_OPEN_STATE_GACHA                                   OpenStateType = 31
	OpenStateType_OPEN_STATE_GUIDE_GACHA                             OpenStateType = 32
	OpenStateType_OPEN_STATE_GUIDE_TEAM                              OpenStateType = 33
	OpenStateType_OPEN_STATE_GUIDE_PROUD                             OpenStateType = 34
	OpenStateType_OPEN_STATE_GUIDE_AVATAR_PROMOTE                    OpenStateType = 35
	OpenStateType_OPEN_STATE_GUIDE_ADVENTURE_CARD                    OpenStateType = 36
	OpenStateType_OPEN_STATE_FORGE                                   OpenStateType = 37
	OpenStateType_OPEN_STATE_GUIDE_BAG                               OpenStateType = 38
	OpenStateType_OPEN_STATE_EXPEDITION                              OpenStateType = 39
	OpenStateType_OPEN_STATE_GUIDE_ADVENTURE_DAILYTASK               OpenStateType = 40
	OpenStateType_OPEN_STATE_GUIDE_ADVENTURE_DUNGEON                 OpenStateType = 41
	OpenStateType_OPEN_STATE_TOWER                                   OpenStateType = 42
	OpenStateType_OPEN_STATE_WORLD_STAMINA                           OpenStateType = 43
	OpenStateType_OPEN_STATE_TOWER_FIRST_ENTER                       OpenStateType = 44
	OpenStateType_OPEN_STATE_RESIN                                   OpenStateType = 45
	OpenStateType_OPEN_STATE_LIMIT_REGION_FRESHMEAT                  OpenStateType = 47
	OpenStateType_OPEN_STATE_LIMIT_REGION_GLOBAL                     OpenStateType = 48
	OpenStateType_OPEN_STATE_MULTIPLAYER                             OpenStateType = 49
	OpenStateType_OPEN_STATE_GUIDE_MOUSEPC                           OpenStateType = 50
	OpenStateType_OPEN_STATE_GUIDE_MULTIPLAYER                       OpenStateType = 51
	OpenStateType_OPEN_STATE_GUIDE_DUNGEONREWARD                     OpenStateType = 52
	OpenStateType_OPEN_STATE_GUIDE_BLOSSOM                           OpenStateType = 53
	OpenStateType_OPEN_STATE_AVATAR_FASHION                          OpenStateType = 54
	OpenStateType_OPEN_STATE_PHOTOGRAPH                              OpenStateType = 55
	OpenStateType_OPEN_STATE_GUIDE_KSLQUEST                          OpenStateType = 56
	OpenStateType_OPEN_STATE_PERSONAL_LINE                           OpenStateType = 57
	OpenStateType_OPEN_STATE_GUIDE_PERSONAL_LINE                     OpenStateType = 58
	OpenStateType_OPEN_STATE_GUIDE_APPEARANCE                        OpenStateType = 59
	OpenStateType_OPEN_STATE_GUIDE_PROCESS                           OpenStateType = 60
	OpenStateType_OPEN_STATE_GUIDE_PERSONAL_LINE_KEY                 OpenStateType = 61
	OpenStateType_OPEN_STATE_GUIDE_WIDGET                            OpenStateType = 62
	OpenStateType_OPEN_STATE_GUIDE_ACTIVITY_SKILL_ASTER              OpenStateType = 63
	OpenStateType_OPEN_STATE_GUIDE_COLDCLIMATE                       OpenStateType = 64
	OpenStateType_OPEN_STATE_DERIVATIVE_MALL                         OpenStateType = 65
	OpenStateType_OPEN_STATE_GUIDE_EXITMULTIPLAYER                   OpenStateType = 66
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD           OpenStateType = 67
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_REBUILD         OpenStateType = 68
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_CARD            OpenStateType = 69
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_MONSTER         OpenStateType = 70
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_MISSION_CHECK   OpenStateType = 71
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD_SELECT    OpenStateType = 72
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_CHALLENGE_START OpenStateType = 73
	OpenStateType_OPEN_STATE_GUIDE_CONVERT                           OpenStateType = 74
	OpenStateType_OPEN_STATE_GUIDE_THEATREMACHANICUS_MULTIPLAYER     OpenStateType = 75
	OpenStateType_OPEN_STATE_GUIDE_COOP_TASK                         OpenStateType = 76
	OpenStateType_OPEN_STATE_GUIDE_HOMEWORLD_ADEPTIABODE             OpenStateType = 77
	OpenStateType_OPEN_STATE_GUIDE_HOMEWORLD_DEPLOY                  OpenStateType = 78
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_EQUIP              OpenStateType = 79
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_MP_SOLUTION        OpenStateType = 80
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_POWER              OpenStateType = 81
	OpenStateType_OPEN_STATE_GUIDE_HIDEANDSEEK_SKILL                 OpenStateType = 82
	OpenStateType_OPEN_STATE_GUIDE_HOMEWORLD_MAPLIST                 OpenStateType = 83
	OpenStateType_OPEN_STATE_GUIDE_RELICRESOLVE                      OpenStateType = 84
	OpenStateType_OPEN_STATE_GUIDE_GGUIDE                            OpenStateType = 85
	OpenStateType_OPEN_STATE_GUIDE_GGUIDE_HINT                       OpenStateType = 86
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_EQUIP_V2           OpenStateType = 87
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_MP_SOLUTION_V2     OpenStateType = 88
	OpenStateType_OPEN_STATE_GUIDE_CHANNELLERSLAB_POWER_V2           OpenStateType = 89
	OpenStateType_OPEN_STATE_GUIDE_QUICK_TEAMMEMBERCHANGE            OpenStateType = 90
	OpenStateType_OPEN_STATE_GGUIDE_FIRSTSHOW                        OpenStateType = 91
	OpenStateType_OPEN_STATE_GGUIDE_MAINPAGE_ENTRY_DISAPPEAR         OpenStateType = 92
	OpenStateType_OPEN_STATE_CITY_REPUATION_MENGDE                   OpenStateType = 800
	OpenStateType_OPEN_STATE_CITY_REPUATION_LIYUE                    OpenStateType = 801
	OpenStateType_OPEN_STATE_CITY_REPUATION_UI_HINT                  OpenStateType = 802
	OpenStateType_OPEN_STATE_CITY_REPUATION_INAZUMA                  OpenStateType = 803
	OpenStateType_OPEN_STATE_SHOP_TYPE_MALL                          OpenStateType = 900
	OpenStateType_OPEN_STATE_SHOP_TYPE_RECOMMANDED                   OpenStateType = 901
	OpenStateType_OPEN_STATE_SHOP_TYPE_GENESISCRYSTAL                OpenStateType = 902
	OpenStateType_OPEN_STATE_SHOP_TYPE_GIFTPACKAGE                   OpenStateType = 903
	OpenStateType_OPEN_STATE_SHOP_TYPE_PAIMON                        OpenStateType = 1001
	OpenStateType_OPEN_STATE_SHOP_TYPE_CITY                          OpenStateType = 1002
	OpenStateType_OPEN_STATE_SHOP_TYPE_BLACKSMITH                    OpenStateType = 1003
	OpenStateType_OPEN_STATE_SHOP_TYPE_GROCERY                       OpenStateType = 1004
	OpenStateType_OPEN_STATE_SHOP_TYPE_FOOD                          OpenStateType = 1005
	OpenStateType_OPEN_STATE_SHOP_TYPE_SEA_LAMP                      OpenStateType = 1006
	OpenStateType_OPEN_STATE_SHOP_TYPE_VIRTUAL_SHOP                  OpenStateType = 1007
	OpenStateType_OPEN_STATE_SHOP_TYPE_LIYUE_GROCERY                 OpenStateType = 1008
	OpenStateType_OPEN_STATE_SHOP_TYPE_LIYUE_SOUVENIR                OpenStateType = 1009
	OpenStateType_OPEN_STATE_SHOP_TYPE_LIYUE_RESTAURANT              OpenStateType = 1010
	OpenStateType_OPEN_STATE_SHOP_TYPE_INAZUMA_SOUVENIR              OpenStateType = 1011
	OpenStateType_OPEN_STATE_SHOP_TYPE_NPC_TOMOKI                    OpenStateType = 1012
	OpenStateType_OPEN_STATE_SHOP_TYPE_INAZUMA_SOUVENIR_BLACK_BAR    OpenStateType = 1013
	OpenStateType_OPEN_ADVENTURE_MANUAL                              OpenStateType = 1100
	OpenStateType_OPEN_ADVENTURE_MANUAL_CITY_MENGDE                  OpenStateType = 1101
	OpenStateType_OPEN_ADVENTURE_MANUAL_CITY_LIYUE                   OpenStateType = 1102
	OpenStateType_OPEN_ADVENTURE_MANUAL_MONSTER                      OpenStateType = 1103
	OpenStateType_OPEN_ADVENTURE_MANUAL_BOSS_DUNGEON                 OpenStateType = 1104
	OpenStateType_OPEN_STATE_ACTIVITY_SEALAMP                        OpenStateType = 1200
	OpenStateType_OPEN_STATE_ACTIVITY_SEALAMP_TAB2                   OpenStateType = 1201
	OpenStateType_OPEN_STATE_ACTIVITY_SEALAMP_TAB3                   OpenStateType = 1202
	OpenStateType_OPEN_STATE_BATTLE_PASS                             OpenStateType = 1300
	OpenStateType_OPEN_STATE_BATTLE_PASS_ENTRY                       OpenStateType = 1301
	OpenStateType_OPEN_STATE_ACTIVITY_CRUCIBLE                       OpenStateType = 1400
	OpenStateType_OPEN_STATE_ACTIVITY_NEWBEEBOUNS_OPEN               OpenStateType = 1401
	OpenStateType_OPEN_STATE_ACTIVITY_NEWBEEBOUNS_CLOSE              OpenStateType = 1402
	OpenStateType_OPEN_STATE_ACTIVITY_ENTRY_OPEN                     OpenStateType = 1403
	OpenStateType_OPEN_STATE_MENGDE_INFUSEDCRYSTAL                   OpenStateType = 1404
	OpenStateType_OPEN_STATE_LIYUE_INFUSEDCRYSTAL                    OpenStateType = 1405
	OpenStateType_OPEN_STATE_SNOW_MOUNTAIN_ELDER_TREE                OpenStateType = 1406
	OpenStateType_OPEN_STATE_MIRACLE_RING                            OpenStateType = 1407
	OpenStateType_OPEN_STATE_COOP_LINE                               OpenStateType = 1408
	OpenStateType_OPEN_STATE_INAZUMA_INFUSEDCRYSTAL                  OpenStateType = 1409
	OpenStateType_OPEN_STATE_FISH                                    OpenStateType = 1410
	OpenStateType_OPEN_STATE_GUIDE_SUMO_TEAM_SKILL                   OpenStateType = 1411
	OpenStateType_OPEN_STATE_GUIDE_FISH_RECIPE                       OpenStateType = 1412
	OpenStateType_OPEN_STATE_HOME                                    OpenStateType = 1500
	OpenStateType_OPEN_STATE_ACTIVITY_HOMEWORLD                      OpenStateType = 1501
	OpenStateType_OPEN_STATE_ADEPTIABODE                             OpenStateType = 1502
	OpenStateType_OPEN_STATE_HOME_AVATAR                             OpenStateType = 1503
	OpenStateType_OPEN_STATE_HOME_EDIT                               OpenStateType = 1504
	OpenStateType_OPEN_STATE_HOME_EDIT_TIPS                          OpenStateType = 1505
	OpenStateType_OPEN_STATE_RELIQUARY_DECOMPOSE                     OpenStateType = 1600
	OpenStateType_OPEN_STATE_ACTIVITY_H5                             OpenStateType = 1700
	OpenStateType_OPEN_STATE_ORAIONOKAMI                             OpenStateType = 2000
	OpenStateType_OPEN_STATE_GUIDE_CHESS_MISSION_CHECK               OpenStateType = 2001
	OpenStateType_OPEN_STATE_GUIDE_CHESS_BUILD                       OpenStateType = 2002
	OpenStateType_OPEN_STATE_GUIDE_CHESS_WIND_TOWER_CIRCLE           OpenStateType = 2003
	OpenStateType_OPEN_STATE_GUIDE_CHESS_CARD_SELECT                 OpenStateType = 2004
	OpenStateType_OPEN_STATE_INAZUMA_MAINQUEST_FINISHED              OpenStateType = 2005
	OpenStateType_OPEN_STATE_PAIMON_LVINFO                           OpenStateType = 2100
	OpenStateType_OPEN_STATE_TELEPORT_HUD                            OpenStateType = 2101
	OpenStateType_OPEN_STATE_GUIDE_MAP_UNLOCK                        OpenStateType = 2102
	OpenStateType_OPEN_STATE_GUIDE_PAIMON_LVINFO                     OpenStateType = 2103
	OpenStateType_OPEN_STATE_GUIDE_AMBORTRANSPORT                    OpenStateType = 2104
	OpenStateType_OPEN_STATE_GUIDE_FLY_SECOND                        OpenStateType = 2105
	OpenStateType_OPEN_STATE_GUIDE_KAEYA_CLUE                        OpenStateType = 2106
	OpenStateType_OPEN_STATE_CAPTURE_CODEX                           OpenStateType = 2107
	OpenStateType_OPEN_STATE_ACTIVITY_FISH_OPEN                      OpenStateType = 2200
	OpenStateType_OPEN_STATE_ACTIVITY_FISH_CLOSE                     OpenStateType = 2201
	OpenStateType_OPEN_STATE_GUIDE_ROGUE_MAP                         OpenStateType = 2205
	OpenStateType_OPEN_STATE_GUIDE_ROGUE_RUNE                        OpenStateType = 2206
	OpenStateType_OPEN_STATE_GUIDE_BARTENDER_FORMULA                 OpenStateType = 2210
	OpenStateType_OPEN_STATE_GUIDE_BARTENDER_MIX                     OpenStateType = 2211
	OpenStateType_OPEN_STATE_GUIDE_BARTENDER_CUP                     OpenStateType = 2212
	OpenStateType_OPEN_STATE_GUIDE_MAIL_FAVORITES                    OpenStateType = 2400
	OpenStateType_OPEN_STATE_GUIDE_POTION_CONFIGURE                  OpenStateType = 2401
	OpenStateType_OPEN_STATE_GUIDE_LANV2_FIREWORK                    OpenStateType = 2402
	OpenStateType_OPEN_STATE_LOADINGTIPS_ENKANOMIYA                  OpenStateType = 2403
	OpenStateType_OPEN_STATE_MICHIAE_CASKET                          OpenStateType = 2500
	OpenStateType_OPEN_STATE_MAIL_COLLECT_UNLOCK_RED_POINT           OpenStateType = 2501
	OpenStateType_OPEN_STATE_LUMEN_STONE                             OpenStateType = 2600
	OpenStateType_OPEN_STATE_GUIDE_CRYSTALLINK_BUFF                  OpenStateType = 2601
	OpenStateType_OPEN_STATE_GUIDE_MUSIC_GAME_V3                     OpenStateType = 2700
	OpenStateType_OPEN_STATE_GUIDE_MUSIC_GAME_V3_REAL_TIME_EDIT      OpenStateType = 2701
	OpenStateType_OPEN_STATE_GUIDE_MUSIC_GAME_V3_TIMELINE_EDIT       OpenStateType = 2702
	OpenStateType_OPEN_STATE_GUIDE_MUSIC_GAME_V3_SETTING             OpenStateType = 2703
	OpenStateType_OPEN_STATE_GUIDE_ROBOTGACHA                        OpenStateType = 2704
	OpenStateType_OPEN_STATE_GUIDE_FRAGILE_RESIN                     OpenStateType = 2800
	OpenStateType_OPEN_ADVENTURE_MANUAL_EDUCATION                    OpenStateType = 2801
	OpenStateType_OPEN_STATE_CITY_REPUATION_SUMERU                   OpenStateType = 3000
	OpenStateType_OPEN_STATE_VANASARA                                OpenStateType = 3001
	OpenStateType_OPEN_STATE_SUMERU_INFUSEDCRYSTAL                   OpenStateType = 3002
	OpenStateType_OPEN_STATE_LIMIT_REGION_WITHERED_FOREST            OpenStateType = 3003
	OpenStateType_OPEN_STATE_SHOP_TYPE_SUMERU_SOUVENIR               OpenStateType = 3004
	OpenStateType_OPEN_STATE_SHOP_TYPE_SUMERU_SOUVENIR_BLACK_BAR     OpenStateType = 3005
	OpenStateType_OPEN_STATE_GUIDE_ACTIVITY_SKILL_MUQADAS            OpenStateType = 3006
	OpenStateType_OPEN_STATE_GUIDE_MUQADAS                           OpenStateType = 3007
	OpenStateType_OPEN_STATE_GUIDE_WINDFIELD_SKILL                   OpenStateType = 3100
	OpenStateType_OPEN_STATE_GUIDE_VINTAGE_MARKET_UPGRADE            OpenStateType = 3101
	OpenStateType_OPEN_STATE_GUIDE_VINTAGE_MARKET_SKILL              OpenStateType = 3102
	OpenStateType_OPEN_STATE_GUIDE_VINTAGE_MARKET_STAFF              OpenStateType = 3103
	OpenStateType_OPEN_STATE_GCG_RESOURCE_MANAGEMENT                 OpenStateType = 3200
	OpenStateType_OPEN_STATE_GUIDE_FUNGUSFIGHTER_1                   OpenStateType = 3201
	OpenStateType_OPEN_STATE_GUIDE_FUNGUSFIGHTER_2                   OpenStateType = 3202
	OpenStateType_OPEN_STATE_GUIDE_FUNGUSFIGHTER_3                   OpenStateType = 3203
)

type PropType = uint32

const (
	PropType_PROP_NONE                            PropType = 0
	PropType_PROP_EXP                             PropType = 1001
	PropType_PROP_BREAK_LEVEL                     PropType = 1002
	PropType_PROP_SATIATION_VAL                   PropType = 1003
	PropType_PROP_SATIATION_PENALTY_TIME          PropType = 1004
	PropType_PROP_GEAR_START_VAL                  PropType = 2001
	PropType_PROP_GEAR_STOP_VAL                   PropType = 2002
	PropType_PROP_LEVEL                           PropType = 4001
	PropType_PROP_LAST_CHANGE_AVATAR_TIME         PropType = 10001
	PropType_PROP_MAX_SPRING_VOLUME               PropType = 10002
	PropType_PROP_CUR_SPRING_VOLUME               PropType = 10003
	PropType_PROP_IS_SPRING_AUTO_USE              PropType = 10004
	PropType_PROP_SPRING_AUTO_USE_PERCENT         PropType = 10005
	PropType_PROP_IS_FLYABLE                      PropType = 10006
	PropType_PROP_IS_WEATHER_LOCKED               PropType = 10007
	PropType_PROP_IS_GAME_TIME_LOCKED             PropType = 10008
	PropType_PROP_IS_TRANSFERABLE                 PropType = 10009
	PropType_PROP_MAX_STAMINA                     PropType = 10010
	PropType_PROP_CUR_PERSIST_STAMINA             PropType = 10011
	PropType_PROP_CUR_TEMPORARY_STAMINA           PropType = 10012
	PropType_PROP_PLAYER_LEVEL                    PropType = 10013
	PropType_PROP_PLAYER_EXP                      PropType = 10014
	PropType_PROP_PLAYER_HCOIN                    PropType = 10015
	PropType_PROP_PLAYER_SCOIN                    PropType = 10016
	PropType_PROP_PLAYER_MP_SETTING_TYPE          PropType = 10017
	PropType_PROP_IS_MP_MODE_AVAILABLE            PropType = 10018
	PropType_PROP_PLAYER_WORLD_LEVEL              PropType = 10019
	PropType_PROP_PLAYER_RESIN                    PropType = 10020
	PropType_PROP_PLAYER_WAIT_SUB_HCOIN           PropType = 10022
	PropType_PROP_PLAYER_WAIT_SUB_SCOIN           PropType = 10023
	PropType_PROP_IS_ONLY_MP_WITH_PS_PLAYER       PropType = 10024
	PropType_PROP_PLAYER_MCOIN                    PropType = 10025
	PropType_PROP_PLAYER_WAIT_SUB_MCOIN           PropType = 10026
	PropType_PROP_PLAYER_LEGENDARY_KEY            PropType = 10027
	PropType_PROP_IS_HAS_FIRST_SHARE              PropType = 10028
	PropType_PROP_PLAYER_FORGE_POINT              PropType = 10029
	PropType_PROP_CUR_CLIMATE_METER               PropType = 10035
	PropType_PROP_CUR_CLIMATE_TYPE                PropType = 10036
	PropType_PROP_CUR_CLIMATE_AREA_ID             PropType = 10037
	PropType_PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE   PropType = 10038
	PropType_PROP_PLAYER_WORLD_LEVEL_LIMIT        PropType = 10039
	PropType_PROP_PLAYER_WORLD_LEVEL_ADJUST_CD    PropType = 10040
	PropType_PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM PropType = 10041
	PropType_PROP_PLAYER_HOME_COIN                PropType = 10042
	PropType_PROP_PLAYER_WAIT_SUB_HOME_COIN       PropType = 10043
	PropType_PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP   PropType = 10044
	PropType_PROP_PLAYER_GCG_COIN                 PropType = 10045
	PropType_PROP_PLAYER_WAIT_SUB_GCG_COIN        PropType = 10046
)
