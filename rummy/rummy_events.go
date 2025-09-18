package rummy

import (
	gc "github.com/nagurumalab/gocards/gocards"
	"github.com/rs/zerolog/log"
)

type TakeCard struct {
	gc.TakeCard
}

func (event TakeCard) Handle(g gc.Game) {
	r := g.(*Rummy)
	log.Debug().Msgf("TakeCard - %v", event)
	card, _ := r.CardPiles[CLS].GetCardEnd()
	r.CardPiles[event.Player.Id].PutCardEnd(card)
}

type DropCard struct {
	gc.DropCard
}

func (event DropCard) Handle(g gc.Game) {
	r := g.(*Rummy)
	log.Debug().Msgf("DropCard event - %v", event)
	found, idx := r.CardPiles[event.Player.Id].SearchCard(event.Card)
	if !found {
		log.Error().Msgf("Dropping card not in hand")
	} else {
		card, _ := r.CardPiles[event.Player.Id].GetCard(idx)
		r.DropCard(card)
	}
}
