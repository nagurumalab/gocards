package gocards

type Session struct {
	Game    *Game
	Players map[string]Player
}

func newSession(players *map[string]Player) Session {
	return Session{Players: *players}
}

func (s *Session) createNewGame() {

}
