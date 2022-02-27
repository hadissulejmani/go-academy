package main

import "fmt"

type Card struct {
	cardSuit  int
	cardValue int
}

func main() {
	cardOne := Card{cardSuit: 3, cardValue: 3}
	cardTwo := Card{cardSuit: 4, cardValue: 3}

	cards := make([]Card, 5)
	cards = append(cards, cardOne)
	cards = append(cards, cardTwo)

	resultFromFunction := maxCard(cards)
	fmt.Println(resultFromFunction)
}

func maxCard(cards []Card) Card {
	maxValue := 0
	maxSuit := 0

	for i, card := range cards {
		if card.cardSuit >= maxSuit && card.cardValue > maxValue {
			maxValue = i
		}
	}
	return cards[maxValue]
}
