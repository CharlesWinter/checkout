package checkout

import "github.com/CharlesWinter/checkout/entities"

func (r *Repository) Scan(itemName entities.ItemName) {
	r.basket[itemName]++
}
