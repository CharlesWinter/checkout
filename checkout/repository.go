package checkout

import "github.com/CharlesWinter/checkout/entities"

// Repository defines a type capable of affording basic checkout capability
type Repository struct {
	priceChecker PriceChecker
	dealGetter   DealGetter

	basket entities.Basket
}

// RepositoryConfig is the config struct for the repostitory
type RepositoryConfig struct {
	PriceChecker PriceChecker
	DealGetter   DealGetter
}

// New returns a new Repository
func New(cfg RepositoryConfig) *Repository {
	return &Repository{
		priceChecker: cfg.PriceChecker,
		dealGetter:   cfg.DealGetter,
		basket:       make(entities.Basket),
	}
}
