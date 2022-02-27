package main

import "fmt"

type CardSuit = int

const (
	club CardSuit = iota + 2
	diamond
	heart
	spade
)

func main() {
	resultFromFunction := compareCards(2, heart, 2, diamond)
	fmt.Println(resultFromFunction)
}

func compareCards(cardOneVal int, cardOneSuit CardSuit, cardTwoVal int, cardTwoSuit CardSuit) int {
	if cardOneVal > cardTwoVal {
		return 1
	} else if cardTwoVal > cardOneVal {
		return -1
	} else {
		if cardOneSuit > cardTwoSuit {
			return 1
		} else if cardOneSuit < cardTwoSuit {
			return -1
		} else {
			return 0
		}
	}
}
