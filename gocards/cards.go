package gocards

import (
	"fmt"
)

type Suit rune

const (
	HEARTS   Suit = '♥'
	CLUBS    Suit = '♣'
	DIAMONDS Suit = '♦'
	SPADES   Suit = '♠'
)

type Number string

const (
	ACE   Number = "1"
	TWO   Number = "2"
	THREE Number = "3"
	FOUR  Number = "4"
	FIVE  Number = "5"
	SIX   Number = "6"
	SEVEN Number = "7"
	EIGHT Number = "8"
	NINE  Number = "9"
	TEN   Number = "10"
	JACK  Number = "J"
	QUEEN Number = "Q"
	KING  Number = "K"
)

type Card struct {
	Suit   Suit
	Number Number
	Show   bool
}

func (c Card) String() string {
	show_str := "-"
	if c.Show {
		show_str = "0"
	}
	return fmt.Sprintf("%c%s %s", c.Suit, c.Number, show_str)
}

type Pile struct {
	Cards []Card
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

func (pile *Pile) RemoveCard(idx int) {
	if idx < len(pile.Cards) {
		pile.Cards = append(pile.Cards[:idx], pile.Cards[idx+1:]...)
	}
}

func (pile *Pile) ShowCards(idxs []int) {
	for _, idx := range idxs {
		pile.Cards[idx].Show = true
	}
}
