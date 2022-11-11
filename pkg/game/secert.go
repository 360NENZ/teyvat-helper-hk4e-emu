package game

import (
	"sync"

	"github.com/teyvat-helper/hk4e-emu/pkg/ec2b"
	"github.com/teyvat-helper/hk4e-emu/pkg/http"
	"github.com/teyvat-helper/hk4e-emu/pkg/mt19937"
)

type Secret struct {
	mutex  sync.Mutex
	Shared *ec2b.Ec2b
	Server *http.PrivateKey
	Client map[uint32]*http.PublicKey
	KeyMap map[uint64]*mt19937.KeyBlock
}

func NewSecret() *Secret {
	s := &Secret{}
	s.Shared = ec2b.NewEc2b()
	s.Server = &http.PrivateKey{}
	s.Client = make(map[uint32]*http.PublicKey)
	s.KeyMap = make(map[uint64]*mt19937.KeyBlock)
	return s
}

func (s *Secret) GetKeyBlock(seed uint64) *mt19937.KeyBlock {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if b, ok := s.KeyMap[seed]; ok {
		return b
	}
	b := mt19937.NewKeyBlock(seed)
	s.KeyMap[seed] = b
	return b
}
