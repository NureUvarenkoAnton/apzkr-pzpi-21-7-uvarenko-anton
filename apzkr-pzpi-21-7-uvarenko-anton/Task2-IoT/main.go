package main

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task3/transport"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// conn := InitWebSocketConention()
	// locaitonHandler := NewLocationHandler(conn)
	// locaitonHandler.StartSendingLocation()

	m := transport.NewCompositeModel()
	tea.NewProgram(m).Run()
}
