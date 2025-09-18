package gameserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nagurumalab/gocards/gocards"
	"github.com/rs/zerolog/log"
)

type GameRoom struct {
	Id          string
	Name        string
	players     map[string]Player
	fromPlayers chan gocards.Event
	game        gocards.Game
}

func (r *GameRoom) Broadcast(e gocards.Event) {
	for _, player := range r.players {
		player.SendMsg(e)
	}
}

func (r *GameRoom) Run() {
	for {
		select {
		case e := <-r.fromPlayers:
			r.game.EventsHandler(e)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Hub) JoinRoom(c *gin.Context) {
	roomId := c.Param("roomId")
	playerId := c.Query("playerId")
	playerName := c.Query("playerName")

	log.Debug().Msgf("Join Room Details - %s %s %s", roomId, playerId, playerName)

	room, ok := h.rooms[roomId]
	if !ok {
		log.Error().Msgf("Room Not Found %s", roomId)
		c.JSON(
			http.StatusNotFound,
			gin.H{"msg": fmt.Sprintf("Room Not Found - %s", roomId)},
		)
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player := Player{Id: playerId, Name: playerName, connection: conn}
	room.players[player.Id] = player
}
