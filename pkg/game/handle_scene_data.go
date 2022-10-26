package game

func (s *Server) AsyncSendSceneData(ctx *Context) error {
	return ctx.Async(
		s.SendPlayerEnterSceneInfoNotify,
		s.SendSceneTeamUpdateNotify,
	)
}
