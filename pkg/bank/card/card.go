package card

import (
	"github.com/nekruz08/bank24/pkg/bank/types"
)

// Issue создаёт экземпляр карты с определёнными полями
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:         1000,
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    0,
		Currency:   currency,
		Color:      color,
		Name:       name,
		Active:     true,
		MinBalance: 0,
	}
}

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
