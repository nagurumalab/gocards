package gameserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type joinRoomParams struct {
	roomId     string `uri:"roomId" binding:"required"`
	playerId   string `form:"playerId"`
	playerName string `form:"playerName"`
}

func (h *Hub) JoinRoom(c *gin.Context) {
	var input joinRoomParams
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	log.Debug().Msgf("Input %v", input)

	room, ok := h.rooms[input.roomId]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"msg": fmt.Sprintf("Room Not Found - %s", input.roomId)})
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player := Player{Id: input.playerId, Name: input.playerName, connection: conn}
	room.players[player.Id] = player
}
