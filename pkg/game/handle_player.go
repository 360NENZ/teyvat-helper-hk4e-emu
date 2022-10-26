package game

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/teyvat-helper/hk4e-emu/pkg/pb"
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
	uid, err := strconv.ParseInt(req.GetAccountUid(), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	player, err := ctx.Session().TokenLogin(ctx, uid, req.GetAccountToken())
	if err != nil {
		return 0, 0, err
	}
	player.SetBinary(&pb.PlayerDataBin{})
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

func (s *Server) SendPlayerLogoutRsp(ctx *Context, req *pb.PlayerLogoutRsp) error {
	panic("not implement")
}

func (s *Server) SendPlayerLogoutNotify(ctx *Context, req *pb.PlayerLogoutNotify) error {
	panic("not implement")
}

func (s *Server) SendPlayerDataNotify(ctx *Context, req *pb.PlayerDataNotify) error {
	panic("not implement")
}

func (s *Server) HandleChangeGameTimeReq(ctx *Context, req *pb.ChangeGameTimeReq) error {
	panic("not implement")
}

func (s *Server) SendChangeGameTimeRsp(ctx *Context, req *pb.ChangeGameTimeRsp) error {
	panic("not implement")
}

func (s *Server) SendPlayerGameTimeNotify(ctx *Context, req *pb.PlayerGameTimeNotify) error {
	panic("not implement")
}

func (s *Server) SendPlayerPropNotify(ctx *Context, req *pb.PlayerPropNotify) error {
	panic("not implement")
}

func (s *Server) SendClientTriggerEventNotify(ctx *Context, req *pb.ClientTriggerEventNotify) error {
	panic("not implement")
}

func (s *Server) HandleSetPlayerPropReq(ctx *Context, req *pb.SetPlayerPropReq) error {
	panic("not implement")
}

func (s *Server) SendSetPlayerPropRsp(ctx *Context, req *pb.SetPlayerPropRsp) error {
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
	avatar := ctx.Session().GetPlayer().Avatar()
	if err := avatar.SetPlayerBornData(ctx, req); err != nil {
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
