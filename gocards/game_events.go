package gocards

type Event interface {
	Handle(Game)
}

func EventToJson(e Event) string {

}

func JsonToEvent(j string, eventMap map[string]Event)

type BaseEvent struct {
	Id     string
	Player Player
}

type TakeCard struct {
	BaseEvent
	FromPile string
	FromTop  bool
	NumCards int
}

type DropCard struct {
	BaseEvent
	ToPile string
	Card   Card
}

type ShowCards struct {
	BaseEvent
	Cards     []Card
	ToPlayers []Player
}

type SeenCards struct {
	BaseEvent
	Show ShowCards
}

type AskCard struct {
	BaseEvent
	ToPlayer Player
	Card     Card
}

type DenyAsk struct {
	BaseEvent
	Ask AskCard
}

type GiveCard struct {
	BaseEvent
	ToPlayer Player
}

type GetCard struct {
	BaseEvent
	FromPlayer Player
}

type DenyGive struct {
	BaseEvent
	Give GiveCard
}

type Declare struct {
	BaseEvent
}
