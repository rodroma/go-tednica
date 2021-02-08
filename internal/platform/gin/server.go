package gin

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string

	PingHandler        PingHandler
	GetItemByIDHandler GetItemByIDHandler
}

func (s *Server) Run() error {
	return s.Engine.Run(":" + s.Port)
}
