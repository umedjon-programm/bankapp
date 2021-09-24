package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	payments := []types.Card{
		{
			Id:      1,
			Balance: 10000,
			Name:    "card",
			Active:  true,
			PAN:     "5058 2702 XXXX 0001",
		},
		{
			Id:      2,
			Balance: 302323,
			Name:    "card",
			Active:  false,
			PAN:     "5058 2702 XXXX 0002",
		},
		{
			Id:      3,
			Balance: 11212,
			Name:    "card",
			Active:  true,
			PAN:     "5058 2702 XXXX 0003",
		},
		{
			Id:      4,
			Balance: -10,
			Name:    "card",
			Active:  true,
			PAN:     "5058 2702 XXXX 0004",
		},
	}
	for _, p := range PaymentSources(payments) {
		fmt.Println(p.Number)
	}
	//Output:
	//5058 2702 XXXX 0001
	//5058 2702 XXXX 0003
}
