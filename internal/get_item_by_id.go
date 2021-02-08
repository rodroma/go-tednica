package internal

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type GetItemByIDResponse struct {
	Item
}

type GetItemByID struct{
	SalePriceGetter interface {
		GetSalePrice(ctx context.Context, itemID string) (SalePrice, error)
	}

	ItemGetter interface {
		GetItem(ctx context.Context, id string) (Item, error)
	}
}

func (uc GetItemByID) GetItemByID(ctx context.Context, id string) (GetItemByIDResponse, error) {
	g, ctx := errgroup.WithContext(ctx)

	var item Item
	g.Go(func() error {
		var err error
		item, err = uc.ItemGetter.GetItem(ctx, id)
		return err
	})

	var salePrice *SalePrice
	g.Go(func() error {
		sp, err := uc.SalePriceGetter.GetSalePrice(ctx, id)
		if err != nil {
			// Log or do some thing
		} else {
			salePrice = &sp
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return GetItemByIDResponse{}, err
	}

	response := GetItemByIDResponse{Item: item}
	if salePrice != nil {
		response.Price = salePrice.Price
		response.OriginalPrice = salePrice.RegularPrice
	}

	return response, nil
}
