package checkout_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
	"github.com/CharlesWinter/checkout/entities"
)

// TestScan is a test for the Scan method receiver on the checkout struct
func TestScan(t *testing.T) {
	t.Run("calling scan successfully adds the item to the checkouts basket", func(t *testing.T) {
		var itemToAdd entities.ItemName = "itemA"
		c := checkout.New(checkout.RepositoryConfig{})

		c.Scan(itemToAdd)

		items := c.ListBasketItems()

		if quantity, ok := items[itemToAdd]; !ok || quantity != 1 {
			t.Fatalf("unexpected basket: %#v", items)
		}
	})
}
