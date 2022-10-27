package game

import (
	"fmt"
	"log"
	"time"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

// send PlayerEnterSceneNotify
//
//	flow:
//		*SEND ··> PlayerEnterSceneNotify
//
//	todo:
//		SEND ··> Unk3100_MDGBODAFNDA
//		SEND ··> Unk3000_KEJGDDMMBLP
func (s *Server) SendPlayerEnterSceneNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	var notify pb.PlayerEnterSceneNotify
	notify.EnterSceneToken = 1 // TODO: random token
	notify.IsFirstLoginEnterScene = true
	notify.TargetUid = uint32(player.ID)
	notify.SceneTagIdList = []uint32{102, 107, 113, 117, 125, 134, 139, 141, 1091, 1094, 1095, 1099, 1101, 1103, 1105, 1110, 1120, 1122, 1125, 1127, 1129, 1131, 1133, 1135, 1137, 1138, 1140, 1143, 1146, 1165, 1168}
	notify.Pos = &pb.Vector{X: 2747.562, Y: 194.633, Z: -1719.386}
	notify.Type = pb.EnterType_ENTER_SELF
	notify.SceneBeginTime = uint64(time.Now().UnixMilli())
	notify.SceneId = 3
	notify.WorldType = 1
	notify.EnterReason = 1
	notify.SceneTransaction = fmt.Sprintf("%d-%d-%d-%d", 3, player.ID, time.Now().Unix(), 10000)
	return s.Send(ctx, &notify)
}

// handle SceneInitFinishReq
//
//	flow:
//		*RECV <·· SceneInitFinishReq
//		SEND ··> [asyncSendSceneData]
//		*SEND ··> SceneInitFinishRsp
//
//	todo:
//		SEND ··> SyncTeamEntityNotify
//		SEND ··> AllMarkPointNotify
//		SEND ··> SyncScenePlayTeamEntityNotify
//		SEND ··> ScenePlayBattleInfoListNotify
//		SEND ··> Unk3000_NNPCGEAHNHM
//		SEND ··> Unk2700_HBOFACHAGIF_ServerNotify
func (s *Server) HandleSceneInitFinishReq(ctx *Context, req *pb.SceneInitFinishReq) error {
	if err := s.handleSceneInitFinishReq(ctx, req); err != nil {
		log.Printf("[GAME] Failed to handle SceneInitFinishReq: %v", err)
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
//		*SEND ··> SceneEntityAppearNotify
func (s *Server) SendSceneEntityAppearNotify(ctx *Context, appearType pb.VisionType, entityList ...*pb.SceneEntityInfo) error {
	var notify pb.SceneEntityAppearNotify
	notify.AppearType = appearType
	notify.EntityList = entityList
	return s.Send(ctx, &notify)
}

// send SceneEntityDisappearNotify
//
//	flow:
//		*SEND ··> SceneEntityDisappearNotify
func (s *Server) SendSceneEntityDisappearNotify(ctx *Context, disappearType pb.VisionType, entityList ...uint32) error {
	var notify pb.SceneEntityDisappearNotify
	notify.DisappearType = disappearType
	notify.EntityList = entityList
	return s.Send(ctx, &notify)
}

// send PlayerEnterSceneInfoNotify
//
//	flow:
//		*SEND ··> PlayerEnterSceneInfoNotify
func (s *Server) SendPlayerEnterSceneInfoNotify(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	var notify pb.PlayerEnterSceneInfoNotify
	notify.MpLevelEntityInfo = &pb.MPLevelEntityInfo{
		AuthorityPeerId: 1,
		EntityId:        11<<24 | 1,
		AbilityInfo:     &pb.AbilitySyncStateInfo{},
	}
	notify.CurAvatarEntityId = 1<<24 | 1
	notify.AvatarEnterInfo = []*pb.AvatarEnterSceneInfo{{
		AvatarGuid:        uint64(player.ID)<<32 | 1, // AvatarGUID
		AvatarEntityId:    1<<24 | 1,
		AvatarAbilityInfo: &pb.AbilitySyncStateInfo{},
		WeaponGuid:        uint64(player.ID)<<32 | 2, // WeaponGUID
		WeaponEntityId:    6<<24 | 1,
		WeaponAbilityInfo: &pb.AbilitySyncStateInfo{},
	}}
	notify.TeamEnterInfo = &pb.TeamEnterSceneInfo{
		TeamEntityId:        9<<24 | 1,
		TeamAbilityInfo:     &pb.AbilitySyncStateInfo{},
		AbilityControlBlock: &pb.AbilityControlBlock{},
	}
	notify.EnterSceneToken = ctx.GetSceneToken()
	return s.Send(ctx, &notify)
}

// handle EnterSceneReadyReq
//
//	flow:
//		*RECV <·· EnterSceneReadyReq
//		*SEND ··> EnterSceneReadyRsp
//		*SEND ··> EnterScenePeerNotify
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
//		*SEND ··> EnterScenePeerNotify
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
//		*RECV <·· EnterSceneDoneReq
//		SEND ··> SceneAreaWeatherNotify
//		*SEND ··> SceneEntityAppearNotify
//		*SEND ··> SceneEntityAppearNotify
//		SEND ··> PlatformStartRouteNotify
//		SEND ··> PlatformStartRouteNotify
//		SEND ··> PlayerEyePointStateNotify
//		*SEND ··> EnterSceneDoneRsp
//		SEND ··> GroupSuiteNotify
func (s *Server) HandleEnterSceneDoneReq(ctx *Context, req *pb.EnterSceneDoneReq) error {
	if err := s.handleEnterSceneDoneReq(ctx, req); err != nil {
		log.Printf("[GAME] Failed to handle EnterSceneDoneReq: %v", err)
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

// handle PostEnterSceneReq
//
//	flow:
//		*RECV <·· PostEnterSceneReq
//		SEND ··> ServerCondMeetQuestListUpdateNotify
//		SEND ··> FinishedParentQuestUpdateNotify
//		SEND ··> QuestProgressUpdateNotify
//		SEND ··> QuestListUpdateNotify
//		*SEND ··> PostEnterSceneRsp
//		SEND ··> AnchorPointDataNotify
//		SEND ··> OneofGatherPointDetectorDataNotify
//		SEND ··> ChatChannelDataNotify
func (s *Server) HandlePostEnterSceneReq(ctx *Context, req *pb.PostEnterSceneReq) error {
	if err := s.handlePostEnterSceneReq(ctx, req); err != nil {
		log.Printf("[GAME] Failed to handle PostEnterSceneReq: %v", err)
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
