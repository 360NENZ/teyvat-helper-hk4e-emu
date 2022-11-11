package game

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-proto/pb"
	"google.golang.org/protobuf/proto"
)

// HandleGetPlayerTokenReq handle GetPlayerTokenReq
//
//	flow:
//		RECV <·· GetPlayerTokenReq
//		SEND ··> GetPlayerTokenRsp
func (s *Server) HandleGetPlayerTokenReq(ctx *Context, req *pb.GetPlayerTokenReq) error {
	clientSeed, keyID, err := s.handleGetPlayerTokenReq(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to handle GetPlayerTokenReq")
		return s.Send(ctx, &pb.GetPlayerTokenRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendGetPlayerTokenRsp(ctx, clientSeed, keyID)
}
func (s *Server) handleGetPlayerTokenReq(ctx *Context, req *pb.GetPlayerTokenReq) (uint64, uint32, error) {
	id, err := strconv.ParseInt(req.GetAccountUid(), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	player, err := ctx.Session().TokenLogin(ctx, id, req.GetAccountToken())
	if err != nil {
		return 0, 0, err
	}
	playerData, err := ctx.Session().GetPlayerData(ctx, player.ID)
	if err != nil {
		return 0, 0, err
	}
	var playerDataBin pb.PlayerDataBin
	if err := proto.Unmarshal(playerData.BinData, &playerDataBin); err != nil {
		return 0, 0, err
	}
	player.SetBinary(&playerDataBin)
	ctx.Session().SetPlayer(player)
	seed, err := s.secret.Server.DecryptBase64(req.GetClientRandKey())
	if err != nil {
		return 0, 0, err
	}
	return binary.BigEndian.Uint64(seed), req.GetKeyId(), nil
}
func (s *Server) SendGetPlayerTokenRsp(ctx *Context, clientSeed uint64, keyID uint32) error {
	var resp pb.GetPlayerTokenRsp
	player := ctx.Session().GetPlayer()
	resp.Uid = uint32(player.ID)
	resp.AccountUid = strconv.FormatInt(player.Account.ID, 10)
	resp.Token = player.Account.ComboToken
	serverSeed := s.random.Uint64()
	seed := make([]byte, 8)
	binary.BigEndian.PutUint64(seed, clientSeed^serverSeed)
	resp.ServerRandKey, _ = s.secret.Client[keyID].EncryptBase64(seed)
	resp.Sign, _ = s.secret.Server.SignBase64(seed)
	resp.KeyId = keyID
	defer ctx.Session().SetSeed(serverSeed)
	resp.AccountType = 1
	resp.ChannelId = 1
	resp.PlatformType = 3
	resp.RegPlatform = 3
	resp.CountryCode = "US"
	resp.ClientIpStr = "127.0.0.1" // TODO: get real ip
	data := make([]byte, 16+32)
	rand.Read(data)
	resp.ClientVersionRandomKey = fmt.Sprintf("%03x-%012x", data[:3], data[4:16])
	resp.SecurityCmdBuffer = data[16:]
	return s.Send(ctx, &resp)
}

// HandlePlayerLoginReq handle GetPlayerTokenReq
//
//	flow:
//		RECV <·· PlayerLoginReq
//		SEND ··> DoSetPlayerBornDataNotify
//		SEND ··> PlayerLoginRsp
//
//	flow:
//		RECV <·· PlayerLoginReq
//		SEND ··> [PlayerData]
//		SEND ··> PlayerEnterSceneNotify
//		SEND ··> PlayerLoginRsp
func (s *Server) HandlePlayerLoginReq(ctx *Context, req *pb.PlayerLoginReq) error {
	if err := s.handlePlayerLoginReq(ctx, req); err != nil {
		log.Error().Err(err).Msg("Failed to handle PlayerLoginReq")
		return s.Send(ctx, &pb.PlayerLoginRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendPlayerLoginRsp(ctx)
}
func (s *Server) handlePlayerLoginReq(ctx *Context, req *pb.PlayerLoginReq) error {
	player := ctx.Session().GetPlayer()
	if player.Avatar().GetChooseAvatarGuid() == 0 {
		if player.ID != 1 {
			return s.SendDoSetPlayerBornDataNotify(ctx)
		}
		if err := player.SetPlayerBornData(ctx, 10000005, "admin"); err != nil {
			return err
		}
	}
	if err := s.AsyncSendPlayerData(ctx); err != nil {
		return err
	}
	return s.SendPlayerEnterSceneNotify(ctx)
}
func (s *Server) SendPlayerLoginRsp(ctx *Context) error {
	var resp pb.PlayerLoginRsp
	resp.IsUseAbilityHash = true
	resp.AbilityHashCode = -1205511399
	resp.CountryCode = "US"
	resp.GameBiz = "hk4e_global"
	resp.RegisterCps = "mihoyo"
	return s.Send(ctx, &resp)
}

func (s *Server) HandlePlayerLogoutReq(ctx *Context, req *pb.PlayerLogoutReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerLogoutRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerLogoutNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerDataNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	var notify pb.PlayerDataNotify
	notify.RegionId = 50
	notify.ServerTime = uint64(time.Now().UnixMilli())
	notify.NickName = player.Basic().GetNickname()
	notify.IsFirstLoginToday = true
	notify.PropMap = PropMap{
		PropType_PROP_IS_SPRING_AUTO_USE:              1,
		PropType_PROP_SPRING_AUTO_USE_PERCENT:         50,
		PropType_PROP_IS_FLYABLE:                      0,
		PropType_PROP_IS_WEATHER_LOCKED:               1,
		PropType_PROP_IS_GAME_TIME_LOCKED:             1,
		PropType_PROP_IS_TRANSFERABLE:                 1,
		PropType_PROP_MAX_STAMINA:                     10000,
		PropType_PROP_CUR_PERSIST_STAMINA:             10000,
		PropType_PROP_CUR_TEMPORARY_STAMINA:           0,
		PropType_PROP_PLAYER_LEVEL:                    1,
		PropType_PROP_PLAYER_EXP:                      0,
		PropType_PROP_PLAYER_HCOIN:                    0,
		PropType_PROP_PLAYER_SCOIN:                    0,
		PropType_PROP_PLAYER_MP_SETTING_TYPE:          2,
		PropType_PROP_PLAYER_WORLD_LEVEL:              0,
		PropType_PROP_PLAYER_RESIN:                    160,
		PropType_PROP_PLAYER_WAIT_SUB_HCOIN:           0,
		PropType_PROP_PLAYER_WAIT_SUB_SCOIN:           0,
		PropType_PROP_PLAYER_MCOIN:                    0,
		PropType_PROP_PLAYER_WAIT_SUB_MCOIN:           0,
		PropType_PROP_PLAYER_LEGENDARY_KEY:            0,
		PropType_PROP_CUR_CLIMATE_METER:               0,
		PropType_PROP_CUR_CLIMATE_TYPE:                0,
		PropType_PROP_CUR_CLIMATE_AREA_ID:             0,
		PropType_PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE:   0,
		PropType_PROP_PLAYER_WORLD_LEVEL_LIMIT:        0,
		PropType_PROP_PLAYER_WORLD_LEVEL_ADJUST_CD:    0,
		PropType_PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM: 0,
		PropType_PROP_PLAYER_HOME_COIN:                0,
		PropType_PROP_PLAYER_WAIT_SUB_HOME_COIN:       0,
		PropType_PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP:   0,
	}.ToPropMap()
	return s.Send(ctx, &notify)
}

func (s *Server) HandleChangeGameTimeReq(ctx *Context, req *pb.ChangeGameTimeReq) error {
	panic("not implement")
}

func (s *Server) SendChangeGameTimeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerGameTimeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerPropNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientTriggerEventNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerPropReq(ctx *Context, req *pb.SetPlayerPropReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerPropRsp(ctx *Context) error {
	panic("not implement")
}

// handle SetPlayerBornDataReq
//
//	flow:
//		RECV <·· SetPlayerBornDataReq
//		SEND ··> [PlayerData]
//		SEND ··> SetPlayerBornDataRsp
//		SEND ··> PlayerEnterSceneNotify
func (s *Server) HandleSetPlayerBornDataReq(ctx *Context, req *pb.SetPlayerBornDataReq) error {
	if err := s.handleSetPlayerBornDataReq(ctx, req); err != nil {
		log.Error().Err(err).Msg("Failed to handle SetPlayerBornDataReq")
		return s.Send(ctx, &pb.SetPlayerBornDataRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendSetPlayerBornDataRsp(ctx)
}
func (s *Server) handleSetPlayerBornDataReq(ctx *Context, req *pb.SetPlayerBornDataReq) error {
	player := ctx.Session().GetPlayer()
	id, name := req.GetAvatarId(), req.GetNickName()
	if err := player.SetPlayerBornData(ctx, id, name); err != nil {
		return err
	}
	return s.AsyncSendPlayerData(ctx)
}
func (s *Server) SendSetPlayerBornDataRsp(ctx *Context) error {
	if err := s.Send(ctx, &pb.SetPlayerBornDataRsp{}); err != nil {
		return err
	}
	return s.SendPlayerEnterSceneNotify(ctx)
}

// send DoSetPlayerBornDataNotify
//
//	flow:
//		SEND ··> DoSetPlayerBornDataNotify
func (s *Server) SendDoSetPlayerBornDataNotify(ctx *Context) error {
	return s.Send(ctx, &pb.DoSetPlayerBornDataNotify{})
}

func (s *Server) SendPlayerPropChangeNotify(ctx *Context) error {
	panic("not implement")
}

// handle SetPlayerNameReq
//
//	flow:
//		RECV <·· SetPlayerNameReq
//		SEND ··> SetPlayerNameRsp
func (s *Server) HandleSetPlayerNameReq(ctx *Context, req *pb.SetPlayerNameReq) error {
	player := ctx.Session().GetPlayer()
	name := req.GetNickName()
	if err := player.Basic().SetNickname(ctx, name); err != nil {
		return s.Send(ctx, &pb.SetPlayerNameRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendSetPlayerNameRsp(ctx, name)
}
func (s *Server) SendSetPlayerNameRsp(ctx *Context, name string) error {
	return s.Send(ctx, &pb.SetPlayerNameRsp{NickName: name})
}

// handle SetOpenStateReq
//
//	flow:
//		RECV <·· SetOpenStateReq
//		SEND ··> SetOpenStateRsp
//		SEND ··> OpenStateChangeNotify
func (s *Server) HandleSetOpenStateReq(ctx *Context, req *pb.SetOpenStateReq) error {
	player := ctx.Session().GetPlayer()
	key, value := req.GetKey(), req.GetValue()
	if err := player.Basic().SetOpenState(ctx, key, value); err != nil {
		log.Error().Err(err).Msg("Failed to handle SetOpenStateReq")
		return s.Send(ctx, &pb.SetOpenStateRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendSetOpenStateRsp(ctx, key, value)
}
func (s *Server) SendSetOpenStateRsp(ctx *Context, key, value uint32) error {
	if err := s.Send(ctx, &pb.SetOpenStateRsp{Key: key, Value: value}); err != nil {
		return err
	}
	return s.SendOpenStateChangeNotify(ctx, key, value)
}

// send OpenStateUpdateNotify
//
//	flow:
//		SEND ··> OpenStateUpdateNotify
func (s *Server) SendOpenStateUpdateNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	return s.Send(ctx, &pb.OpenStateUpdateNotify{OpenStateMap: player.Basic().GetOpenStateMap()})
}

// send OpenStateChangeNotify
//
//	flow:
//		SEND ··> OpenStateChangeNotify
func (s *Server) SendOpenStateChangeNotify(ctx *Context, key, value uint32) error {
	return s.Send(ctx, &pb.OpenStateChangeNotify{OpenStateMap: map[uint32]uint32{key: value}})
}

func (s *Server) HandlePlayerCookReq(ctx *Context, req *pb.PlayerCookReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerCookRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePlayerRandomCookReq(ctx *Context, req *pb.PlayerRandomCookReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerRandomCookRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCookDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCookRecipeDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCookGradeDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePlayerCompoundMaterialReq(ctx *Context, req *pb.PlayerCompoundMaterialReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerCompoundMaterialRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakeCompoundOutputReq(ctx *Context, req *pb.TakeCompoundOutputReq) error {
	panic("not implement")
}

func (s *Server) SendTakeCompoundOutputRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCompoundDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetCompoundDataReq(ctx *Context, req *pb.GetCompoundDataReq) error {
	panic("not implement")
}

func (s *Server) SendGetCompoundDataRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerTimeNotify(ctx *Context) error {
	panic("not implement")
}

// handle PlayerSetPauseReq
//
//	flow:
//		RECV <·· PlayerSetPauseReq
//		SEND ··> PlayerSetPauseRsp
func (s *Server) HandlePlayerSetPauseReq(ctx *Context, req *pb.PlayerSetPauseReq) error {
	return s.SendPlayerSetPauseRsp(ctx, req.GetIsPaused())
}
func (s *Server) SendPlayerSetPauseRsp(ctx *Context, paused bool) error {
	return s.Send(ctx, &pb.PlayerSetPauseRsp{})
}

func (s *Server) HandlePlayerSetLanguageReq(ctx *Context, req *pb.PlayerSetLanguageReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerSetLanguageRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDataResVersionNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDailyTaskDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDailyTaskProgressNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDailyTaskScoreRewardNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendWorldOwnerDailyTaskNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAddRandTaskInfoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendRemoveRandTaskInfoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakePlayerLevelRewardReq(ctx *Context, req *pb.TakePlayerLevelRewardReq) error {
	panic("not implement")
}

func (s *Server) SendTakePlayerLevelRewardRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerLevelRewardUpdateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGivingRecordNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGivingRecordChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleItemGivingReq(ctx *Context, req *pb.ItemGivingReq) error {
	panic("not implement")
}

func (s *Server) SendItemGivingRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePlayerCookArgsReq(ctx *Context, req *pb.PlayerCookArgsReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerCookArgsRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerLuaShellNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendServerDisconnectClientNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAntiAddictNotify(ctx *Context) error {
	panic("not implement")
}

// handle PlayerForceExitReq
//
//	flow:
//		RECV <·· PlayerForceExitReq
//		SEND ··> PlayerForceExitRsp
func (s *Server) HandlePlayerForceExitReq(ctx *Context, req *pb.PlayerForceExitReq) error {
	defer s.deletePlayerSession(ctx, ctx.Session())
	return s.SendPlayerForceExitRsp(ctx)
}
func (s *Server) SendPlayerForceExitRsp(ctx *Context) error {
	return s.Send(ctx, &pb.PlayerForceExitRsp{})
}

func (s *Server) SendPlayerInjectFixNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendTaskVarNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientLockGameTimeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetNextResourceInfoReq(ctx *Context, req *pb.GetNextResourceInfoReq) error {
	panic("not implement")
}

func (s *Server) SendGetNextResourceInfoRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleAdjustWorldLevelReq(ctx *Context, req *pb.AdjustWorldLevelReq) error {
	panic("not implement")
}

func (s *Server) SendAdjustWorldLevelRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleDailyTaskFilterCityReq(ctx *Context, req *pb.DailyTaskFilterCityReq) error {
	panic("not implement")
}

func (s *Server) SendDailyTaskFilterCityRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDailyTaskUnlockedCitiesNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendExclusiveRuleNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCompoundUnlockNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetGameplayRecommendationReq(ctx *Context, req *pb.GetGameplayRecommendationReq) error {
	panic("not implement")
}

func (s *Server) SendGetGameplayRecommendationRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleTakeBackGivingItemReq(ctx *Context, req *pb.TakeBackGivingItemReq) error {
	panic("not implement")
}

func (s *Server) SendTakeBackGivingItemRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerNicknameAuditDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerNicknameNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendNicknameAuditConfigNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleReadNicknameAuditReq(ctx *Context, req *pb.ReadNicknameAuditReq) error {
	panic("not implement")
}

func (s *Server) SendReadNicknameAuditRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePlayerCompoundMaterialBoostReq(ctx *Context, req *pb.PlayerCompoundMaterialBoostReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerCompoundMaterialBoostRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerGameTimeByLuaNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerIpRegionNotify(ctx *Context) error {
	panic("not implement")
}
