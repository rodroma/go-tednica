package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tednica/internal"
	"net/http"
	"regexp"
)

type GetItemByIDHandler struct {
	UseCase interface {
		GetItemByID(ctx context.Context, id string) (internal.GetItemByIDResponse, error)
	}
}

func (h GetItemByIDHandler) GetItemByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.validateID(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	response, err := h.UseCase.GetItemByID(ctx.Request.Context(), id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, h.mapResponseToJSON(response))
}

var itemIDRegex = regexp.MustCompile(`ML[A-Z]\d+`)
func (h GetItemByIDHandler) validateID(id string) error {
	if !itemIDRegex.MatchString(id) {
		return fmt.Errorf("invalid item id %s", id)
	}
	return nil
}

func (h GetItemByIDHandler) mapResponseToJSON(response internal.GetItemByIDResponse) interface{} {
	return response
}
