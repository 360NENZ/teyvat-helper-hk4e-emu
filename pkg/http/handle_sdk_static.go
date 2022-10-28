package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *Server) handleSDKGetAgreementInfos(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, gin.H{
		"marketing_agreements": []any{},
	}))
}

type sdkCompareProtocolVersionRequestData struct {
	ID        string `json:"id"`
	AppID     string `json:"app_id"`
	ChannelID string `json:"channel_id"`
	Language  string `json:"language"`
	Major     string `json:"major"`
	Minimum   string `json:"minimum"`
}

type sdkProtocolVersion struct {
	ID            int64  `json:"id"`
	AppID         int64  `json:"app_id"`
	Language      string `json:"language"`
	UserProto     string `json:"user_proto"`
	PrivProto     string `json:"priv_proto"`
	Major         int32  `json:"major"`
	Minimum       int32  `json:"minimum"`
	CreateTime    string `json:"create_time"`
	TeenagerProto string `json:"teenager_proto"`
	ThirdProto    string `json:"third_proto"`
}

type sdkCompareProtocolVersionResponseData struct {
	Modified bool                `json:"modified"`
	Protocol *sdkProtocolVersion `json:"protocol"`
}

func (s *Server) handleSDKCompareProtocolVersion(c *gin.Context) {
	var req sdkCompareProtocolVersionRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		c.AbortWithStatusJSON(http.StatusOK, newSDKResponse(-106, nil))
		return
	}
	var resp sdkCompareProtocolVersionResponseData
	resp.Modified = true // TODO: check version
	resp.Protocol = &sdkProtocolVersion{
		AppID:      4,
		Language:   req.Language,
		CreateTime: "0",
	}
	switch req.Language {
	case "es", "pt", "ru":
		resp.Protocol.Major = 5
	case "de", "fr", "id", "ja", "ko", "th", "vi":
		resp.Protocol.Major = 6
	case "zh-cn":
		resp.Protocol.Major = 6
		resp.Protocol.Minimum = 1
	case "en":
		resp.Protocol.Major = 9
	case "zh-tw":
		resp.Protocol.Major = 10
	}
	c.JSON(http.StatusOK, newSDKResponse(0, &resp))
}

func (s *Server) handleSDKGetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, gin.H{
		"protocol":                  true,
		"qr_enabled":                false,
		"log_level":                 "INFO",
		"announce_url":              "https://" + s.config.BaseDomain + "/hk4e/announcement/index.html?sdk_presentation_style=fullscreen&announcement_version=1.21&sdk_screen_transparent=true&game_biz=hk4e_global&auth_appid=announcement&game=hk4e#/",
		"push_alias_type":           2,
		"disable_ysdk_guard":        false,
		"enable_announce_pic_popup": true,
	}))
}

func (s *Server) handleSDKLoadConfig(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, gin.H{
		"id":                     6,
		"game_key":               "hk4e_global",
		"client":                 "PC",
		"identity":               "I_IDENTITY",
		"guest":                  false,
		"ignore_versions":        "",
		"scene":                  "S_NORMAL",
		"name":                   "HK4E Global",
		"disable_regist":         false,
		"enable_email_captcha":   false,
		"thirdparty":             []string{"fb", "tw"},
		"disable_mmt":            false,
		"server_guest":           false,
		"thirdparty_ignore":      gin.H{},
		"enable_ps_bind_account": false,
		"thirdparty_login_configs": gin.H{
			"fb": gin.H{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 2592000},
			"tw": gin.H{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 2592000},
		},
	}))
}

func (s *Server) handleSDKConfigCombo(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, gin.H{
		"vals": gin.H{
			"pay_payco_centered_host":    "bill.payco.com",
			"email_bind_remind":          "true",
			"disable_email_bind_skip":    "false",
			"email_bind_remind_interval": "7",
			"network_report_config":      "{ \"enable\": 1, \"status_codes\": [206], \"url_paths\": [\"dataUpload\"] }",
			"kibana_pc_config":           "{ \"enable\": 1, \"level\": \"Info\",\"modules\": [\"download\"]",
			"kcp_enable":                 "false",
		},
	}))
}

func (s *Server) handleABTest(c *gin.Context) {
	c.JSON(http.StatusOK, newSDKResponse(0, []gin.H{{
		"code":      1000,
		"type":      2,
		"config_id": "14",
		"period_id": "6036_99",
		"version":   "3",
		"configs":   gin.H{"cardType": "direct"},
	}}))
}
