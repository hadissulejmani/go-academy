package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	cardSuit  int
	cardValue int
}

type Deck struct {
	cards []Card
}

func (deck *Deck) New() {
	for i := 1; i < 14; i++ {
		for cardSuit := 0; cardSuit < 4; cardSuit++ {
			card := Card{
				cardSuit:  cardSuit,
				cardValue: i,
			}
			deck.cards = append(deck.cards, card)
		}
	}
	return
}

func (deck *Deck) Shuffle() {
	for i := 1; i < 14; i++ {
		r := rand.Intn(i)
		if i != r {
			deck.cards[r], deck.cards[i] = deck.cards[i], deck.cards[r]
		}
	}
}

func (deck *Deck) Deal() (card *Card) {
	if deck.cards == nil {
		return
	}
	card = &deck.cards[0]
	deck.cards = append(deck.cards[1:])
	return card
}

func main() {
	deck := Deck{}
	deck.New()

	fmt.Println(deck.cards)

	deck.Shuffle()
	fmt.Println(*deck.Deal())
	fmt.Println(deck.cards)
}
