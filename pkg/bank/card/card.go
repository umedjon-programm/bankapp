package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource{
	
	var ps []types.PaymentSource
	for _,payment:=range cards{
		if payment.Balance>0 && payment.Active{
	ps = append(ps, types.PaymentSource{Type:payment.Name,Balance:payment.Balance,Number:string(payment.Pan)})

		}
	}
	return ps
}