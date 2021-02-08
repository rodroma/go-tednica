package gin

func (s *Server) MapRoutesToHandlers() {
	r := s.Engine

	r.Use(errorMiddleware())

	r.GET("/ping", s.PingHandler.Ping)
	r.GET("/items/:id", s.GetItemByIDHandler.GetItemByID)
}
