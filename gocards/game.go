package gocards

type Config struct {
	NumDecks int `json:"num_decks"`
}

type State struct {
	Desc       string `json:"description"`
	NextAction string `json:"next_action"`
}

type Game struct {
	Config
	CardPiles map[string]Pile `json:"card_piles"`
	State
	Players map[string]Player `json:"players"`
}

func newGame() *Game {
	return &Game{}
}

func (game *Game) Start() {

}

func (game *Game) CurrentState() {

}
