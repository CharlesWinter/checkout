package checkout

import "github.com/CharlesWinter/checkout/entities"

func (r *Repository) ListBasketItems() entities.Basket {
	return r.basket
}
