package gocards

import (
	"fmt"
	"log"
)

type Config struct {
	NumDecks int `json:"num_decks"`
}

func (c Config) String() string {
	return fmt.Sprintf("(Config : %d)", c.NumDecks)
}

type State struct {
	Desc     string `json:"description"`
	NextMove string `json:"next_move"`
}

func (s State) String() string {
	return fmt.Sprintf("(State : %s - %s)", s.Desc, s.NextMove)
}

type Game struct {
	Config
	CardPiles map[string]*Pile `json:"card_piles"`
	State
	Players map[string]Player `json:"players"`
}

func (g Game) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s", g.Players, g.CardPiles, g.Config, g.State)
}

func NewGame(num_decks int, players map[string]Player) Game {
	return Game{Config: Config{NumDecks: 2}, Players: players}
}

// Starting a game usually means the following
// 1. Shuffle
// 2. Split the decks into different table piles
// 3. Deal the cards to players from one of table piles
// 4. Assign a player action to one of the player
func (game *Game) Start() {
	log.Println("Starting a game")
}

func (game *Game) CurrentState() {

}

func (game *Game) ActionHandler() {
	log.Println("Starting the game and running the game loop")

}
