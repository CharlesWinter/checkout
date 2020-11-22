package entities

// Deal is a struct representing an applicable discount for a set of items
type Deal struct {
	RequiredSubbasket Basket

	Price uint
}

// IsApplicable checks if the passed in basket meets the deals requirements
func (d Deal) IsApplicable(basket Basket) bool {
	// for every item the deal requires, check if that item is present in the
	// basket, and the amount is equal to or greater than the required quantity
	for dealItemName, dealItemQuantity := range d.RequiredSubbasket {
		if q, ok := basket[dealItemName]; !ok || q < dealItemQuantity {
			return false
		}
	}

	return true
}

// GetPrice simply returns the price of the deal
func (d Deal) GetPrice() uint {
	return d.Price
}

// GetItems returns the items that the deal requires to be valid
func (d Deal) GetItems() Basket {
	return d.RequiredSubbasket
}
