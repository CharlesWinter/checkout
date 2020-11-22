package deals

import "github.com/CharlesWinter/checkout/entities"

// Repository in this instance is a small, pre-emptive abstraction to return
// deals to the checkout. Its purpose is to facilitate the later removal of the
// deal functionality to somewhere else, if required.
type Repository struct {
	Deals []entities.Deal
}

// GetDeals returns the deals from the repository
func (r Repository) GetDeals() []entities.Deal {
	return r.Deals
}
