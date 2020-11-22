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

	t.Run("returns the total price correctly for a basket with a deal and products", func(t *testing.T) {
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

	t.Run("if there is an applicable deal and other items in the basket", func(t *testing.T) {
		var dealPriceA uint = 130
		var dealPriceB uint = 45

		ctrl := gomock.NewController(t)

		deals := []entities.Deal{
			{
				RequiredSubbasket: entities.Basket{
					itemNameA: 3,
				},
				Price: dealPriceA,
			},
			{
				RequiredSubbasket: entities.Basket{
					itemNameB: 2,
				},
				Price: dealPriceB,
			},
		}

		mockDealGetter := NewMockDealGetter(ctrl)
		mockDealGetter.EXPECT().GetDeals().Return(deals)

		mockPriceChecker := NewMockPriceChecker(ctrl)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameA).Times(1).Return(itemPriceA)

		repo := checkout.New(checkout.RepositoryConfig{
			DealGetter:   mockDealGetter,
			PriceChecker: mockPriceChecker,
		})

		repo.Scan(itemNameB)
		repo.Scan(itemNameA)
		repo.Scan(itemNameB)

		got := repo.GetTotalPrice()
		want := uint(95)
		if got != want {
			t.Fatalf("wanted price %d, got %d", want, got)
		}
	})

	t.Run("checks that GetTotalPrice is idempotent", func(t *testing.T) {
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
		mockDealGetter.EXPECT().GetDeals().Return(deals)

		mockPriceChecker := NewMockPriceChecker(ctrl)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameA).Times(4).Return(itemPriceA)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameB).Times(2).Return(itemPriceB)
		mockPriceChecker.EXPECT().GetItemPrice(itemNameC).Times(2).Return(itemPriceC)

		repo := checkout.New(checkout.RepositoryConfig{
			DealGetter:   mockDealGetter,
			PriceChecker: mockPriceChecker,
		})

		repo.Scan(itemNameA)
		repo.Scan(itemNameA)
		repo.Scan(itemNameA)

		repo.Scan(itemNameB)
		repo.Scan(itemNameB)
		repo.Scan(itemNameB)

		repo.Scan(itemNameC)
		repo.Scan(itemNameC)
		repo.Scan(itemNameC)
		repo.Scan(itemNameC)

		got := repo.GetTotalPrice()
		got = repo.GetTotalPrice()
		want := uint(dealPrice + 2*itemPriceA + itemPriceB + itemPriceC)
		if got != want {
			t.Fatalf("wanted price %d, got %d", want, got)
		}
	})
}
