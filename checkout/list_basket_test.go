package checkout_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
)

// TestListItems is a test for the List Items method receiver on the checkout struct
func TestListItems(t *testing.T) {
	t.Run("calling get items returns a count of all items", func(t *testing.T) {
		var (
			itemA = "itemA"
			itemB = "itemB"

			expectedQuantityItemA uint = 2
			expectedQuantityItemB uint = 1
		)

		c := checkout.New(checkout.RepositoryConfig{})

		c.Scan(itemA)
		c.Scan(itemA)
		c.Scan(itemB)

		items := c.ListBasketItems()

		if quantity, ok := items[itemA]; !ok || quantity != expectedQuantityItemA {
			t.Fatalf("expected %d items in basket for %s, got basket: %#v", expectedQuantityItemA, itemB, items)
		}

		if quantity, ok := items[itemB]; !ok || quantity != expectedQuantityItemB {
			t.Fatalf("expected %d items in basket for %s, got basket: %#v", expectedQuantityItemB, itemB, items)
		}
	})
}
