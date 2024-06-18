package pkg

import (
	"fmt"
	"log"
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task3/model"

	"github.com/gorilla/websocket"
)

func InitWebSocketConention() *websocket.Conn {
	header := http.Header{}
	header.Add("Authorization", "Bearer "+model.ApiKey)
	conn, _, err := websocket.
		DefaultDialer.
		Dial("ws://localhost:8080/ws", header)
	if err != nil {
		log.Fatal(fmt.Errorf("can't open connection: [%w]", err))
	}

	return conn
}
