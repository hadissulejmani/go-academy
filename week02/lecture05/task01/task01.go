package main

import "fmt"

type Card struct {
	cardSuit  int
	cardValue int
}

func main() {
	cardOne := Card{cardSuit: 3, cardValue: 3}
	cardTwo := Card{cardSuit: 4, cardValue: 3}

	resultFromFunction := compareCards(cardOne, cardTwo)
	fmt.Println(resultFromFunction)
}

func compareCards(cardOne Card, cardTwo Card) int {
	if cardOne.cardValue > cardTwo.cardValue {
		return 1
	} else if cardTwo.cardValue > cardOne.cardValue {
		return -1
	} else {
		if cardOne.cardSuit > cardTwo.cardSuit {
			return 1
		} else if cardOne.cardSuit < cardTwo.cardSuit {
			return -1
		} else {
			return 0
		}
	}
}
