package transport

import (
	"fmt"
	"time"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task3/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

type LocationModel struct {
	callBack           tea.Cmd
	isConnectionOpened bool
	currentLocation    *pkg.LocationPayload
	finish             chan struct{}
}

func NewLocationModel(callBack tea.Cmd) LocationModel {
	m := LocationModel{
		callBack:        callBack,
		finish:          make(chan struct{}),
		currentLocation: &pkg.LocationPayload{},
	}
	return m
}

func (m LocationModel) Init() tea.Cmd {
	return locationCmd(m.currentLocation)
}

func (m LocationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.isConnectionOpened {
		conn := pkg.InitWebSocketConention()
		locationHandler := pkg.NewLocationHandler(conn)
		go locationHandler.StartSendingLocation(m.currentLocation, m.finish)
		m.isConnectionOpened = true
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.finish <- struct{}{}
			return m, m.callBack
		}
	case LocationMsg:
		*m.currentLocation = pkg.LocationPayload(msg)
	}

	return m, locationCmd(m.currentLocation)
}

func (m LocationModel) View() string {
	return fmt.Sprintf("Current Location is X: %v, Y: %v", m.currentLocation.X, m.currentLocation.Y)
}

type LocationMsg pkg.LocationPayload

func locationCmd(lc *pkg.LocationPayload) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second)
		return LocationMsg{
			X: lc.X,
			Y: lc.Y,
		}
	}
}

func tick() tea.Msg {
	return LocationMsg{
		X: 1,
		Y: 1,
	}
}
