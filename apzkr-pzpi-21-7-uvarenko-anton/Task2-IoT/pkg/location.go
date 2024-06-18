package pkg

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type LocationHandler struct {
	conn            *websocket.Conn
	currentLocation LocationPayload
}

func NewLocationHandler(w *websocket.Conn) *LocationHandler {
	return &LocationHandler{
		conn: w,
		currentLocation: LocationPayload{
			X: 1,
			Y: 1,
		},
	}
}

type LocationPayload struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var defaultVector = LocationPayload{
	X: 1,
	Y: 1,
}

func (h *LocationHandler) getCurrentLocation() LocationPayload {
	h.currentLocation = LocationPayload{
		X: h.currentLocation.X + defaultVector.X,
		Y: h.currentLocation.Y + defaultVector.Y,
	}

	return h.currentLocation
}

func (h *LocationHandler) StartSendingLocation(lc *LocationPayload, finish <-chan struct{}) {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			currentLocation := h.getCurrentLocation()
			h.SendLocation(currentLocation)
			*lc = currentLocation
		case <-finish:
			return
		}
	}
}

func (h *LocationHandler) SendLocation(location LocationPayload) {
	body, _ := json.Marshal(location)
	h.conn.WriteMessage(websocket.TextMessage, body)
}
