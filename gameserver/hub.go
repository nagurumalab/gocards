package gameserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nagurumalab/gocards/gocards"
)

type Hub struct {
	rooms map[string]*GameRoom
}

func NewHub() *Hub {
	return &Hub{rooms: map[string]*GameRoom{}}
}

type createRoomReq struct {
	Name string `json:"name" binding:"required"`
}

func (h *Hub) CreateRoom(c *gin.Context) {
	var input createRoomReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room := &GameRoom{
		Name:        input.Name,
		Id:          uuid.NewString(),
		players:     map[string]Player{},
		fromPlayers: make(chan gocards.Event),
	}
	h.rooms[room.Id] = room
	c.JSON(http.StatusOK, gin.H{"id": room.Id})
}

func (h *Hub) ListRooms(c *gin.Context) {
	rooms := []map[string]string{}

	for _, r := range h.rooms {
		rooms = append(rooms, map[string]string{"id": r.Id, "name": r.Name})
	}

	c.JSON(http.StatusOK, rooms)
}
