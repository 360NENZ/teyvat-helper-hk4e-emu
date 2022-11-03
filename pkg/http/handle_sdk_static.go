package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/teyvat-helper/hk4e-emu/pkg/sdk"
)

func (s *Server) handleSDKGetAgreementInfos(c *gin.Context) {
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
		"marketing_agreements": []gin.H{},
	}))
}

func (s *Server) handleSDKGetShopPriceTier(c *gin.Context) {
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
		"suggest_currency":   "USD",
		"price_tier_version": "0",
		"tiers":              []gin.H{},
	}))
}

func (s *Server) handleSDKCompareProtocolVersion(c *gin.Context) {
	var req sdk.CompareProtocolVersionRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		c.AbortWithStatusJSON(http.StatusOK, sdk.NewResponse(-106, nil))
		return
	}
	var resp sdk.CompareProtocolVersionResponseData
	resp.Modified = true // TODO: check version
	resp.Protocol = &sdk.ProtocolVersion{
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
	c.JSON(http.StatusOK, sdk.NewResponse(0, &resp))
}

func (s *Server) handleSDKGetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
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
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
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
	c.JSON(http.StatusOK, sdk.NewResponse(0, gin.H{
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

func (s *Server) handleSDKABTest(c *gin.Context) {
	c.JSON(http.StatusOK, sdk.NewResponse(0, []gin.H{{
		"code":      1000,
		"type":      2,
		"config_id": "14",
		"period_id": "6036_99",
		"version":   "3",
		"configs":   gin.H{"cardType": "direct"},
	}}))
}
