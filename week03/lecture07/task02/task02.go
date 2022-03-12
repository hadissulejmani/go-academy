package main

import "fmt"

type Card struct {
	cardSuite int
	cardValue int
}

type CardComparator func(cOne Card, cTwo Card) int

func (cardComparator CardComparator) compare(cardOne, cardTwo Card) int {
	if cardOne.cardSuite > cardTwo.cardSuite {
		return 1
	} else if cardOne.cardSuite == cardTwo.cardSuite && cardOne.cardValue == cardTwo.cardValue {
		return 0
	} else {
		return -1
	}
}

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	maxCard := Card{cardSuite: 0, cardValue: 0}
	for _, card := range cards {
		if comparatorFunc.compare(card, maxCard) > 0 {
			maxCard = card
		}
	}
	return maxCard
}

func main() {
	cards := []Card{
		{cardSuite: 2, cardValue: 1},
		{cardSuite: 4, cardValue: 1},
		{cardSuite: 6, cardValue: 1},
		{cardSuite: 1, cardValue: 1},
		{cardSuite: 2, cardValue: 1},
	}

	fmt.Println(maxCard(cards, func(cardOne, cardTwo Card) int { return 0 }))
}
