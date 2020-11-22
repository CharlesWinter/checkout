package checkout_test

import (
	"fmt"

	"github.com/CharlesWinter/checkout/checkout"
	"github.com/CharlesWinter/checkout/deals"
	"github.com/CharlesWinter/checkout/entities"
	"github.com/CharlesWinter/checkout/prices"
)

func Example() {
	var (
		itemA entities.ItemName = "itemA"
		itemB entities.ItemName = "itemB"
		itemC entities.ItemName = "itemB"

		itemAPrice uint = 30
		itemBPrice uint = 40
		itemCPrice uint = 50
	)

	// Load up a deals repository with a single deal; in this instance, get item
	// A and item B for 10
	dealsRepository := deals.Repository{
		Deals: []entities.Deal{
			{
				RequiredSubbasket: entities.Basket{
					itemA: 1,
					itemB: 1,
				},
				Price: 10,
			},
		},
	}

	// Load up a pricing repository with information about the prices of the
	// various items.
	pricesRepository := prices.Repository{
		Prices: map[entities.ItemName]uint{
			itemA: itemAPrice,
			itemB: itemBPrice,
			itemC: itemCPrice,
		},
	}

	// Inject the deal repository and the price repository
	checkoutRepository := checkout.New(checkout.RepositoryConfig{
		DealGetter:   dealsRepository,
		PriceChecker: pricesRepository,
	})

	// Scan some items and get the total price
	checkoutRepository.Scan(itemA)
	checkoutRepository.Scan(itemA)

	checkoutRepository.Scan(itemB)
	checkoutRepository.Scan(itemB)

	checkoutRepository.Scan(itemC)

	fmt.Printf("the total price at checkout is: %d\n", checkoutRepository.GetTotalPrice())
}
