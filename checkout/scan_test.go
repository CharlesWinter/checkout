package checkout

import (
	"testing"

	"github.com/CharlesWinter/checkout/checkout"
)

// TestScan is a test for the Scan method receiver on the checkout struct
func TestScan(t *testing.T) {
	t.Run("calling scan successfully adds the item to the checkouts basket", func(t *testing.T) {
		itemToAdd := "itemA"
		c := checkout.New(checkout.RepositoryConfig{})

		c.Scan(itemToAdd)

		if items := c.ListBasketItems(); len(items) != 1 || items[0] != itemToAdd {
			t.Fatalf("expected item %s in basket, got items: %v", itemToAdd, items)
		}
	})
}
