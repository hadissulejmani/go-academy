package carddraw

import (
	"cardgame"
)

type dealer interface {
	Deal() *cardgame.Card
}

type Dealer []dealer

func (d Dealer) Deal() *cardgame.Card {
	return &cardgame.Card{}
}

func DrawAllCards(dealer Dealer) []cardgame.Card {
	// call the dealer's Draw() method, until you reach a nil Card
	cards := []cardgame.Card{}
	for _, value := range dealer {
		cards = append(cards, *value.Deal())
	}

	return cards
}
