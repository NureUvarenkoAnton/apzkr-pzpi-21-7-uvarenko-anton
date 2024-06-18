package transport

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task3/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	generalListDisplay = iota
	keyDisplay
	sendLocationDisplay
)

type compositeModel struct {
	listModel     tea.Model
	inputModel    tea.Model
	locationModel tea.Model
	toDisplay     int
}

func NewCompositeModel() compositeModel {
	compositeModel := compositeModel{
		toDisplay: generalListDisplay,
	}

	listItems := []list.Item{
		compositeModel.initAddApiKeyItem(),
		compositeModel.initStartSendingLocation(),
	}
	compositeModel.listModel = NewListModel(listItems)
	compositeModel.inputModel = NewInputModel("Add key", "key", func(data string) tea.Cmd {
		return func() tea.Msg {
			model.ApiKey = data
			return switchViewMsg(generalListDisplay)
		}
	})
	compositeModel.locationModel = NewLocationModel(func() tea.Msg {
		return switchViewMsg(generalListDisplay)
	})

	return compositeModel
}

func (m compositeModel) Init() tea.Cmd {
	return nil
}

func (m compositeModel) View() string {
	switch m.toDisplay {
	case generalListDisplay:
		return m.listModel.View()
	case keyDisplay:
		return m.inputModel.View()
	case sendLocationDisplay:
		return m.locationModel.View()
	}
	return ""
}

func (m compositeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case switchViewMsg:
		m.toDisplay = int(msg)
		return m, cmd
	default:
		switch m.toDisplay {
		case generalListDisplay:
			m.listModel, cmd = m.listModel.Update(msg)
		case keyDisplay:
			m.inputModel, cmd = m.inputModel.Update(msg)
		case sendLocationDisplay:
			m.locationModel, cmd = m.locationModel.Update(msg)
		}
	}

	return m, cmd
}

type switchViewMsg int

func (m *compositeModel) initAddApiKeyItem() Item {
	message := "Add api key"
	return Item{
		message: &message,
		callBack: func() tea.Msg {
			return switchViewMsg(keyDisplay)
		},
	}
}

func (m *compositeModel) initStartSendingLocation() Item {
	message := "Sending location"
	return Item{
		message: &message,
		callBack: func() tea.Msg {
			return switchViewMsg(sendLocationDisplay)
		},
	}
}
