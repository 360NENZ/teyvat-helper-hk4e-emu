package http

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teyvat-helper/hk4e-proto/pb"
	"google.golang.org/protobuf/proto"
)

type customConfig struct {
	SDKEnvironment               string  `json:"sdkenv,omitempty"`
	ShowException                bool    `json:"showexception,omitempty"`
	RegionConfig                 string  `json:"regionConfig,omitempty"`
	DisableRazorChromaInit       bool    `json:"disableRazorChromaInit,omitempty"`
	DownloadEnableParallelVerify bool    `json:"downloadEnableParallelVerify,omitempty"`
	CodeSwitch                   []int32 `json:"codeSwitch,omitempty"`
}

func (s *Server) handleQueryRegionList() gin.HandlerFunc {
	m := pb.QueryRegionListHttpRsp{}
	for _, c := range s.config.GateServer {
		m.RegionList = append(m.RegionList, &pb.RegionSimpleInfo{
			Name:        c.Name,
			Title:       c.Title,
			Type:        "DEV_PUBLIC",
			DispatchUrl: "https://example.com/query_cur_region/" + c.Name,
		})
	}
	m.ClientSecretKey = s.secret.Shared.Key()
	m.ClientCustomConfigEncrypted, _ = json.Marshal(customConfig{
		SDKEnvironment:         "2",
		RegionConfig:           "pm|fk|add",
		DisableRazorChromaInit: true,
	})
	s.secret.Shared.Xor(m.GetClientCustomConfigEncrypted())
	m.EnableLoginPc = true
	b, _ := proto.Marshal(&m)
	return func(c *gin.Context) {
		c.String(http.StatusOK, base64.StdEncoding.EncodeToString(b))
	}
}

func (s *Server) handleQueryCurrentRegion() gin.HandlerFunc {
	regionMap := map[string][]byte{}
	defaultRegion := []byte{}
	for i, c := range s.config.GateServer {
		addr, _ := net.ResolveUDPAddr("udp", c.Addr)
		m := pb.QueryCurrRegionHttpRsp{}
		m.RegionInfo = &pb.RegionInfo{
			GateserverIp:   addr.IP.String(),
			GateserverPort: uint32(addr.Port),
			SecretKey:      s.secret.Shared.Key(),
		}
		m.ClientSecretKey = s.secret.Shared.Key()
		m.RegionCustomConfigEncrypted, _ = json.Marshal(customConfig{})
		s.secret.Shared.Xor(m.GetRegionCustomConfigEncrypted())
		b, _ := proto.Marshal(&m)
		regionMap[c.Name] = b
		if i == 0 {
			defaultRegion = b
		}
	}
	return func(c *gin.Context) {
		b, ok := regionMap[c.Param("id")]
		if !ok {
			b = defaultRegion
		}
		switch id := c.Query("key_id"); id {
		case "":
			c.String(http.StatusOK, base64.StdEncoding.EncodeToString(b))
		case "2", "3":
			pub, ok := s.secret.Client[id]
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "missing public key " + id})
				return
			}
			content, _ := pub.Encrypt(b)
			sign, _ := s.secret.Server.Sign(b)
			c.JSON(http.StatusOK, gin.H{
				"content": content,
				"sign":    sign,
			})
		default:
			c.JSON(http.StatusOK, gin.H{"error": "unknown public key " + id})
		}
	}
}
