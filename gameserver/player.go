package gameserver

import (
	"github.com/gorilla/websocket"
	gc "github.com/nagurumalab/gocards/gocards"
	"github.com/rs/zerolog/log"
)

type Player struct {
	Id         string
	Name       string
	conn       *websocket.Conn
	fromPlayer chan gc.Event
}

func (p Player) SendMsg(e gc.Event) error {
	if err := p.conn.WriteJSON(e); err != nil {
		log.Error().Err(err).Msgf("Something went wrong when sending message to Player - %s : %s", p.Id, e)
		return err
	}
	return nil
}

func (p Player) RecvMsg() {
	for {
		_, msg, err := p.conn.ReadMessage()
		if err != nil {
			log.Error().Err(err).Msgf("Something went wrong when reading from websocket for player - %s", p.Id)
			break
		}

	}
}
