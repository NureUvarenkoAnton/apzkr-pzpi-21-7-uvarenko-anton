package transport

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type errMsg error

type inputModel struct {
	textInput textinput.Model
	title     string
	callBack  func(string) tea.Cmd
	err       error
}

func NewInputModel(title, placeHodler string, callBack func(string) tea.Cmd) inputModel {
	ti := textinput.New()
	ti.Placeholder = "Api key"
	ti.Focus()
	ti.CharLimit = 1000
	ti.Width = 200

	return inputModel{
		textInput: ti,
		callBack:  callBack,
		err:       nil,
		title:     title,
	}
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, m.callBack(m.textInput.Value())
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf(
		"%s:\n\n%s\n\n%s",
		m.title,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
