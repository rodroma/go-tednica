package di

import (
	actualGin "github.com/gin-gonic/gin"
	"go-tednica/internal/platform/gin"
	"os"
)

func ProvideServer() *gin.Server {
	s := &gin.Server{
		Engine:      actualGin.Default(),
		Port:        providePort(),
		PingHandler: gin.PingHandler{},
		GetItemByIDHandler: gin.GetItemByIDHandler{
			UseCase: provideGetItemByIDHandler(),
		},
	}

	s.MapRoutesToHandlers()

	return s
}

func providePort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}
