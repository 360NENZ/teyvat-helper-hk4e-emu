package game

import "github.com/teyvat-helper/hk4e-proto/pb"

func (s *Server) HandleGmTalkReq(ctx *Context, req *pb.GmTalkReq) error {
	panic("not implement")
}

// handle PingReq
//
//	flow:
//		*RECV <·· PingReq
//		*SEND ··> PingRsp
func (s *Server) HandlePingReq(ctx *Context, req *pb.PingReq) error {
	return s.Send(ctx, &pb.PingRsp{ClientTime: req.GetClientTime()})
}

func (s *Server) HandleGetOnlinePlayerListReq(ctx *Context, req *pb.GetOnlinePlayerListReq) error {
	panic("not implement")
}

func (s *Server) HandleGetOnlinePlayerInfoReq(ctx *Context, req *pb.GetOnlinePlayerInfoReq) error {
	panic("not implement")
}

func (s *Server) HandleCheckSegmentCRCReq(ctx *Context, req *pb.CheckSegmentCRCReq) error {
	panic("not implement")
}
