package prices_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/entities"
	"github.com/CharlesWinter/checkout/prices"
)

func TestGetItemPrice(t *testing.T) {
	var (
		itemNameA entities.ItemName = "A"
		itemNameB entities.ItemName = "B"

		itemPriceA uint = 50
	)

	t.Run("returns an items price if present", func(t *testing.T) {
		repo := prices.Repository{
			Prices: map[entities.ItemName]uint{
				itemNameA: itemPriceA,
			},
		}

		if got := repo.GetItemPrice(itemNameA); got != itemPriceA {
			t.Fatalf("expected item price %d, got %d", itemPriceA, got)
		}
	})

	// NOTE: This would obviously return an error if done in real life
	t.Run("returns 0 if the item is not present", func(t *testing.T) {
		repo := prices.Repository{
			Prices: map[entities.ItemName]uint{
				itemNameA: itemPriceA,
			},
		}

		var wantPrice uint = 0
		if got := repo.GetItemPrice(itemNameB); got != wantPrice {
			t.Fatalf("expected item price %d, got %d", itemPriceA, got)
		}
	})
}
