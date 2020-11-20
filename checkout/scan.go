package checkout

func (r *Repository) Scan(itemName string) {
	r.basket[itemName]++
}
