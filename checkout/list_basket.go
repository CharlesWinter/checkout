package checkout

func (r *Repository) ListBasketItems() map[string]uint {
	return r.basket
}
