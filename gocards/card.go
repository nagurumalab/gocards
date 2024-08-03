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

func (cs Suit) String() string {
	suites := [...]string{"-", "♠", "♥", "♦", "♣"}
	if cs < NOSUITE || cs > CLUBS {
		return "-"
	}
	return suites[cs]
}

type Number int

const (
	JOKER Number = iota
	ACE
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

func (cn Number) String() string {
	pip := [...]string{"🃟", "A", "2", "3", "4", "5", "6",
		"7", "8", "9", "10", "J", "Q", "K"}
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
	show := "0"
	if !c.Show {
		show = "#"
	}
	return fmt.Sprintf("|%s%s-%s|", c.Suit, c.Number, show)
}

func (c Card) Eq(card Card) bool {
	if c.Number == card.Number && c.Suit == card.Suit {
		return true
	}
	return false
}
