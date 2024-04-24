package gocards

import (
	"fmt"
)

type CardSuit rune

const (
	HEARTS   CardSuit = '♥'
	CLUBS    CardSuit = '♣'
	DIAMONDS CardSuit = '♦'
	SPADES   CardSuit = '♠'
)

type CardNumber string

const (
	ACE   CardNumber = "1"
	TWO   CardNumber = "2"
	THREE CardNumber = "3"
	FOUR  CardNumber = "4"
	FIVE  CardNumber = "5"
	SIX   CardNumber = "6"
	SEVEN CardNumber = "7"
	EIGHT CardNumber = "8"
	NINE  CardNumber = "9"
	TEN   CardNumber = "10"
	JACK  CardNumber = "J"
	QUEEN CardNumber = "Q"
	KING  CardNumber = "K"
)

type Card struct {
	Suit   CardSuit
	Number CardNumber
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
