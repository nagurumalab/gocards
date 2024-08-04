package gameserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Hub struct {
	rooms map[string]*GameRoom
}

func NewHub() *Hub {
	return &Hub{rooms: map[string]*GameRoom{}}
}

type createRoomReq struct {
	name string `json:"name", binding:"required"`
}

func (h *Hub) CreateRoom(c *gin.Context) {
	var input createRoomReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room := &GameRoom{Name: input.name, Id: uuid.NewString()}
	h.rooms[room.Id] = room
	c.JSON(http.StatusOK, gin.H{"room_id": room.Id})
}
