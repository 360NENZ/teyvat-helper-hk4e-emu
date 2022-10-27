package game

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/teyvat-helper/hk4e-proto/pb"
	"google.golang.org/protobuf/proto"
)

// HandleGetPlayerTokenReq handle GetPlayerTokenReq
//
//	flow:
//		*RECV <·· GetPlayerTokenReq
//		*SEND ··> GetPlayerTokenRsp
func (s *Server) HandleGetPlayerTokenReq(ctx *Context, req *pb.GetPlayerTokenReq) error {
	clientSeed, keyID, err := s.handleGetPlayerTokenReq(ctx, req)
	if err != nil {
		log.Printf("[GAME] Failed to handle GetPlayerTokenReq: %v", err)
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
	log.Println(binary.BigEndian.Uint64(seed), binary.LittleEndian.Uint64(seed), hex.EncodeToString(seed))
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
//		*RECV <·· PlayerLoginReq
//		*SEND ··> DoSetPlayerBornDataNotify
//		*SEND ··> PlayerLoginRsp
//
//	flow:
//		*RECV <·· PlayerLoginReq
//		*SEND ··> [PlayerData]
//		*SEND ··> PlayerEnterSceneNotify
//		*SEND ··> PlayerLoginRsp
func (s *Server) HandlePlayerLoginReq(ctx *Context, req *pb.PlayerLoginReq) error {
	if err := s.handlePlayerLoginReq(ctx, req); err != nil {
		log.Printf("[GAME] Failed to handle PlayerLoginReq: %v", err)
		return s.Send(ctx, &pb.PlayerLoginRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendPlayerLoginRsp(ctx)
}
func (s *Server) handlePlayerLoginReq(ctx *Context, req *pb.PlayerLoginReq) error {
	if ctx.Session().GetPlayer().Avatar().GetChooseAvatarGuid() == 0 {
		return s.SendDoSetPlayerBornDataNotify(ctx)
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
	notify.NickName = player.Basic().GetNickname()
	notify.ServerTime = uint64(time.Now().UnixMilli())
	notify.RegionId = 50
	notify.PropMap = map[uint32]*pb.PropValue{
		10004: {Type: 10004, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		10005: {Type: 10005, Value: &pb.PropValue_Ival{Ival: 50}, Val: 50},
		10006: {Type: 10006, Value: &pb.PropValue_Ival{Ival: 0}},
		10007: {Type: 10007, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		10008: {Type: 10008, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		10009: {Type: 10009, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		10010: {Type: 10010, Value: &pb.PropValue_Ival{Ival: 10000}, Val: 10000},
		10011: {Type: 10011, Value: &pb.PropValue_Ival{Ival: 10000}, Val: 10000},
		10012: {Type: 10012, Value: &pb.PropValue_Ival{Ival: 0}},
		10013: {Type: 10013, Value: &pb.PropValue_Ival{Ival: 1}, Val: 1},
		10014: {Type: 10014, Value: &pb.PropValue_Ival{Ival: 0}, Val: 0},
		10015: {Type: 10015, Value: &pb.PropValue_Ival{Ival: 0}, Val: 0},
		10016: {Type: 10016, Value: &pb.PropValue_Ival{Ival: 0}, Val: 0},
		10017: {Type: 10017, Value: &pb.PropValue_Ival{Ival: 2}, Val: 2},
		10019: {Type: 10019, Value: &pb.PropValue_Ival{Ival: 0}},
		10020: {Type: 10020, Value: &pb.PropValue_Ival{Ival: 160}, Val: 160},
		10022: {Type: 10022, Value: &pb.PropValue_Ival{Ival: 0}},
		10023: {Type: 10023, Value: &pb.PropValue_Ival{Ival: 0}},
		10025: {Type: 10025, Value: &pb.PropValue_Ival{Ival: 0}},
		10026: {Type: 10026, Value: &pb.PropValue_Ival{Ival: 0}},
		10027: {Type: 10027, Value: &pb.PropValue_Ival{Ival: 0}},
		10035: {Type: 10035, Value: &pb.PropValue_Ival{Ival: 0}},
		10036: {Type: 10036, Value: &pb.PropValue_Ival{Ival: 0}},
		10037: {Type: 10037, Value: &pb.PropValue_Ival{Ival: 0}},
		10038: {Type: 10038, Value: &pb.PropValue_Ival{Ival: 0}},
		10039: {Type: 10039, Value: &pb.PropValue_Ival{Ival: 0}},
		10040: {Type: 10040, Value: &pb.PropValue_Ival{Ival: 0}},
		10041: {Type: 10041, Value: &pb.PropValue_Ival{Ival: 0}},
		10042: {Type: 10042, Value: &pb.PropValue_Ival{Ival: 0}},
		10043: {Type: 10043, Value: &pb.PropValue_Ival{Ival: 0}},
		10044: {Type: 10044, Value: &pb.PropValue_Ival{Ival: 0}},
	}
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
//		*RECV <·· SetPlayerBornDataReq
//		*SEND ··> [PlayerData]
//		*SEND ··> SetPlayerBornDataRsp
//		*SEND ··> PlayerEnterSceneNotify
func (s *Server) HandleSetPlayerBornDataReq(ctx *Context, req *pb.SetPlayerBornDataReq) error {
	if err := s.handleSetPlayerBornDataReq(ctx, req); err != nil {
		log.Printf("[GAME] Failed to handle SetPlayerBornDataReq: %v", err)
		return s.Send(ctx, &pb.SetPlayerBornDataRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendSetPlayerBornDataRsp(ctx)
}
func (s *Server) handleSetPlayerBornDataReq(ctx *Context, req *pb.SetPlayerBornDataReq) error {
	id, name := req.GetAvatarId(), req.GetNickName()
	if id != 10000005 && id != 10000007 {
		return errors.New("bad avatar id")
	}
	player := ctx.Session().GetPlayer()
	if err := player.Basic().SetPlayerBornData(ctx, id, name); err != nil {
		return err
	}
	if err := player.Avatar().SetPlayerBornData(ctx, id); err != nil {
		return err
	}
	if err := player.Social().SetPlayerBornData(ctx, id); err != nil {
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
//		*SEND ··> DoSetPlayerBornDataNotify
func (s *Server) SendDoSetPlayerBornDataNotify(ctx *Context) error {
	return s.Send(ctx, &pb.DoSetPlayerBornDataNotify{})
}

func (s *Server) SendPlayerPropChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerNameReq(ctx *Context, req *pb.SetPlayerNameReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerNameRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetOpenStateReq(ctx *Context, req *pb.SetOpenStateReq) error {
	panic("not implement")
}

func (s *Server) SendSetOpenStateRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendOpenStateUpdateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendOpenStateChangeNotify(ctx *Context) error {
	panic("not implement")
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
//		*RECV <·· PlayerSetPauseReq
//		*SEND ··> PlayerSetPauseRsp
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
//		*RECV <·· PlayerForceExitReq
//		*SEND ··> PlayerForceExitRsp
func (s *Server) HandlePlayerForceExitReq(ctx *Context, req *pb.PlayerForceExitReq) error {
	if _, err := ctx.Session().UpdatePlayer(ctx); err != nil {
		return err
	}
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
