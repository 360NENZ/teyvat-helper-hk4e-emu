package game

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/net"
	"github.com/teyvat-helper/hk4e-emu/pkg/store"
	"github.com/teyvat-helper/hk4e-proto/pb"
	"google.golang.org/protobuf/proto"
)

type PlayerSession struct {
	*net.Session
	seed   uint64
	secret *Secret
	server *Server
	player *Player
}

func NewPlayerSession(sess *net.Session, secret *Secret, server *Server) *PlayerSession {
	session := &PlayerSession{
		Session: sess,
		secret:  secret,
		server:  server,
	}
	return session
}

func (s *PlayerSession) CreatePlayer(ctx context.Context, id int64) (*store.Player, error) {
	player := new(store.Player)
	player.AccountID = id
	if err := s.server.Store().Player().CreatePlayer(ctx, player); err != nil {
		return nil, err
	}
	if err := s.server.Store().PlayerData().CreatePlayerData(ctx, &store.PlayerData{ID: player.ID}); err != nil {
		return nil, err
	}
	return player, nil
}

func (s *PlayerSession) UpdatePlayer(ctx context.Context) (*store.Player, error) {
	data, err := proto.Marshal(s.player.binary)
	if err != nil {
		return nil, err
	}
	if err := s.server.Store().PlayerData().UpdatePlayerData(ctx, &store.PlayerData{ID: s.player.ID, BinData: data}); err != nil {
		return nil, err
	}
	return s.player.Player, nil
}

func (s *PlayerSession) GetPlayerData(ctx context.Context, id int32) (*store.PlayerData, error) {
	return s.server.Store().PlayerData().GetPlayerData(ctx, id)
}

func (s *PlayerSession) SetPlayer(player *Player) { s.player = player }
func (s *PlayerSession) GetPlayer() *Player       { return s.player }

func (s *PlayerSession) SetSeed(seed uint64) { s.seed = seed }

func (s *PlayerSession) handlePacket(packet *net.Packet, onPacket func(*Packet)) {
	data := packet.Payload()
	defer packet.Release()
	if s.seed == 0 {
		s.secret.Shared.Xor(data)
	} else {
		s.secret.GetKeyBlock(s.seed).Xor(data)
	}
	l := len(data)
	if l < 12 {
		log.Warn().Msgf("packet too short: %d", l)
		return
	}
	if data[0] != 0x45 || data[1] != 0x67 {
		log.Warn().Msgf("invalid packet prefix: %s", hex.EncodeToString(data[:2]))
		return
	}
	if data[l-2] != 0x89 || data[l-1] != 0xAB {
		log.Warn().Msgf("invalid packet suffix: %s", hex.EncodeToString(data[l-2:]))
		return
	}
	buf := bytes.NewBuffer(data[2 : l-2])
	cmd := pb.ProtoCommand(binary.BigEndian.Uint16(buf.Next(2)))
	headLength := binary.BigEndian.Uint16(buf.Next(2))
	bodyLength := binary.BigEndian.Uint32(buf.Next(4))
	if uint32(l) != 12+uint32(headLength)+bodyLength {
		log.Warn().Msgf("invalid packet length: %d", l)
		return
	}
	head := &pb.PacketHead{}
	if headLength > 0 {
		_ = proto.Unmarshal(buf.Next(int(headLength)), head)
	}
	body := pb.ProtoCommandNewFuncMap.New(cmd)
	var temp []byte
	if body != nil && bodyLength > 0 {
		_ = proto.Unmarshal(buf.Next(int(bodyLength)), body)
	} else if bodyLength > 0 {
		temp = make([]byte, bodyLength)
		copy(temp, buf.Bytes())
	}
	go onPacket(&Packet{session: s, head: head, message: body, command: cmd, rawData: temp})
}

func (s *PlayerSession) Send(head *pb.PacketHead, body pb.ProtoMessage) error {
	p1, err := proto.Marshal(head)
	if err != nil {
		return err
	}
	p2, err := proto.Marshal(body)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	buf.WriteByte(0x45)
	buf.WriteByte(0x67)
	binary.Write(buf, binary.BigEndian, body.ProtoCommand())
	binary.Write(buf, binary.BigEndian, uint16(len(p1)))
	binary.Write(buf, binary.BigEndian, uint32(len(p2)))
	buf.Write(p1)
	buf.Write(p2)
	buf.WriteByte(0x89)
	buf.WriteByte(0xAB)
	b := buf.Bytes()
	if s.seed == 0 {
		s.secret.Shared.Xor(b)
	} else {
		s.secret.GetKeyBlock(s.seed).Xor(b)
	}
	return s.Session.Send(b)
}
