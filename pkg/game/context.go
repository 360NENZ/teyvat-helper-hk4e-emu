package game

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/aj3423/aproto"
	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-proto/pb"
	"google.golang.org/protobuf/proto"
)

func (s *Server) Context(packet *Packet) *Context {
	head, _ := json.Marshal(packet.head)
	if packet.message != nil {
		body, _ := json.Marshal(convertMessage(packet.message))
		log.Debug().RawJSON("head", head).RawJSON("body", body).
			Msgf("RECV <·· %5d - %5d:%s", packet.head.GetClientSequenceId(), packet.message.ProtoCommand(), packet.message.ProtoMessageType())
	} else {
		log.Debug().RawJSON("head", head).Hex("body", packet.rawData).
			Msgf("RECV <·· %5d - %5d:*****", packet.head.GetClientSequenceId(), packet.command)
		fmt.Println(aproto.Dump(packet.rawData))
	}
	return &Context{Context: context.Background(), session: packet.session, head: packet.head}
}

func (s *Server) Send(ctx *Context, message pb.ProtoMessage) error {
	name := message.ProtoMessageType()
	if block := s.filter.Map(ctx, name); block {
		log.Warn().Uint16("cmdid", uint16(message.ProtoCommand())).Str("name", string(name)).Msg("Blocked send packet")
		return nil
	}
	head, _ := json.Marshal(ctx.head)
	body, _ := json.Marshal(message)
	log.Debug().RawJSON("head", head).RawJSON("body", body).
		Msgf("SEND ··> %5d - %5d:%s", ctx.head.GetClientSequenceId(), message.ProtoCommand(), name)
	return ctx.Session().Send(ctx.head, message)
}

func (s *Server) SendFunc(ctx *Context, message pb.ProtoMessage) func(ctx *Context) error {
	return func(ctx *Context) error { return s.Send(ctx, message) }
}

type Context struct {
	context.Context
	session    *PlayerSession
	head       *pb.PacketHead
	sceneToken uint32
}

func (ctx *Context) Session() *PlayerSession { return ctx.session }

func (ctx *Context) GetSceneToken() uint32      { return ctx.sceneToken }
func (ctx *Context) SetSceneToken(token uint32) { ctx.sceneToken = token }

func (ctx *Context) Async(fs ...func(ctx *Context) error) error {
	errCh := make(chan error)
	wgDone := make(chan bool)
	var wg sync.WaitGroup
	for _, f := range fs {
		wg.Add(1)
		go func(f func(ctx *Context) error) {
			if err := f(ctx); err != nil {
				errCh <- err
			}
			wg.Done()
		}(f)
	}
	go func() {
		wg.Wait()
		close(wgDone)
	}()
	select {
	case err := <-errCh:
		return err
	case <-wgDone:
		return nil
	}
}

type UnionCmdNotify struct {
	CmdList []*UnionCmd `json:"cmd_list"`
}

type UnionCmd struct {
	MessageID pb.ProtoCommand `json:"message_id"`
	Body      pb.ProtoMessage `json:"body"`
}

func convertMessage(message pb.ProtoMessage) any {
	var v any
	if message.ProtoMessageType() == "UnionCmdNotify" {
		failed := false
		notify := UnionCmdNotify{}
		for _, cmd := range message.(*pb.UnionCmdNotify).CmdList {
			id := pb.ProtoCommand(cmd.GetMessageId())
			item := UnionCmd{
				MessageID: id,
				Body:      pb.ProtoCommandNewFuncMap.New(id),
			}
			if item.Body != nil {
				_ = proto.Unmarshal(cmd.GetBody(), item.Body)
				notify.CmdList = append(notify.CmdList, &item)
			} else {
				failed = true
			}
		}
		if !failed {
			v = notify
		} else {
			v = message
		}
	} else {
		v = message
	}
	return v
}
