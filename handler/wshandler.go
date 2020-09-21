package handler

import (
	"time"

	"github.com/danielthank/exchat-server/infra"
	"github.com/danielthank/websocket-pubsub"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	*infra.RedisHandler
	*websocket.Upgrader
	*wspubsub.Config
}

func NewWSHandler(redisHandler *infra.RedisHandler) *WSHandler {
	return &WSHandler{redisHandler,
		&websocket.Upgrader{},
		&wspubsub.Config{
			PingPeriod:     50 * time.Second,
			PongWait:       60 * time.Second,
			WriteWait:      10 * time.Second,
			MaxMessageSize: 8192,
		}}
}

func (t *WSHandler) Handle(c *gin.Context) {
	wsConn, err := t.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	wsPubSub := &wspubsub.WSPubSub{
		WSConn:      wsConn,
		RedisClient: t.Client,
		Config:      t.Config,
	}

	go wsPubSub.Run()

}
