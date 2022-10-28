package game

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-proto/pb"
)

// send PlayerEnterSceneNotify
//
//	flow:
//		SEND ··> PlayerEnterSceneNotify
func (s *Server) SendPlayerEnterSceneNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	scene := player.Scene()
	var notify pb.PlayerEnterSceneNotify
	notify.EnterSceneToken = 1 // TODO: random token
	notify.IsFirstLoginEnterScene = true
	notify.TargetUid = uint32(player.ID)
	notify.SceneTagIdList = []uint32{102, 107, 113, 117, 125, 134, 139, 141, 1091, 1094, 1095, 1099, 1101, 1103, 1105, 1110, 1120, 1122, 1125, 1127, 1129, 1131, 1133, 1135, 1137, 1138, 1140, 1143, 1146, 1165, 1168}
	notify.Pos = scene.GetPos()
	notify.Type = pb.EnterType_ENTER_SELF
	notify.SceneBeginTime = uint64(time.Now().UnixMilli())
	notify.SceneId = scene.GetMyCurSceneId()
	notify.WorldType = 1
	notify.EnterReason = 1
	notify.SceneTransaction = fmt.Sprintf("%d-%d-%d-%d", 3, player.ID, time.Now().Unix(), 10000)
	return s.Send(ctx, &notify)
}

func (s *Server) HandleLeaveSceneReq(ctx *Context, req *pb.LeaveSceneReq) error {
	panic("not implement")
}

func (s *Server) SendLeaveSceneRsp(ctx *Context) error {
	panic("not implement")
}

