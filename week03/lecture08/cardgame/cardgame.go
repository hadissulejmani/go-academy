package cardgame

type Card struct {
	cardSuit  int
	cardValue int
}

type Deck struct {
	cards []Card
}

func NewCard(cardSuit int, cardValue int) *Card {
	return &Card{cardSuit: cardSuit, cardValue: cardValue}
}

func NewDeck(cards []Card) *Deck {
	return &Deck{cards: cards}
}
