package gin

import (
	"github.com/gin-gonic/gin"
	"go-tednica/internal/commonerrors"
	"net/http"
)

func errorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		ginErr := ctx.Errors.Last()

		if ginErr == nil {
			return
		}

		err := ginErr.Err

		if commonerrors.IsBadArgument(err) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		}
	}
}
