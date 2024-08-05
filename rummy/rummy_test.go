package rummy_test

import (
	"github.com/google/uuid"
	"testing"

	gc "github.com/nagurumalab/gocards/gocards"
	rmy "github.com/nagurumalab/gocards/rummy"
)

func TestRummy(t *testing.T) {
	player_1 := gc.Player{Id: "1", Name: "Bala"}
	player_2 := gc.Player{Id: "2", Name: "Murugan"}
	newGame := rmy.NewRummy(map[string]gc.Player{player_1.Id: player_1, player_2.Id: player_2})
	newGame.Start()
	newGame.HandleEvent(
		gc.TakeCard{
			Event:    gc.Event{Player: player_1, Id: uuid.NewString()},
			FromPile: "Closed",
			FromTop:  true,
			NumCards: 1,
		})
}
