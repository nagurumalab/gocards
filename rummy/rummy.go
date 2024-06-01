package rummy

import (
	"fmt"
	"log"

	gc "github.com/nagurumalab/gocards/gocards"
)

type Rummy struct {
	gc.Game
	JokerCard     gc.Card
	DiscardedPile *gc.Pile
	ClosedPile    *gc.Pile
}

func NewRummy(players map[string]gc.Player) Rummy {
	log.Println("Init a new Rummy game")
	game := gc.NewGame(2, players)
	log.Println("Initialized the game")
	r := Rummy{Game: game}
	r.CardPiles = map[string]*gc.Pile{}
	r.CardPiles["Discarded"] = &gc.Pile{}
	r.DiscardedPile = r.CardPiles["Discarded"]
	r.CardPiles["Closed"] = &gc.Pile{}
	r.ClosedPile = r.CardPiles["Closed"]
	for pId := range players {
		r.CardPiles[pId] = &gc.Pile{}
	}
	log.Println("Initialized the cardpiles")
	return r
}

func (r *Rummy) DealCards(fromPile *gc.Pile, numCards int) {
	for i := range numCards {
		_ = i
		for playerId := range r.Players {
			_, c := fromPile.RemoveCard(0)
			r.CardPiles[playerId].AddCardToEnd(c)
		}
	}

}

func (r *Rummy) Start() {
	log.Println("Starting rummy")
	log.Println("Getting Decks")
	decks := gc.GetDecks(r.NumDecks, r.NumDecks)
	log.Println("Shuffling Decks")
	decks.Shuffle()
	// deal cards
	r.CardPiles["Closed"], r.ClosedPile = &decks, &decks
	r.DealCards(r.ClosedPile, 13)
	_, r.JokerCard = r.ClosedPile.RemoveCard(0)
	_, openCard := r.ClosedPile.RemoveCard(0)
	openCard.Show = true
	*r.DiscardedPile = append(*r.DiscardedPile, openCard)
	r.State.NextMove = "START"
	log.Println("Game Started")
	//fmt.Println("Game started")
	log.Println(r)
	// fmt.Println(r)
}

func (r *Rummy) getNextMove() {
	r.State.NextMove = "END"
}

func (r *Rummy) Run() {
	r.Start()
	for r.State.NextMove != "END" {
		r.getNextMove()
	}
	fmt.Println("Game Ended")
}
