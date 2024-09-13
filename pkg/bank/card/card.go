package card

import "github.com/nekruz08/bank24/pkg/bank/types"

// Withdraw снимает деньги с краты
func Withdraw(card *types.Card, amount types.Money) {
	const WithdrawLimit = 20_000_00
	if amount < 0 {
		return
	}
	if amount > WithdrawLimit {
		return
	}
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}

	card.Balance = card.Balance - amount
}
