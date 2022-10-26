package game

func (s *Server) AsyncSendPlayerData(ctx *Context) error {
	return ctx.Async(
		s.SendAvatarDataNotify,
	)
}
