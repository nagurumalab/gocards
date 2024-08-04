package gameserver

import "github.com/gin-gonic/gin"

type GameRoom struct {
	Id        string
	Name      string
	players   map[string]Player
	broadcast chan interface{}
}

func (s *GameRoom) Run() {
	for {
		select {}
	}
}

func (r *GameRoom) JoinRoom(c gin.Context) {

}
