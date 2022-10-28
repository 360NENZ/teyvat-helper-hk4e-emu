package game

import "github.com/teyvat-helper/hk4e-proto/pb"

// async send SceneData
//
//	flow:
//		SEND ··> PlayerEnterSceneInfoNotify
//		SEND ··> SceneDataNotify
//		SEND ··> SceneForceUnlockNotify
//		SEND ··> SceneTeamUpdateNotify
//		SEND ··> HostPlayerNotify
//		SEND ··> SceneTimeNotify
//		SEND ··> SceneAreaWeatherNotify
func (s *Server) AsyncSendSceneData(ctx *Context) error {
	player := ctx.Session().GetPlayer()
	return ctx.Async(
		s.SendPlayerEnterSceneInfoNotify,
		s.SendSceneTeamUpdateNotify,

		s.SendSceneForceUnlockNotify,

		s.SendFunc(ctx, &pb.HostPlayerNotify{
			HostUid:    uint32(player.ID),
			HostPeerId: 1,
		}),
		s.SendFunc(ctx, &pb.SceneTimeNotify{
			SceneId: 3,
		}),
		s.SendFunc(ctx, &pb.SceneAreaWeatherNotify{
			WeatherAreaId: 1,
			ClimateType:   1,
		}),
		s.SendFunc(ctx, &pb.WorldDataNotify{WorldPropMap: map[uint32]*pb.PropValue{
			1: {Type: 1, Value: &pb.PropValue_Ival{Ival: 0}},
			2: {Type: 2, Value: &pb.PropValue_Ival{Ival: 0}},
		}}),
		s.SendFunc(ctx, &pb.WorldOwnerDailyTaskNotify{}),
		s.SendFunc(ctx, &pb.AllMarkPointNotify{}),
	)
}
