package gocards

type Event struct {
	Player Player
}

type TakeCard struct {
	Event
	FromPile Pile
	FromTop  bool
	NumCards int
}

type DropCard struct {
	Event
	ToPile Pile
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
