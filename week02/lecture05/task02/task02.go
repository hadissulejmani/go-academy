package main

import "fmt"

type Card struct {
	cardSuit  int
	cardValue int
}

func main() {
	cardOne := Card{cardSuit: 3, cardValue: 3}
	cardTwo := Card{cardSuit: 4, cardValue: 3}
	cardThree := Card{cardSuit: 5, cardValue: 2}

	cards := make([]Card, 5)
	cards = append(cards, cardOne)
	cards = append(cards, cardTwo)
	cards = append(cards, cardThree)

	resultFromFunction := maxCard(cards)
	fmt.Println(resultFromFunction)
}

func maxCard(cards []Card) Card {
	maxCard := Card{cardSuit: 0, cardValue: 0}

	for _, card := range cards {
		if card.cardSuit >= maxCard.cardSuit && card.cardValue >= maxCard.cardValue {
			maxCard.cardSuit = card.cardSuit
			maxCard.cardValue = card.cardValue
		}
	}
	return maxCard
}
