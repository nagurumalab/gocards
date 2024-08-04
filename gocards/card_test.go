package gocards_test

import (
	"testing"

	"github.com/nagurumalab/gocards/gocards"
)

func TestAddCardToPile(t *testing.T) {
	pile := gocards.Pile{}
	new_card := gocards.Card{Suit: gocards.SPADES, Number: gocards.ACE}
	pile.PutCardEnd(new_card)
	if pile[pile.Len()-1] != new_card {
		t.Errorf("Card not added to end of the pile")
	}
}
