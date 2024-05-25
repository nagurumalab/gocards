package gocards

import "testing"

func TestAddCardToPile(t *testing.T) {
	pile := NewPile()
	new_card := Card{Suit: SPADES, Number: ACE}
	pile.AddCardLast(new_card)
	if pile.Cards[pile.Len()-1] != new_card {
		t.Errorf("Card not added to end of the pile")
	}
}
