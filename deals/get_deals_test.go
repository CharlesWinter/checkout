package deals_test

import (
	"testing"

	"github.com/CharlesWinter/checkout/deals"
	"github.com/CharlesWinter/checkout/entities"
)

func TestGetDeals(t *testing.T) {
	t.Run("returns a list of deals", func(t *testing.T) {
		dealA := entities.Deal{
			Price: 10,
		}
		dealB := entities.Deal{
			Price: 20,
		}

		repository := deals.Repository{
			Deals: []entities.Deal{dealA, dealB},
		}

		got := repository.GetDeals()

		// This test can be improved through the introduction of deep equals or
		// even better, an external package like cmp Diff. However for the purpose
		// of this exercise lets be content to compare the price
		if got[0].Price != dealA.Price {
			t.Fatalf("expected deal %#v, got deals %#v", dealA, got)
		}

		if got[1].Price != dealB.Price {
			t.Fatalf("expected deal %#v, got deals %#v", dealB, got)
		}
	})
}
