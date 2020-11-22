package entities_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/entities"
)

func TestDealIsApplicable(t *testing.T) {
	var (
		itemA entities.ItemName = "itemA"
		itemB entities.ItemName = "itemB"
		itemC entities.ItemName = "itemC"

		itemAQuantity entities.ItemQuantity = 1
		itemBQuantity entities.ItemQuantity = 2
		itemCQuantity entities.ItemQuantity = 3
	)

	requiredBasket := entities.Basket{
		itemA: itemAQuantity,
		itemB: itemBQuantity,
		itemC: itemCQuantity,
	}

	t.Run("returns true if the deal is applicable", func(t *testing.T) {
		basket := entities.Basket{
			itemA: itemAQuantity,
			itemB: itemBQuantity,
			itemC: itemCQuantity,
		}

		deal := entities.Deal{
			RequiredSubbasket: requiredBasket,
		}

		if applicable := deal.IsApplicable(basket); !applicable {
			t.Fatal("expected the deal to be applicable")
		}
	})

	t.Run("returns false if the deal is not applicable", func(t *testing.T) {
		basket := entities.Basket{
			itemA: itemAQuantity,
			itemB: itemBQuantity,
		}

		deal := entities.Deal{
			RequiredSubbasket: requiredBasket,
		}

		if applicable := deal.IsApplicable(basket); applicable {
			t.Fatal("expected the deal to not be applicable")
		}
	})
}

func TestGetPrice(t *testing.T) {
	t.Skip("obviously in a professional context I would add tests for these methods, but won't do so here")
}

func TestGetItems(t *testing.T) {
	t.Skip("obviously in a professional context I would add tests for these methods, but won't do so here")
}
