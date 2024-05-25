package gocards

import (
	"fmt"
)

type Suit int

const (
	NOSUITE Suit = iota
	SPADES
	HEARTS
	DIAMONDS
	CLUBS
)
const NOSUIT Suit = -1

func (cs Suit) String() string {
	suites := [...]string{"-", "♠", "♥", "♦", "♣"}
	if cs < NOSUITE || cs > CLUBS {
		return "-"
	}
	return suites[cs]
}

type Number int

const (
	ACE Number = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

const JOKER Number = -1

func (cn Number) String() string {
	pip := [...]string{"🃟", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	if cn < JOKER || cn > KING {
		return "<INVALID>"
	}
	return pip[cn]
}

type Card struct {
	Suit   Suit
	Number Number
	Show   bool
}

func (c Card) String() string {
	if !c.Show {
		return "|##|"
	}
	return fmt.Sprintf("|%s%s|", c.Suit, c.Number)
}

type Pile struct {
	Cards []Card
}

func NewPile() *Pile {
	return &Pile{}
}

func (pile *Pile) AddCard(card Card, idx int) {
	if idx == -1 {
		// log.Printf("idx is %d", idx)
		idx = len(pile.Cards)
	}
	if len(pile.Cards) == idx {
		pile.Cards = append(pile.Cards, card)
		// log.Print("idx is ", pile)
	} else {
		pile.Cards = append(pile.Cards[:idx+1], pile.Cards[idx:]...)
		pile.Cards[idx] = card
	}

}

func (pile *Pile) AddCardLast(card Card) {
	pile.AddCard(card, -1)
}

func (pile *Pile) AddCardFirst(card Card) {
	pile.AddCard(card, 0)
}

func (pile *Pile) RemoveCard(idx int) {
	if idx < len(pile.Cards) {
		pile.Cards = append(pile.Cards[:idx], pile.Cards[idx+1:]...)
	}
}

func (pile *Pile) ShowCard(idx int) {
	pile.Cards[idx].Show = true
}

func (pile *Pile) ShowCards(idxs []int) {
	for _, idx := range idxs {
		pile.ShowCard(idx)
	}
}

func (pile *Pile) ShowAllCards() {
	for idx := range pile.Cards {
		pile.ShowCard(idx)
	}
}

func (a Pile) Len() int      { return len(a.Cards) }
func (a Pile) Swap(i, j int) { a.Cards[i], a.Cards[j] = a.Cards[j], a.Cards[i] }
func (a Pile) Less(i, j int) bool {
	return a.Cards[i].Suit < a.Cards[j].Suit && a.Cards[i].Number < a.Cards[j].Number
}

func GetDecks(n int) {

}
