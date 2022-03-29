package carddraw

import (
	"cardgame"
	"errors"
)

type dealer interface {
	Draw() *cardgame.Card
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	// call the dealer's Draw() method, until you reach a nil Card
	cards := []cardgame.Card{}
	if dealer.Done() {
		return nil, errors.New("Deck is empty")
	}
	for val := dealer.Draw(); val != nil; {
		cards = append(cards, *val)
	}
	return cards, nil
}
