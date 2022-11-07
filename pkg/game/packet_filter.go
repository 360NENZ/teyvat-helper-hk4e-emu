package game

import (
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/config"
	"github.com/teyvat-helper/hk4e-proto/pb"
)

type PacketFilter struct {
	config *config.FilterConfig
	policy bool
	packet map[pb.ProtoMessageType]bool
}

func NewPacketFilter(cfg *config.Config) *PacketFilter {
	f := &PacketFilter{
		config: &cfg.GameServer.Filter,
		policy: true,
		packet: make(map[pb.ProtoMessageType]bool),
	}
	f.init()
	return f
}

func (f *PacketFilter) init() {
	if f.config == nil {
		return
	}
	for _, rule := range f.config.Rules {
		params := strings.Split(rule, ":")
		if len(params) != 2 {
			continue
		}
		policy := true
		switch params[0] {
		case "allow":
			policy = false
		case "block":
			policy = true
		}
		name := pb.ProtoMessageType(params[1])
		if name == "*" {
			f.policy = policy
			log.Debug().Str("policy", params[0]).Msg("PacketFilter policy updated")
		} else {
			f.packet[name] = policy
			log.Debug().Str("policy", params[0]).Str("packet", string(name)).Msg("PacketFilter packet updated")
		}
	}
}

func (f *PacketFilter) Map(ctx *Context, name pb.ProtoMessageType) bool {
	block, ok := f.packet[name]
	if !ok {
		return f.policy
	}
	return block
}
