package checkout

import (
	"github.com/CharlesWinter/checkout/entities"
)

// GetTotalPrice calculates and returns the total price of all items in the basket.
func (r Repository) GetTotalPrice() uint {
	basket := r.copyBasket()

	var totalPrice uint

	// Work through the shopping cart, applying any deals.
	for _, deal := range r.dealGetter.GetDeals() {
		if !deal.IsApplicable(basket) {
			continue
		}

		// if the deal is applied, remove the applied items from the cart and keep
		// going
		removeDealItems(basket, deal.GetItems())

		totalPrice = totalPrice + deal.GetPrice()
	}

	for name, quantity := range basket {
		if quantity == 0 {
			continue
		}
		totalPrice += uint(quantity) * r.priceChecker.GetItemPrice(name)
	}
	return totalPrice
}

// copyBasket returns a copy of the repo's basket. This allows us to remove
// items from it as we apply offers in the GetTotalPrice method but still keep
// the method idempotent.
func (r Repository) copyBasket() entities.Basket {
	dst := make(entities.Basket, len(r.basket))
	for k, v := range r.basket {
		dst[k] = v
	}
	return dst
}

func removeDealItems(basket entities.Basket, dealItems entities.Basket) {
	for dealItemName, dealItemQuantity := range dealItems {
		basket[dealItemName] = basket[dealItemName] - dealItemQuantity
	}
}
