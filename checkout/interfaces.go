package checkout

//go:generate mockgen -destination=mocks_test.go -package=checkout_test . PriceChecker

// PriceChecker is a simple abstraction around an entity capable of providing information on the price of items
type PriceChecker interface {
	GetItemPrice(name string) uint
}
