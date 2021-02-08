package gin

import "github.com/gin-gonic/gin"

func (s *Server) MapRoutesToHandlers() {
	r := s.Engine

	r.Use(errorMiddleware())

	r.GET("/ping", s.PingHandler.Ping)
	r.GET("/items/:id", liftFallibleHandler(s.GetItemByIDHandler.GetItemByID))
}

type fallibleHandler = func(*gin.Context) error
func liftFallibleHandler(h fallibleHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := h(ctx); err != nil {
			_ = ctx.Error(err)
		}
	}
}