// handle SceneInitFinishReq
//
//	flow:
//		RECV <·· SceneInitFinishReq
//		SEND ··> [asyncSendSceneData]
//		SEND ··> SceneInitFinishRsp
func (s *Server) HandleSceneInitFinishReq(ctx *Context, req *pb.SceneInitFinishReq) error {
	if err := s.handleSceneInitFinishReq(ctx, req); err != nil {
		log.Error().Err(err).Msg("Failed to handle SceneInitFinishReq")
		return s.Send(ctx, &pb.SceneInitFinishRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendSceneInitFinishRsp(ctx, req.GetEnterSceneToken())
}
func (s *Server) handleSceneInitFinishReq(ctx *Context, req *pb.SceneInitFinishReq) error {
	ctx.SetSceneToken(req.GetEnterSceneToken())
	return s.AsyncSendSceneData(ctx)
}
func (s *Server) SendSceneInitFinishRsp(ctx *Context, token uint32) error {
	return s.Send(ctx, &pb.SceneInitFinishRsp{EnterSceneToken: token})
}

// send SceneEntityAppearNotify
//
//	flow:
//		SEND ··> SceneEntityAppearNotify
func (s *Server) SendSceneEntityAppearNotify(ctx *Context, appearType pb.VisionType, entityList ...*pb.SceneEntityInfo) error {
	var notify pb.SceneEntityAppearNotify
	notify.AppearType = appearType
	notify.EntityList = entityList
	return s.Send(ctx, &notify)
}

// send SceneEntityDisappearNotify
//
//	flow:
//		SEND ··> SceneEntityDisappearNotify
func (s *Server) SendSceneEntityDisappearNotify(ctx *Context, disappearType pb.VisionType, entityList ...uint32) error {
	var notify pb.SceneEntityDisappearNotify
	notify.DisappearType = disappearType
	notify.EntityList = entityList
	return s.Send(ctx, &notify)
}

func (s *Server) HandleSceneEntityMoveReq(ctx *Context, req *pb.SceneEntityMoveReq) error {
	panic("not implement")
}

func (s *Server) SendSceneEntityMoveRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneAvatarStaminaStepReq(ctx *Context, req *pb.SceneAvatarStaminaStepReq) error {
	panic("not implement")
}

func (s *Server) SendSceneAvatarStaminaStepRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneEntityMoveNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendScenePlayerLocationNotify(ctx *Context, notify *pb.ScenePlayerLocationNotify) error {
	panic("not implement")
}

// handle GetScenePointReq
//
//	flow:
//		RECV <·· GetScenePointReq
//		SEND ··> GetScenePointRsp
func (s *Server) HandleGetScenePointReq(ctx *Context, req *pb.GetScenePointReq) error {
	return s.SendGetScenePointRsp(ctx, req.GetSceneId(), req.GetBelongUid())
}
func (s *Server) SendGetScenePointRsp(ctx *Context, sceneId, belongId uint32) error {
	var resp pb.GetScenePointRsp
	resp.SceneId = sceneId
	switch sceneId {
	case 3:
		resp.NotExploredDungeonEntryList = []uint32{1, 2, 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 1001}
		if belongId == 1 {
			for i := 1; i < 79; i++ {
				resp.UnlockAreaList = append(resp.UnlockAreaList, uint32(i))
			}
			resp.UnlockedPointList = []uint32{3, 6, 4, 5, 7, 8, 10, 11, 12, 13, 14, 15, 29, 33, 34, 35, 36, 37, 57, 58, 59, 60, 61, 69, 72, 73, 74, 75, 76, 77, 78, 79, 81, 82, 84, 85, 86, 87, 91, 92, 93, 95, 96, 97, 99, 100, 103, 104, 105, 114, 116, 121, 122, 127, 128, 152, 153, 154, 155, 156, 162, 165, 166, 167, 168, 180, 181, 182, 197, 200, 204, 205, 206, 208, 209, 210, 211, 212, 213, 214, 222, 225, 228, 234, 235, 236, 241, 242, 244, 245, 246, 296, 298, 301, 302, 304, 305, 306, 331, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 293, 307, 314, 318, 319, 320, 321, 322, 316, 324, 325, 326, 327, 328, 329, 330, 332, 365, 366, 380, 381, 382, 383, 384, 432, 435, 462, 463, 464, 465, 471, 472, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 575, 586, 587, 588, 589, 598, 604, 605, 606, 615, 616, 676, 442, 443, 444, 445, 446, 438, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 555, 557, 558, 626, 653, 629, 517, 518, 519, 577, 599, 602, 603, 612, 623, 625, 652, 655, 665, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 666, 667, 668, 691, 702, 703, 706}
		}
	case 5, 6, 7:
	case 9:
		resp.UnlockAreaList = []uint32{403}
	}
	return s.Send(ctx, &resp)
}

func (s *Server) SendEnterTransPointRegionNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendExitTransPointRegionNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendScenePointUnlockNotify(ctx *Context) error {
	return s.Send(ctx, &pb.ScenePointUnlockNotify{SceneId: 3, PointList: []uint32{}})
}

func (s *Server) HandleSceneTransToPointReq(ctx *Context, req *pb.SceneTransToPointReq) error {
	panic("not implement")
}

func (s *Server) SendSceneTransToPointRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEntityJumpNotify(ctx *Context) error {
	panic("not implement")
}

// handle GetSceneAreaReq
//
//	flow:
//		RECV <·· GetSceneAreaReq
//		SEND ··> GetSceneAreaRsp
func (s *Server) HandleGetSceneAreaReq(ctx *Context, req *pb.GetSceneAreaReq) error {
	return s.SendGetSceneAreaRsp(ctx, req.GetSceneId(), req.GetBelongUid())
}
func (s *Server) SendGetSceneAreaRsp(ctx *Context, sceneId, belongId uint32) error {
	var resp pb.GetSceneAreaRsp
	resp.SceneId = sceneId
	switch sceneId {
	case 3:
		resp.CityInfoList = []*pb.CityInfo{
			{CityId: 1, Level: 1},
			{CityId: 2, Level: 1},
			{CityId: 3, Level: 1},
			{CityId: 4, Level: 1},
			{CityId: 99, Level: 1},
			{CityId: 100, Level: 1},
			{CityId: 101, Level: 1},
			{CityId: 102, Level: 1},
		}
		if belongId == 1 {
			for i := 1; i < 79; i++ {
				resp.AreaIdList = append(resp.AreaIdList, uint32(i))
			}
		}
	case 5, 6, 7:
	case 9:
		resp.CityInfoList = []*pb.CityInfo{
			{CityId: 1, Level: 1},
			{CityId: 2, Level: 1},
			{CityId: 3, Level: 1},
			{CityId: 4, Level: 1},
			{CityId: 99, Level: 1},
			{CityId: 100, Level: 1},
			{CityId: 101, Level: 1},
			{CityId: 102, Level: 1},
		}
	}
	return s.Send(ctx, &resp)
}

func (s *Server) SendSceneAreaUnlockNotify(ctx *Context) error {
	return s.Send(ctx, &pb.SceneAreaUnlockNotify{SceneId: 3, AreaList: []uint32{}})
}

func (s *Server) HandleSceneEntityDrownReq(ctx *Context, req *pb.SceneEntityDrownReq) error {
	panic("not implement")
}

func (s *Server) SendSceneEntityDrownRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneCreateEntityReq(ctx *Context, req *pb.SceneCreateEntityReq) error {
	panic("not implement")
}

func (s *Server) SendSceneCreateEntityRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneDestroyEntityReq(ctx *Context, req *pb.SceneDestroyEntityReq) error {
	panic("not implement")
}

func (s *Server) SendSceneDestroyEntityRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneForceUnlockNotify(ctx *Context) error {
	return s.Send(ctx, &pb.SceneForceUnlockNotify{})
}

func (s *Server) SendSceneForceLockNotify(ctx *Context) error {
	panic("not implement")
}

// HandleEnterWorldAreaReq handle EnterWorldAreaReq
//
//	flow:
//		RECV <·· EnterWorldAreaReq
//		SEND ··> EnterWorldAreaRsp
func (s *Server) HandleEnterWorldAreaReq(ctx *Context, req *pb.EnterWorldAreaReq) error {
	return s.SendEnterWorldAreaRsp(ctx, req.GetAreaType(), req.GetAreaId())
}
func (s *Server) SendEnterWorldAreaRsp(ctx *Context, typ, id uint32) error {
	return s.Send(ctx, &pb.EnterWorldAreaRsp{AreaType: typ, AreaId: id})
}

func (s *Server) HandleEntityForceSyncReq(ctx *Context, req *pb.EntityForceSyncReq) error {
	panic("not implement")
}

func (s *Server) SendEntityForceSyncRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetAreaExplorePointReq(ctx *Context, req *pb.GetAreaExplorePointReq) error {
	panic("not implement")
}

func (s *Server) SendGetAreaExplorePointRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleClientTransmitReq(ctx *Context, req *pb.ClientTransmitReq) error {
	panic("not implement")
}

func (s *Server) SendClientTransmitRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEnterSceneWeatherAreaNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendExitSceneWeatherAreaNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneAreaWeatherNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendScenePlayerInfoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendWorldPlayerLocationNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendBeginCameraSceneLookNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEndCameraSceneLookNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendMarkEntityInMinMapNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUnmarkEntityInMinMapNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleExecuteGroupTriggerReq(ctx *Context, req *pb.ExecuteGroupTriggerReq) error {
	panic("not implement")
}

func (s *Server) SendExecuteGroupTriggerRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleLevelupCityReq(ctx *Context, req *pb.LevelupCityReq) error {
	panic("not implement")
}

func (s *Server) SendLevelupCityRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneRouteChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlatformStartRouteNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlatformStopRouteNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlatformChangeRouteNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendScenePlayerSoundNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandlePersonalSceneJumpReq(ctx *Context, req *pb.PersonalSceneJumpReq) error {
	panic("not implement")
}

func (s *Server) SendPersonalSceneJumpRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSealBattleBeginNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSealBattleEndNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSealBattleProgressNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientPauseNotify(ctx *Context) error {
	panic("not implement")
}

// send PlayerEnterSceneInfoNotify
//
//	flow:
//		SEND ··> PlayerEnterSceneInfoNotify
func (s *Server) SendPlayerEnterSceneInfoNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	var notify pb.PlayerEnterSceneInfoNotify
	notify.MpLevelEntityInfo = &pb.MPLevelEntityInfo{
		AuthorityPeerId: uint32(player.ID),
		EntityId:        uint32(pb.ProtEntityType_PROT_ENTITY_MP_LEVEL)<<24 | 1,
		AbilityInfo:     &pb.AbilitySyncStateInfo{},
	}
	notify.CurAvatarEntityId = uint32(pb.ProtEntityType_PROT_ENTITY_AVATAR)<<24 | 1
	notify.AvatarEnterInfo = []*pb.AvatarEnterSceneInfo{{
		AvatarGuid:        uint64(player.ID)<<32 | 1, // AvatarGUID
		AvatarEntityId:    uint32(pb.ProtEntityType_PROT_ENTITY_AVATAR)<<24 | 1,
		AvatarAbilityInfo: &pb.AbilitySyncStateInfo{},
		WeaponGuid:        uint64(player.ID)<<32 | 2, // WeaponGUID
		WeaponEntityId:    uint32(pb.ProtEntityType_PROT_ENTITY_WEAPON)<<24 | 1,
		WeaponAbilityInfo: &pb.AbilitySyncStateInfo{},
	}}
	notify.TeamEnterInfo = &pb.TeamEnterSceneInfo{
		TeamEntityId:        uint32(pb.ProtEntityType_PROT_ENTITY_TEAM)<<24 | 1,
		TeamAbilityInfo:     &pb.AbilitySyncStateInfo{},
		AbilityControlBlock: &pb.AbilityControlBlock{},
	}
	notify.EnterSceneToken = ctx.GetSceneToken()
	return s.Send(ctx, &notify)
}

func (s *Server) HandleJoinPlayerSceneReq(ctx *Context, req *pb.JoinPlayerSceneReq) error {
	panic("not implement")
}

func (s *Server) SendJoinPlayerSceneRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneKickPlayerReq(ctx *Context, req *pb.SceneKickPlayerReq) error {
	panic("not implement")
}

func (s *Server) SendSceneKickPlayerRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneKickPlayerNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendHitClientTrivialNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleBackMyWorldReq(ctx *Context, req *pb.BackMyWorldReq) error {
	panic("not implement")
}

func (s *Server) SendBackMyWorldRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSeeMonsterReq(ctx *Context, req *pb.SeeMonsterReq) error {
	panic("not implement")
}

func (s *Server) SendSeeMonsterRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAddSeenMonsterNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAllSeenMonsterNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneTimeNotify(ctx *Context) error {
	panic("not implement")
}

// handle EnterSceneReadyReq
//
//	flow:
//		RECV <·· EnterSceneReadyReq
//		SEND ··> EnterSceneReadyRsp
//		SEND ··> EnterScenePeerNotify
func (s *Server) HandleEnterSceneReadyReq(ctx *Context, req *pb.EnterSceneReadyReq) error {
	return s.SendEnterSceneReadyRsp(ctx, req.GetEnterSceneToken())
}
func (s *Server) SendEnterSceneReadyRsp(ctx *Context, token uint32) error {
	var resp pb.EnterSceneReadyRsp
	resp.EnterSceneToken = token
	if err := s.Send(ctx, &resp); err != nil {
		return err
	}
	return s.SendEnterScenePeerNotify(ctx, token)
}

// send EnterScenePeerNotify
//
//	flow:
//		SEND ··> EnterScenePeerNotify
func (s *Server) SendEnterScenePeerNotify(ctx *Context, token uint32) error {
	var notify pb.EnterScenePeerNotify
	notify.PeerId = 1
	notify.EnterSceneToken = token
	notify.DestSceneId = 3
	notify.HostPeerId = 1
	return s.Send(ctx, &notify)
}

// handle EnterSceneDoneReq
//
//	flow:
//		RECV <·· EnterSceneDoneReq
//		SEND ··> SceneAreaWeatherNotify
//		SEND ··> SceneEntityAppearNotify
//		SEND ··> SceneEntityAppearNotify
//		SEND ··> PlatformStartRouteNotify
//		SEND ··> PlatformStartRouteNotify
//		SEND ··> PlayerEyePointStateNotify
//		SEND ··> EnterSceneDoneRsp
//		SEND ··> GroupSuiteNotify
func (s *Server) HandleEnterSceneDoneReq(ctx *Context, req *pb.EnterSceneDoneReq) error {
	if err := s.handleEnterSceneDoneReq(ctx, req); err != nil {
		log.Error().Err(err).Msg("Failed to handle EnterSceneDoneReq")
		return s.Send(ctx, &pb.EnterSceneDoneRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendEnterSceneDoneRsp(ctx, req.GetEnterSceneToken())
}
func (s *Server) handleEnterSceneDoneReq(ctx *Context, req *pb.EnterSceneDoneReq) error {
	player := ctx.Session().GetPlayer()
	avatar := player.Avatar().GetAvatarList()[0]
	return s.SendSceneEntityAppearNotify(ctx, pb.VisionType_VISION_BORN, getAvatarEntityInfo(player, avatar))
}
func (s *Server) SendEnterSceneDoneRsp(ctx *Context, token uint32) error {
	return s.Send(ctx, &pb.EnterSceneDoneRsp{EnterSceneToken: token})
}

func (s *Server) SendWorldPlayerDieNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWorldPlayerReviveReq(ctx *Context, req *pb.WorldPlayerReviveReq) error {
	panic("not implement")
}

func (s *Server) SendWorldPlayerReviveRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendJoinPlayerFailNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSetSceneWeatherAreaReq(ctx *Context, req *pb.SetSceneWeatherAreaReq) error {
	panic("not implement")
}

func (s *Server) SendSetSceneWeatherAreaRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleExecuteGadgetLuaReq(ctx *Context, req *pb.ExecuteGadgetLuaReq) error {
	panic("not implement")
}

func (s *Server) SendExecuteGadgetLuaRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCutSceneBeginNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCutSceneFinishNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCutSceneEndNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientScriptEventNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneEntitiesMovesReq(ctx *Context, req *pb.SceneEntitiesMovesReq) error {
	panic("not implement")
}

func (s *Server) SendSceneEntitiesMovesRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneEntitiesMoveCombineNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleUnlockTransPointReq(ctx *Context, req *pb.UnlockTransPointReq) error {
	panic("not implement")
}

func (s *Server) SendUnlockTransPointRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleSceneWeatherForcastReq(ctx *Context, req *pb.SceneWeatherForcastReq) error {
	panic("not implement")
}

func (s *Server) SendSceneWeatherForcastRsp(ctx *Context) error {
	panic("not implement")
}

// handle MarkMapReq
//
//	flow:
//		RECV <·· MarkMapReq
//		SEND ··> MarkMapRsp
func (s *Server) HandleMarkMapReq(ctx *Context, req *pb.MarkMapReq) error {
	return s.SendMarkMapRsp(ctx, req.GetMark().GetPos())
}
func (s *Server) SendMarkMapRsp(ctx *Context, pos *pb.Vector) error {
	if err := s.Send(ctx, &pb.MarkMapRsp{}); err != nil {
		return err
	}
	pos.Y = 512
	ctx.Session().GetPlayer().Scene().SetPos(pos)
	return s.SendPlayerEnterSceneNotify(ctx)
}

func (s *Server) SendAllMarkPointNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendWorldDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEntityMoveRoomNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendWorldPlayerInfoNotify(ctx *Context) error {
	panic("not implement")
}

// handle PostEnterSceneReq
//
//	flow:
//		RECV <·· PostEnterSceneReq
//		SEND ··> ServerCondMeetQuestListUpdateNotify
//		SEND ··> FinishedParentQuestUpdateNotify
//		SEND ··> QuestProgressUpdateNotify
//		SEND ··> QuestListUpdateNotify
//		SEND ··> PostEnterSceneRsp
//		SEND ··> AnchorPointDataNotify
//		SEND ··> OneofGatherPointDetectorDataNotify
//		SEND ··> ChatChannelDataNotify
func (s *Server) HandlePostEnterSceneReq(ctx *Context, req *pb.PostEnterSceneReq) error {
	if err := s.handlePostEnterSceneReq(ctx, req); err != nil {
		log.Error().Err(err).Msg("Failed to handle PostEnterSceneReq")
		return s.Send(ctx, &pb.PostEnterSceneRsp{Retcode: int32(pb.Retcode_RET_FAIL)})
	}
	return s.SendPostEnterSceneRsp(ctx, req.GetEnterSceneToken())
}
func (s *Server) handlePostEnterSceneReq(ctx *Context, req *pb.PostEnterSceneReq) error {
	return nil
}
func (s *Server) SendPostEnterSceneRsp(ctx *Context, token uint32) error {
	return s.Send(ctx, &pb.PostEnterSceneRsp{EnterSceneToken: token})
}

func (s *Server) HandlePlayerChatReq(ctx *Context, req *pb.PlayerChatReq) error {
	panic("not implement")
}

func (s *Server) SendPlayerChatRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerChatNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerChatCDNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendChatHistoryNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDungeonEntryToBeExploreNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetDungeonEntryExploreConditionReq(ctx *Context, req *pb.GetDungeonEntryExploreConditionReq) error {
	panic("not implement")
}

func (s *Server) SendGetDungeonEntryExploreConditionRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUnfreezeGroupLimitNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSetEntityClientDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGroupSuiteNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGroupUnloadNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendMonsterAIConfigHashNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendShowTemplateReminderNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendShowCommonTipsNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCloseCommonTipsNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendChangeWorldToSingleModeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSyncScenePlayTeamEntityNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDelScenePlayTeamEntityNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerEyePointStateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetMapMarkTipsReq(ctx *Context, req *pb.GetMapMarkTipsReq) error {
	panic("not implement")
}

func (s *Server) SendGetMapMarkTipsRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleChangeWorldToSingleModeReq(ctx *Context, req *pb.ChangeWorldToSingleModeReq) error {
	panic("not implement")
}

func (s *Server) SendChangeWorldToSingleModeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetWorldMpInfoReq(ctx *Context, req *pb.GetWorldMpInfoReq) error {
	panic("not implement")
}

func (s *Server) SendGetWorldMpInfoRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEntityConfigHashNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendForceDragAvatarNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendMonsterPointArrayRouteUpdateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendForceDragBackTransferNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetScenePerformanceReq(ctx *Context, req *pb.GetScenePerformanceReq) error {
	panic("not implement")
}

func (s *Server) SendGetScenePerformanceRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneAudioNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendHitTreeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEntityTagChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAvatarFollowRouteNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendSceneEntityUpdateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientHashDebugNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendPlayerWorldSceneInfoListNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendLuaEnvironmentEffectNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientLoadingCostumeVerificationNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendShowClientGuideNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendShowClientTutorialNotify(ctx *Context) error {
	panic("not implement")
}

// handle GetMapAreaReq
//
//	flow:
//		RECV <·· GetMapAreaReq
//		SEND ··> GetMapAreaRsp
func (s *Server) HandleGetMapAreaReq(ctx *Context, req *pb.GetMapAreaReq) error {
	return s.SendGetMapAreaRsp(ctx)
}
func (s *Server) SendGetMapAreaRsp(ctx *Context) error {
	return s.Send(ctx, &pb.GetMapAreaRsp{MapAreaInfoList: []*pb.MapAreaInfo{
		{MapAreaId: 1, IsOpen: true},
		{MapAreaId: 2, IsOpen: true},
		{MapAreaId: 3, IsOpen: true},
		{MapAreaId: 4, IsOpen: true},
		{MapAreaId: 5},
		{MapAreaId: 6, IsOpen: true},
		{MapAreaId: 7},
		{MapAreaId: 8},
		{MapAreaId: 9},
		{MapAreaId: 10, IsOpen: true},
		{MapAreaId: 11, IsOpen: true},
		{MapAreaId: 12, IsOpen: true},
		{MapAreaId: 13, IsOpen: true},
		{MapAreaId: 14, IsOpen: true},
		{MapAreaId: 15, IsOpen: true},
		{MapAreaId: 16},
		{MapAreaId: 17, IsOpen: true},
		{MapAreaId: 18, IsOpen: true},
		{MapAreaId: 20, IsOpen: true},
		{MapAreaId: 21, IsOpen: true},
		{MapAreaId: 22, IsOpen: true},
		{MapAreaId: 23, IsOpen: true},
		{MapAreaId: 24, IsOpen: true},
		{MapAreaId: 25, IsOpen: true},
		{MapAreaId: 30, IsOpen: true},
		{MapAreaId: 31, IsOpen: true},
		{MapAreaId: 32, IsOpen: true},
		{MapAreaId: 33, IsOpen: true},
		{MapAreaId: 36, IsOpen: true},
		{MapAreaId: 37, IsOpen: true},
		{MapAreaId: 38, IsOpen: true},
		{MapAreaId: 39, IsOpen: true},
	}})
}

func (s *Server) SendMapAreaChangeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendLeaveWorldNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGuestBeginEnterSceneNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGuestPostEnterSceneNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendLevelTagDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendStopReminderNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendAreaPlayInfoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleCheckGroupReplacedReq(ctx *Context, req *pb.CheckGroupReplacedReq) error {
	panic("not implement")
}

func (s *Server) SendCheckGroupReplacedRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendDeathZoneObserveNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendWorldChestOpenNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleWidgetQuickHitTreeReq(ctx *Context, req *pb.WidgetQuickHitTreeReq) error {
	panic("not implement")
}

func (s *Server) SendWidgetQuickHitTreeRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendBeginCameraSceneLookWithTemplateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendRefreshEntityAuthNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendScenePlayerBackgroundAvatarRefreshNotify(ctx *Context) error {
	panic("not implement")
}
