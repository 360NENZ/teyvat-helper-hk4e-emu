package game

import (
	"crypto/x509"
	"encoding/pem"
	"os"
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
	rest, _ := os.ReadFile("data/secret.pem")
	var block *pem.Block
	for {
		block, rest = pem.Decode(rest)
		switch block.Type {
		case "DISPATCH SERVER RSA PRIVATE KEY":
			s.Server.PrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
		case "DISPATCH CLIENT RSA PUBLIC KEY 2":
			k, _ := x509.ParsePKCS1PublicKey(block.Bytes)
			s.Client[2] = &http.PublicKey{PublicKey: k}
		case "DISPATCH CLIENT RSA PUBLIC KEY 3":
			k, _ := x509.ParsePKCS1PublicKey(block.Bytes)
			s.Client[3] = &http.PublicKey{PublicKey: k}
		case "DISPATCH CLIENT RSA PUBLIC KEY 4":
			k, _ := x509.ParsePKCS1PublicKey(block.Bytes)
			s.Client[4] = &http.PublicKey{PublicKey: k}
		case "DISPATCH CLIENT RSA PUBLIC KEY 5":
			k, _ := x509.ParsePKCS1PublicKey(block.Bytes)
			s.Client[5] = &http.PublicKey{PublicKey: k}
		}
		if len(rest) == 0 {
			break
		}
	}
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
