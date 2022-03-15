package main

import (
	"carddraw"
	"cardgame"
	"fmt"
)

func main() {
	deck := cardgame.Deck{}

	dealer := carddraw.Dealer{}

	for i := 0; i < 10; i++ {
		dealer = append(dealer, cardgame.NewCard(int(i), int(i)))
	}

	carddraw.DrawAllCards(dealer)
	fmt.Println(deck)
}
