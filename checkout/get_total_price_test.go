package checkout_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
	"github.com/CharlesWinter/checkout/entities"
	gomock "github.com/golang/mock/gomock"
)

const (
	itemNameA entities.ItemName = "A"
	itemNameB entities.ItemName = "B"
	itemNameC entities.ItemName = "C"

	itemPriceA uint = 50
	itemPriceB uint = 30
	itemPriceC uint = 20
)

func TestGetTotalPrice(t *testing.T) {
	t.Run("returns the correct total price for a basket with different items", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockPriceChecker := NewMockPriceChecker(ctrl)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameA).Return(itemPriceA)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameB).Return(itemPriceB)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameC).Return(itemPriceC)

		c := checkout.New(checkout.RepositoryConfig{
			PriceChecker: mockPriceChecker,
		})

		c.Scan(itemNameA)
		c.Scan(itemNameB)
		c.Scan(itemNameC)

		var want uint = itemPriceA + itemPriceB + itemPriceC
		if got := c.GetTotalPrice(); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})

	t.Run("returns the total price correctly for a basket with a deal", func(t *testing.T) {
		var dealPrice uint = 43

		ctrl := gomock.NewController(t)

		deals := []entities.Deal{
			{
				RequiredSubbasket: entities.Basket{
					itemNameA: 1,
					itemNameB: 2,
					itemNameC: 3,
				},
				Price: dealPrice,
			},
		}

		mockDealGetter := NewMockDealGetter(ctrl)
		mockDealGetter.EXPECT().GetDeals().Return(deals)

		mockPriceChecker := NewMockPriceChecker(ctrl)

		repo := checkout.New(checkout.RepositoryConfig{
			DealGetter:   mockDealGetter,
			PriceChecker: mockPriceChecker,
		})

		repo.Scan(itemNameA)

		repo.Scan(itemNameB)
		repo.Scan(itemNameB)

		repo.Scan(itemNameC)
		repo.Scan(itemNameC)
		repo.Scan(itemNameC)

		got := repo.GetTotalPrice()
		if got != dealPrice {
			t.Fatalf("wanted price %d, got %d", dealPrice, got)
		}
	})
}
