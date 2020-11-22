package prices

import "github.com/CharlesWinter/checkout/entities"

// Repository in this instance is a small, pre-emptive abstraction to return
// prices to the checkout. Its purpose is to facilitate the later removal of the
// prices functionality to somewhere else, if required.
type Repository struct {
	Prices map[entities.ItemName]uint
}

// GetItemPrice returns the deals from the repository
func (r Repository) GetItemPrice(name entities.ItemName) uint {
	return r.Prices[name]
}
