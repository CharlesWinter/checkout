package checkout_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
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
		c := checkout.New(checkout.RepositoryConfig{})

		c.Scan(itemNameA)
		c.Scan(itemNameB)
		c.Scan(itemNameC)

		var want uint = 100
		if got := c.GetTotalPrice(); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})
}
