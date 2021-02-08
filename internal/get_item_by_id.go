package internal

import (
	"context"
	"sync"
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
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	var wg sync.WaitGroup
	wg.Add(2)

	var err error

	var item Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		item, err = uc.ItemGetter.GetItem(ctx, id)

		if err != nil {
			cancelFunc()
		}
	}(&wg)

	var salePrice *SalePrice
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		sp, err := uc.SalePriceGetter.GetSalePrice(ctx, id)

		if err != nil {
			// Do stuff, log maybe?
			return
		}

		salePrice = &sp
	}(&wg)

	wg.Wait()

	if err != nil {
		return GetItemByIDResponse{}, err
	}

	response := GetItemByIDResponse{Item: item}
	if salePrice != nil {
		response.Price = salePrice.Price
		response.OriginalPrice = salePrice.RegularPrice
	}

	return response, nil
}
