package game

import (
	"log"

	"github.com/teyvat-helper/hk4e-emu/pkg/pb"
)

type Packet struct {
	session *PlayerSession
	head    *pb.PacketHead
	message pb.ProtoMessage
	command pb.ProtoCommand
	rawData []byte
}

func (s *Server) onPacket(packet *Packet) {
	var err error
	ctx := s.Context(packet)
	switch message := packet.message.(type) {

	// handle misc
	case *pb.PingReq:
		err = s.HandlePingReq(ctx, message)

	// handle player
	case *pb.GetPlayerTokenReq:
		err = s.HandleGetPlayerTokenReq(ctx, message)
	case *pb.PlayerLoginReq:
		err = s.HandlePlayerLoginReq(ctx, message)
	case *pb.PlayerLogoutReq:
		err = s.HandlePlayerLogoutReq(ctx, message)
	case *pb.ChangeGameTimeReq:
		err = s.HandleChangeGameTimeReq(ctx, message)
	case *pb.SetPlayerPropReq:
		err = s.HandleSetPlayerPropReq(ctx, message)
	case *pb.SetPlayerBornDataReq:
		err = s.HandleSetPlayerBornDataReq(ctx, message)

	// handle scene
	case *pb.SceneInitFinishReq:
		err = s.HandleSceneInitFinishReq(ctx, message)
	case *pb.EnterSceneReadyReq:
		err = s.HandleEnterSceneReadyReq(ctx, message)
	case *pb.EnterSceneDoneReq:
		err = s.HandleEnterSceneDoneReq(ctx, message)
	case *pb.PostEnterSceneReq:
		err = s.HandlePostEnterSceneReq(ctx, message)

	}
	if err != nil {
		log.Printf("[GAME] Failed to handle %s: %v", packet.message.ProtoMessageType(), err)
	}
}
