package main

import (
	"fmt"

	"github.com/nagurumalab/gocards/gocards"
)

func main() {
	//r := gin.Default()

	var pile gocards.Pile
	fmt.Println(pile)
	ace_spade := gocards.Card{Suit: gocards.SPADES, Number: gocards.ACE}
	pile.AddCardLast(ace_spade)
	pile.AddCardFirst(ace_spade)
	fmt.Println(pile)
	pile.ShowAllCards()
	fmt.Println(pile)
}
