package http

func (s *Server) initRouter() {
	s.router.GET("/", s.handleDefault)
	s.router.POST("/", s.handleDefault)

	s.router.OPTIONS("/ping", s.handleOptionsPing)
	s.router.GET("/ping", s.handlePing)
	s.router.POST("/ping", s.handlePing)

	// api handlers
	api := s.router.Group("/api")
	{
		api.GET("/publicKey", s.handleAPIPublicKey)
		api.GET("/status", s.handleAPIStatus)
		api.GET("/status/legacy", s.handleAPIStatusLegacy)

		api.POST("/account/token/check", s.handleAPITokenCheck)
	}

	// Mainland: gameapi-account.*
	// Overseas: api-account-os.*
	s.router.POST("/account/risky/api/check", s.handleSDKRiskyCheck)
	s.router.POST("/account/device/api/listNewerDevices")

	// Overseas: account.*
	s.router.GET("/sdkFacebookLogin.html")
	s.router.GET("/sdkTwitterLogin.html")

	// Mainland: hk4e-sdk.*
	// Overseas: hk4e-sdk-os.*
	s.router.POST("/hk4e_cn/mdk/shield/api/emailCaptcha")
	s.router.POST("/hk4e_cn/mdk/shield/api/login", s.handleSDKShieldLogin)
	s.router.POST("/hk4e_cn/mdk/shield/api/loginMobile")
	s.router.POST("/hk4e_cn/mdk/shield/api/loginCaptcha")
	s.router.POST("/hk4e_cn/mdk/shield/api/verify", s.handleSDKShieldVerify)
	s.router.POST("/hk4e_global/mdk/shield/api/emailCaptcha")
	s.router.POST("/hk4e_global/mdk/shield/api/login", s.handleSDKShieldLogin)
	s.router.POST("/hk4e_global/mdk/shield/api/loginMobile")
	s.router.POST("/hk4e_global/mdk/shield/api/loginCaptcha")
	s.router.POST("/hk4e_global/mdk/shield/api/verify", s.handleSDKShieldVerify)

	// Mainland: hk4e-sdk.*
	// Overseas: hk4e-sdk-os.*
	s.router.POST("/hk4e_cn/combo/granter/login/v2/login", s.handleSDKComboLogin)
	s.router.POST("/hk4e_global/combo/granter/login/v2/login", s.handleSDKComboLogin)

	// Mainland: dispatchcnglobal.*
	// Overseas: dispatchosglobal.*
	s.router.GET("/query_region_list", s.handleQueryRegionList())
	s.router.GET("/query_cur_region", s.handleQueryCurrentRegion())
	s.router.GET("/query_cur_region/:id", s.handleQueryCurrentRegion())

	// Mainland: uspider.*
	// Overseas: overseauspider.*
	s.router.POST("/log", s.handleLogUpload)

	// Mainland: log-upload.*
	// Overseas: log-upload-os.*
	s.router.POST("/2g/dataUpload", s.handleLogUpload)
	s.router.POST("/crash/dataUpload", s.handleLogUpload)
	s.router.POST("/log/sdk/upload", s.handleLogUpload)
	s.router.POST("/perf/config/verify", s.handleLogUpload)
	s.router.POST("/perf/dataUpload", s.handleLogUpload)
	s.router.POST("/sdk/dataUpload", s.handleLogUpload)

	// Mainland: webstatic.*
	// Overseas: webstatic-sea.*
	s.router.GET("/admin/mi18n/plat_oversea/*any", s.handleWebStaticJSON)
	s.router.GET("/hk4e/announcement/index.html")

	// Mainland: hk4e-sdk.*
	// Overseas: hk4e-sdk-os.*
	s.router.GET("/common/hk4e_global/announcement/api/getAlertPic")
	s.router.GET("/common/hk4e_global/announcement/api/getAlertAnn")
	s.router.GET("/common/hk4e_global/announcement/api/getAnnList")
	s.router.GET("/hk4e_cn/mdk/agreement/api/getAgreementInfos", s.handleSDKGetAgreementInfos)
	s.router.GET("/hk4e_global/mdk/agreement/api/getAgreementInfos", s.handleSDKGetAgreementInfos)
	s.router.POST("/hk4e_cn/mdk/shopwindow/shopwindow/listPriceTier", s.handleSDKGetShopPriceTier)
	s.router.POST("/hk4e_global/mdk/shopwindow/shopwindow/listPriceTier", s.handleSDKGetShopPriceTier)
	s.router.POST("/hk4e_cn/combo/granter/api/compareProtocolVersion", s.handleSDKCompareProtocolVersion)
	s.router.POST("/hk4e_global/combo/granter/api/compareProtocolVersion", s.handleSDKCompareProtocolVersion)

	// Mainland: hk4e-sdk-static.*
	// Overseas: hk4e-sdk-os-static.*
	s.router.GET("/common/hk4e_global/announcement/api/getAnnContent")
	s.router.GET("/hk4e_cn/combo/granter/api/getConfig", s.handleSDKGetConfig)
	s.router.POST("/hk4e_cn/mdk/shield/api/loadConfig", s.handleSDKLoadConfig)
	s.router.GET("/hk4e_global/combo/granter/api/getConfig", s.handleSDKGetConfig)
	s.router.POST("/hk4e_global/mdk/shield/api/loadConfig", s.handleSDKLoadConfig)

	// Overseas: sdk-os-static.*
	s.router.GET("/combo/box/api/config/sdk/combo", s.handleSDKConfigCombo)

	// Overseas: abtest-api-data-sg.*
	s.router.POST("/data_abtest_api/config/experiment/list", s.handleSDKABTest)
}
