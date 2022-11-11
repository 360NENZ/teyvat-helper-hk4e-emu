package game

import "github.com/teyvat-helper/hk4e-proto/pb"

func (s *Server) SendKeepAliveNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGmTalkReq(ctx *Context, req *pb.GmTalkReq) error {
	panic("not implement")
}

func (s *Server) SendGmTalkRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendShowMessageNotify(ctx *Context) error {
	panic("not implement")
}

// handle PingReq
//
//	flow:
//		RECV <·· PingReq
//		SEND ··> PingRsp
func (s *Server) HandlePingReq(ctx *Context, req *pb.PingReq) error {
	return s.SendPingRsp(ctx, req.GetClientTime())
}
func (s *Server) SendPingRsp(ctx *Context, clientTime uint32) error {
	return s.Send(ctx, &pb.PingRsp{ClientTime: clientTime})
}

func (s *Server) HandleGetOnlinePlayerListReq(ctx *Context, req *pb.GetOnlinePlayerListReq) error {
	panic("not implement")
}

func (s *Server) SendGetOnlinePlayerListRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendServerTimeNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendServerLogNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientReconnectNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendRobotPushPlayerDataNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientReportNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUnionCmdNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleGetOnlinePlayerInfoReq(ctx *Context, req *pb.GetOnlinePlayerInfoReq) error {
	panic("not implement")
}

func (s *Server) SendGetOnlinePlayerInfoRsp(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendCheckSegmentCRCNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) HandleCheckSegmentCRCReq(ctx *Context, req *pb.CheckSegmentCRCReq) error {
	panic("not implement")
}

func (s *Server) SendWorldPlayerRTTNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendEchoNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendUpdateRedPointNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendClientBulletCreateNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendChangeServerGlobalValueNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendGmTalkNotify(ctx *Context) error {
	panic("not implement")
}

func (s *Server) SendLastPacketPrintNotify(ctx *Context) error {
	panic("not implement")
}
