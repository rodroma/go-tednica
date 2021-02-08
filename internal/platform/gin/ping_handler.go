package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingHandler struct{}

func (h PingHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "PONG")
}