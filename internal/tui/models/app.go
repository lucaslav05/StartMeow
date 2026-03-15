package models

import (
	style "StartMeow/internal/tui/styles"
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	history     []string
	questions   []Question
	answerField textinput.Model
	index       int
	width       int
	height      int
	styles      *style.MainStyle
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.Next()
			current.Answer = m.answerField.Value()
			m.SetAnswerValue()
			return m, nil
		case "up":
			current.Prev()
		case "down":
			current.Next()
		case "y":
			if current.QuestionType == Verify {
				log.Println("Verified!")
				return m, tea.Quit
			}
		case "n":
			if current.QuestionType == Verify {
				m.ClearAnswers()
				m.Next()
			}
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	var l string

	currentQuestion := m.questions[m.index]

	switch currentQuestion.QuestionType {
	// current question is prompt
	case Prompt:
		l = lipgloss.JoinVertical(
			lipgloss.Center,
			m.styles.Title.Render(m.questions[m.index].Question),
			m.styles.InputField.Render(
				m.answerField.View(),
			),
		)

	// current question is select
	case Select:
		// log.Println("Case: Select")
		l = lipgloss.JoinVertical(
			lipgloss.Left,
			m.styles.Title.Render(m.questions[m.index].Question),
			m.ViewSelect(),
		)

	case Verify:
		// current question is verify
		l = lipgloss.JoinVertical(
			lipgloss.Left,
			m.styles.Title.Render(m.questions[m.index].Question),
			m.ViewVerify(),
		)
	}

	m.history = append(m.history, l)
	v := tea.NewView(strings.Join(m.history, "\n"))
	v.BackgroundColor = lipgloss.Color(style.BackgroundBlack)
	v.AltScreen = true
	return v
}

func (m model) ViewSelect() string {
	var rows []string
	currentQuestion := m.questions[m.index]
	options := currentQuestion.Options
	optionIndex := currentQuestion.OptionIndex

	for i, v := range options {
		if optionIndex == i {
			rows = append(rows, m.styles.OptionSelect.Render(fmt.Sprintf("[x] %s", v)))
		} else {
			rows = append(rows, m.styles.Options.Render(fmt.Sprintf("[ ] %s", v)))
		}
	}

	return strings.Join(rows, "\n")
}

func (m model) ViewVerify() string {
	var rows []string

	for _, v := range m.questions {
		var s string
		switch v.QuestionType {
		case Select:
			s = fmt.Sprintf("%s: %s", m.styles.Title.Render(v.Question), m.styles.OptionSelect.Render(v.Options[v.OptionIndex]))
		case Prompt:
			s = fmt.Sprintf("%s: %s", m.styles.Title.Render(v.Question), m.styles.OptionSelect.Render(v.Answer))
		}

		if s != "" {
			rows = append(rows, s)
		}
	}

	s := "\nConfirm project? (Y/N):"
	rows = append(rows, s)

	return strings.Join(rows, "\n")
}

func (m *model) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (m *model) Prev() {
	if m.index > 0 {
		m.index--
	} else {
		m.index = len(m.questions) - 1
	}
}

func (m *model) SetAnswerValue() {
	currentAnswer := m.questions[m.index].Answer

	if currentAnswer != "" {
		m.answerField.SetValue(currentAnswer)
	} else {
		m.answerField.SetValue("")
	}
}

func (m *model) ClearAnswers() {
	q := m.questions

	for i := range q {
		q[i].Answer = ""
	}
}

func NewDefaultModel(questions []Question) *model {
	mainStyle := style.DefaultStyles()

	answerField := textinput.New()
	answerField.Focus()

	return &model{
		index:       0,
		questions:   questions,
		answerField: answerField,
		styles:      mainStyle,
	}
}
