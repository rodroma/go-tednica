package internal

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetItemByID_Should_Fail_If_ItemGetter_Fails(t *testing.T) {
	err := errors.New("oh no")
	uc := GetItemByID{
		ItemGetter: itemGetterFunc(func(_ context.Context, id string) (Item, error) {
			return Item{}, err
		}),
		SalePriceGetter: salePriceGetterFunc(func(_ context.Context, itemID string) (SalePrice, error) {
			return SalePrice{}, nil
		}),
	}

	_, err = uc.GetItemByID(context.Background(), "MLA123")

	assert.Error(t, err)
}

func Test_GetItemByID_Should_Take_Item_Price_If_SalePriceGetter_Fails(t *testing.T) {
	given := Item{
		ID:            "MLA123",
		Title:         "Title",
		OriginalPrice: 100,
		Price:         75,
	}

	uc := GetItemByID{
		ItemGetter:      itemGetterFunc(func(ctx context.Context, id string) (Item, error) {
			return given, nil
		}),
		SalePriceGetter: salePriceGetterFunc(func(ctx context.Context, itemID string) (SalePrice, error) {
			return SalePrice{}, errors.New("oh no")
		}),
	}

	actual, err := uc.GetItemByID(context.Background(), "MLA123")

	assert.NoError(t, err)
	assert.Equal(t, given.ID, actual.ID)
	assert.Equal(t, given.Title, actual.Title)
	assert.Equal(t, given.OriginalPrice, actual.OriginalPrice)
	assert.Equal(t, given.Price, actual.Price)
}

func Test_GetItemByID_Should_SalePrice_If_Getter_Succeeds(t *testing.T) {
	item := Item{
		ID:            "MLA123",
		Title:         "Title",
		OriginalPrice: 100,
		Price:         75,
	}

	salePrice := SalePrice{
		RegularPrice: 50,
		Price:        20,
	}

	uc := GetItemByID{
		ItemGetter:      itemGetterFunc(func(ctx context.Context, id string) (Item, error) {
			return item, nil
		}),
		SalePriceGetter: salePriceGetterFunc(func(ctx context.Context, itemID string) (SalePrice, error) {
			return salePrice, nil
		}),
	}

	actual, err := uc.GetItemByID(context.Background(), item.ID)

	assert.NoError(t, err)
	assert.Equal(t, item.ID, actual.ID)
	assert.Equal(t, item.Title, actual.Title)
	assert.Equal(t, salePrice.Price, actual.Price)
	assert.Equal(t, salePrice.RegularPrice, actual.OriginalPrice)
}

// fakes

type salePriceGetterFunc func(ctx context.Context, itemID string) (SalePrice, error)

func (f salePriceGetterFunc) GetSalePrice(ctx context.Context, itemID string) (SalePrice, error) {
	return f(ctx, itemID)
}

type itemGetterFunc func(ctx context.Context, id string) (Item, error)

func (f itemGetterFunc) GetItem(ctx context.Context, id string) (Item, error) {
	return f(ctx, id)
}
