package main

import (
        "fmt"

        "github.com/nagurumalab/gocards/gocards"
)

func main() {
        //r := gin.Default()

        var pile gocards.Pile
        fmt.Println(pile)
        ace_spade := gocards.Card{Suit: gocards.SPADES, Number: "1"}
        pile.AddCard(ace_spade, -1)
        fmt.Println(pile)
}
