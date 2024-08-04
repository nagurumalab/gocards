package gameserver

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Id         string
	Name       string
	connection *websocket.Conn
}
