package rummy

import (
	"errors"

	gc "github.com/nagurumalab/gocards/gocards"
	"github.com/rs/zerolog/log"
)

const (
	CLS = "Closed"
	DSC = "Discarded"
)

type Rummy struct {
	gc.Game
	JokerCard gc.Card
}

func NewRummy(players map[string]gc.Player) Rummy {
	log.Info().Msg("Init a new Rummy game")
	game := gc.NewGame(2, players)
	log.Info().Msg("Initialized the game")
	r := Rummy{Game: game}
	r.CardPiles = map[string]*gc.Pile{}
	r.CardPiles[CLS] = &gc.Pile{}
	r.CardPiles[DSC] = &gc.Pile{}
	for pId := range players {
		r.CardPiles[pId] = &gc.Pile{}
	}
	log.Info().Msg("Initialized the cardpiles")
	return r
}

func (r *Rummy) DealCards(fromPile *gc.Pile, numCards int) {
	for i := range numCards {
		_ = i
		for playerId := range r.Players {
			c, _ := fromPile.GetCardEnd()
			r.CardPiles[playerId].PutCardEnd(c)
		}
	}
}

func (r *Rummy) TakeCard() gc.Card {
	card, err := r.CardPiles[CLS].GetCardEnd()
	if err != nil {
		log.Error().Err(err)
		if errors.Is(err, gc.ErrRemoveEmptyPile) {
			top_card, _ := r.CardPiles[DSC].GetCardEnd()
			r.CardPiles[CLS], r.CardPiles[DSC] = r.CardPiles[DSC], r.CardPiles[CLS]
			r.CardPiles[CLS].HideAllCards()
			r.CardPiles[CLS].Shuffle()
			r.CardPiles[DSC].PutCardEnd(top_card)
			card, _ := r.CardPiles[CLS].GetCardEnd()
			return card
		}
		// FIXME: Handle other Error if possible
		return gc.Card{}
	}
	return card
}

func (r *Rummy) DropCard(card gc.Card) {
	card.Show = true
	r.CardPiles[DSC].PutCardEnd(card)
}

func (r *Rummy) Start() {
	log.Info().Msg("Starting rummy")
	log.Info().Msg("Getting Decks")
	decks := gc.GetDecks(r.NumDecks, r.NumDecks)
	log.Info().Msg("Shuffling Decks")
	decks.Shuffle()
	// deal cards
	r.CardPiles["Closed"], r.CardPiles[CLS] = &decks, &decks
	r.DealCards(r.CardPiles[CLS], 13)

	r.JokerCard, _ = r.CardPiles[CLS].GetCardEnd()

	openCard, _ := r.CardPiles[CLS].GetCardEnd()
	openCard.Show = true
	r.CardPiles[DSC].PutCardEnd(openCard)

	r.State.NextMove = "START"
	log.Info().Msg("Game Started")
	//fmt.Println("Game started")
	log.Debug().Msgf("%s", r)
	// fmt.Println(r)
}

func (r *Rummy) HandleEvent(event interface{}) bool {
	switch event := event.(type) {
	case gc.TakeCard:
		log.Debug().Msgf("Take Card Event has been passed - %v", event)
		card, _ := r.CardPiles[CLS].GetCardEnd()
		r.CardPiles[event.Event.Player.Id].PutCardEnd(card)
	default:
		log.Error().Msgf("Unsupported Event - %s", event)
	}
	return false
}
