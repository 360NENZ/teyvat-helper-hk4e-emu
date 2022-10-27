package game

func (s *Server) AsyncSendPlayerData(ctx *Context) error {
	return ctx.Async(
		// avatar
		s.SendAvatarDataNotify,

		// item
		s.SendPlayerStoreNotify,
		s.SendStoreWeightLimitNotify,

		// player
		s.SendPlayerDataNotify,
	)
}
