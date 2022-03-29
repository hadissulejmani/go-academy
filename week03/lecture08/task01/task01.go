package main

import (
	"cardgame"
	"fmt"
)

func main() {
	cards := []cardgame.Card{
		*cardgame.NewCard(1, 2),
		*cardgame.NewCard(2, 2),
		*cardgame.NewCard(3, 4),
		*cardgame.NewCard(4, 1),
		*cardgame.NewCard(1, 2),
	}

	d := cardgame.NewDeck(cards)

	fmt.Println(d)
}
