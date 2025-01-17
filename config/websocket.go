package config

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader websocket.Upgrader

func Websocket() {
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins (for testing purposes, but adjust for production)
		},
	}
}
