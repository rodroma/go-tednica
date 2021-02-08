package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-tednica/internal"
	"go-tednica/internal/commonerrors"
	"net/http"
	"regexp"
)

type GetItemByIDHandler struct {
	UseCase interface {
		GetItemByID(ctx context.Context, id string) (internal.GetItemByIDResponse, error)
	}
}

func (h GetItemByIDHandler) GetItemByID(ctx *gin.Context) {
	if err := h.getItemByID(ctx); err != nil {
		_ = ctx.Error(err)
	}
}

func (h GetItemByIDHandler) getItemByID(ctx *gin.Context) error {
	id := ctx.Param("id")
	if err := h.validateID(id); err != nil {
		return err
	}

	response, err := h.UseCase.GetItemByID(ctx.Request.Context(), id)

	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, h.mapResponseToJSON(response))

	return nil
}

var itemIDRegex = regexp.MustCompile(`ML[A-Z]\d+`)
func (h GetItemByIDHandler) validateID(id string) error {
	if !itemIDRegex.MatchString(id) {
		return commonerrors.BadArgument(fmt.Errorf("invalid item id %s", id))
	}
	return nil
}

func (h GetItemByIDHandler) mapResponseToJSON(response internal.GetItemByIDResponse) interface{} {
	return response
}
