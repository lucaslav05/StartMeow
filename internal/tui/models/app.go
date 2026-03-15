package internal

import (
	"StartMeow/internal"
	style "StartMeow/internal/tui/styles"
	"fmt"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	styles *style.MainStyle
	width  int
	height int

	questions   []Question // list of questions
	project     internal.Project
	qIndex      int
	answerField textinput.Model
	help        help.Model
	keys        help.KeyMap
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	currentQ := &m.questions[m.qIndex] // get current question

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "?":
			m.help.ShowAll = !m.help.ShowAll
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			// if we are at the end of the form, do not go to next
			if currentQ.Prompt.PromptType == internal.Info {
				return m, tea.Quit
			}

			m.NextQuestion()
			currentQ.Prompt.Input = m.answerField.Value()
			m.SetAnswerValue()
			//CRITICAL SECTION
			return m, nil

		case "up":
			// if select form, then move option index
			if currentQ.Prompt.PromptType == internal.Select {
				currentQ.PrevOption()
				return m, nil
			}

		case "down":
			// if select form, then move option index
			if currentQ.Prompt.PromptType == internal.Select {
				currentQ.NextOption()
				return m, nil
			}

		case "y":
			// if we are at the end of the form, make Y a command
			if currentQ.Prompt.PromptType == internal.Info {
				return m, tea.Quit
			}

		case "n":
			// if we are at the end of the form, make N a command
			if currentQ.Prompt.PromptType == internal.Info {
				m.ClearAnswers()
				m.NextQuestion()
			}
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	var tabView string

	currentQuestion := m.questions[m.qIndex]

	switch currentQuestion.Prompt.PromptType {
	// current question is prompt
	case internal.Field:
		tabView = m.ViewTab(
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.styles.TabStyle.Render(
					m.ViewField(),
				),
			))

	// current question is select
	case internal.Select:
		// log.Println("Case: Select")
		tabView = m.ViewTab(
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.styles.TabStyle.Render(
					m.ViewSelect(),
				),
			),
		)

	case internal.Info:
		// current question is verify
		tabView = m.ViewTab(
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.styles.TabStyle.Render(
					m.ViewVerify(),
				),
			))
	}

	helpView := lipgloss.Place(
		m.width,
		2,
		lipgloss.Center,
		lipgloss.Bottom,
		m.styles.Title.Render(m.help.View(m.keys)),
	)

	finalView := lipgloss.JoinVertical(lipgloss.Center, tabView, helpView)

	v := tea.NewView(finalView)
	v.BackgroundColor = lipgloss.Color(style.BackgroundBlack)
	v.AltScreen = true
	return v
}

func (m model) ViewTab(view string) string {
	return lipgloss.Place(
		m.width,
		m.height/2,
		lipgloss.Center,
		lipgloss.Bottom,
		view,
	)
}

func (m model) ViewField() string {
	var rows []string

	rows = append(rows, m.styles.Title.Render(m.questions[m.qIndex].Prompt.Question))
	rows = append(rows, m.styles.InputField.Render(m.answerField.View()))

	return strings.Join(rows, "\n")
}

func (m model) ViewSelect() string {
	var rows []string

	rows = append(rows, m.styles.Title.Render(m.questions[m.qIndex].Prompt.Question))

	currentQuestion := m.questions[m.qIndex]
	options := currentQuestion.Prompt.Options
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

	rows = append(rows, m.styles.Title.Render(m.questions[m.qIndex].Prompt.Question))
	// rows = append(rows, "")

	for _, v := range m.questions {
		var s string
		switch v.Prompt.PromptType {
		case internal.Select:
			s = fmt.Sprintf("%s: %s", m.styles.Options.Render(v.Prompt.Question), m.styles.OptionSelect.Render(v.Prompt.Options[v.OptionIndex]))
		case internal.Field:
			s = fmt.Sprintf("%s: %s", m.styles.Options.Render(v.Prompt.Question), m.styles.OptionSelect.Render(v.Prompt.Input))
		}

		if s != "" {
			rows = append(rows, s)
		}
	}

	s := m.styles.OptionSelect.Render("\nConfirm project? (Y/n):")
	rows = append(rows, s)

	return strings.Join(rows, "\n")
}

func (m *model) NextQuestion() {
	if m.qIndex < len(m.questions)-1 {
		m.qIndex++
	} else {
		m.qIndex = 0
	}
}

func (m *model) PrevQuestion() {
	if m.qIndex > 0 {
		m.qIndex--
	} else {
		m.qIndex = len(m.questions) - 1
	}
}

func (m *model) SetAnswerValue() {
	currentAnswer := m.questions[m.qIndex].Prompt.Input

	if currentAnswer != "" {
		m.answerField.SetValue(currentAnswer)
	} else {
		m.answerField.SetValue("")
	}
}

func (m *model) ClearAnswers() {
	q := m.questions

	for i := range q {
		q[i].Prompt.Input = ""
	}
}

func NewDefaultModel(questions []Question) *model {
	mainStyle := style.DefaultStyles()

	answerField := textinput.New()
	answerField.Focus()

	return &model{
		qIndex:      0,
		questions:   questions,
		answerField: answerField,
		styles:      mainStyle,
		keys:        keys,
		help:        help.New(),
	}
}
