package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	style "StartMeow/internal/styles"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	Select = iota
	Prompt
	Verify
)

const (
	White       = "#f8edeb"
	LightOrange = "#f9dcc4"
	Orange      = "#fec89a"
	Black       = "#d6d0c8"
	Black2      = "#dcd2cc"
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

type OptionType struct {
}

type Question struct {
	questionType int
	question     string
	options      []string
	optionIndex  int
	answer       string
}

func NewQuestion(question string, qtype int, options []string) Question {
	return Question{questionType: qtype, question: question, options: options, optionIndex: 0}
}

func DefaultModel(questions []Question) *model {
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
			current.answer = m.answerField.Value()
			m.SetAnswerValue()
			return m, nil
		case "up":
			current.Prev()
		case "down":
			current.Next()
		case "y":
			if current.questionType == Verify {
				log.Println("Verified!")
				return m, tea.Quit
			}
		case "n":
			if current.questionType == Verify {
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

	switch currentQuestion.questionType {
	// current question is prompt
	case Prompt:
		l = lipgloss.JoinVertical(
			lipgloss.Center,
			m.styles.Title.Render(m.questions[m.index].question),
			m.styles.InputField.Render(
				m.answerField.View(),
			),
		)

	// current question is select
	case Select:
		// log.Println("Case: Select")
		l = lipgloss.JoinVertical(
			lipgloss.Left,
			m.styles.Title.Render(m.questions[m.index].question),
			m.ViewSelect(),
		)

	case Verify:
		// current question is verify
		l = lipgloss.JoinVertical(
			lipgloss.Left,
			m.styles.Title.Render(m.questions[m.index].question),
			m.ViewVerify(),
		)
	}

	m.history = append(m.history, l)
	v := tea.NewView(strings.Join(m.history, "\n"))
	v.BackgroundColor = lipgloss.Color(style.BackgroundBlack)
	v.AltScreen = true
	return v
}

func (m *model) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (q *Question) Next() {
	if q.optionIndex < len(q.options)-1 {
		q.optionIndex++
	} else {
		q.optionIndex = 0
	}
}

func (q *Question) Prev() {
	if q.optionIndex > 0 {
		q.optionIndex--
	} else {
		q.optionIndex = len(q.options) - 1
	}
}

func (m *model) SetAnswerValue() {
	currentAnswer := m.questions[m.index].answer

	if currentAnswer != "" {
		m.answerField.SetValue(currentAnswer)
	} else {
		m.answerField.SetValue("")
	}
}

func (m *model) ClearAnswers() {
	q := m.questions

	for i := range q {
		q[i].answer = ""
	}
}

func (m model) ViewSelect() string {
	var rows []string
	currentQuestion := m.questions[m.index]
	options := currentQuestion.options
	optionIndex := currentQuestion.optionIndex

	for i, v := range options {
		if optionIndex == i {
			rows = append(rows, m.styles.OptionSelect.Render(fmt.Sprintf("[x] %s", v)))
		} else {
			rows = append(rows, m.styles.Options.Render(fmt.Sprintf("[ ] %s", v)))
		}
	}

	return strings.Join(rows, "\n")
}

func (m *model) ViewVerify() string {
	var rows []string

	for _, v := range m.questions {
		var s string
		switch v.questionType {
		case Select:
			s = fmt.Sprintf("%s: %s", m.styles.Title.Render(v.question), m.styles.OptionSelect.Render(v.options[v.optionIndex]))
		case Prompt:
			s = fmt.Sprintf("%s: %s", m.styles.Title.Render(v.question), m.styles.OptionSelect.Render(v.answer))
		}

		if s != "" {
			rows = append(rows, s)
		}
	}

	s := "\nConfirm project? (Y/N):"
	rows = append(rows, s)

	return strings.Join(rows, "\n")
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	questions := []Question{
		NewQuestion("Select a language", Select, []string{"C", "C#", "Java", "JavaScript"}),
		NewQuestion("Insert project name", Prompt, []string{}),
		NewQuestion("Verify project structure", Verify, []string{}),
	}

	m := DefaultModel(questions)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}
}
