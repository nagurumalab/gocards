package gocards

type Event struct {
	Id     string
	Player Player
}

type TakeCard struct {
	Event
	FromPile string
	FromTop  bool
	NumCards int
}

type DropCard struct {
	Event
	ToPile string
	Card   Card
}

type ShowCards struct {
	Event
	Cards     []Card
	ToPlayers []Player
}

type SeenCards struct {
	Event
	Show ShowCards
}

type AskCard struct {
	Event
	ToPlayer Player
	Card     Card
}

type DenyAsk struct {
	Event
	Ask AskCard
}

type GiveCard struct {
	Event
	ToPlayer Player
}

type GetCard struct {
	Event
	FromPlayer Player
}

type DenyGive struct {
	Event
	Give GiveCard
}

type Declare struct {
	Event
}
