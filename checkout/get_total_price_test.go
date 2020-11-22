package checkout_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
	gomock "github.com/golang/mock/gomock"
)

const (
	itemNameA = "A"
	itemNameB = "B"
	itemNameC = "C"

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
}
