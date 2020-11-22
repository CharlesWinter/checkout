package checkout

// GetTotalPrice calculates and returns the total price of all items in the basket.
func (r Repository) GetTotalPrice() uint {
	var totalPrice uint
	for name, quantity := range r.basket {
		totalPrice += uint(quantity) * r.priceChecker.GetItemPrice(name)
	}
	return totalPrice
}
