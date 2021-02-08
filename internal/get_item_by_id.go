package internal

import (
	"context"
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
	item, err := uc.ItemGetter.GetItem(ctx, id)
	if err != nil {
		return GetItemByIDResponse{}, err
	}
	response := GetItemByIDResponse{Item: item}

	salePrice, err := uc.SalePriceGetter.GetSalePrice(ctx, id)
	if err != nil {
		return response, nil
	}

	response.OriginalPrice = salePrice.RegularPrice
	response.Price = salePrice.Price

	return response, nil
}
