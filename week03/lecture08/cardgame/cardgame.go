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

func (d *Deck) Done() bool {
	return d.cards == nil
}

func (d *Deck) Draw() (*Card, error) {
	if d == nil {
		return nil, nil
	}
	card := &d.cards[0]
	d.cards = append(d.cards[:0], d.cards[1:]...)
	return card, nil
}
