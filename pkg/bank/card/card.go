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
