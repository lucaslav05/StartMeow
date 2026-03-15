package internal

import "charm.land/lipgloss/v2"

const (
	White           = "#f8edeb"
	LightOrange     = "#f9dcc4"
	Orange          = "#fec89a"
	Black           = "#d6d0c8"
	BackgroundBlack = "#3A2E2E"
	Black2          = "#dcd2cc"
)

type MainStyle struct {
	Title        lipgloss.Style
	Options      lipgloss.Style
	OptionSelect lipgloss.Style
	InputField   lipgloss.Style
	TabStyle     lipgloss.Style
}

func DefaultStyles() *MainStyle {
	m := new(MainStyle)

	m.Title = TitleStyle()
	m.Options = OptionsStyle()
	m.OptionSelect = OptionSelectStyle()
	m.InputField = InputFieldStyle()
	m.TabStyle = TabStyle()

	return m
}

func InputFieldStyle() lipgloss.Style {
	s := lipgloss.NewStyle().
		BorderForeground(
			lipgloss.Color(Orange),
		).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(40)

	return s
}

func TabStyle() lipgloss.Style {
	s := lipgloss.NewStyle().
		BorderForeground(
			lipgloss.Color(Black),
		).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1, 3)

	return s
}

func TitleStyle() lipgloss.Style {
	s := lipgloss.NewStyle().
		Foreground(
			lipgloss.Color(Orange),
		)

	return s
}

func OptionsStyle() lipgloss.Style {
	s := lipgloss.NewStyle().
		Foreground(
			lipgloss.Color(White),
		)

	return s
}

func OptionSelectStyle() lipgloss.Style {
	s := lipgloss.NewStyle().
		Foreground(
			lipgloss.Color(LightOrange),
		)

	return s
}
