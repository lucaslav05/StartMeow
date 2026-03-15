package main

import (
	"fmt"
	"log"
	"os"

	models "StartMeow/internal/models"

	tea "charm.land/bubbletea/v2"
)

func NewQuestion(question string, qtype int, options []string) models.Question {
	return models.Question{QuestionType: qtype, Question: question, Options: options, OptionIndex: 0}
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	questions := []models.Question{
		NewQuestion("Select a language", models.Select, []string{"C", "C#", "Java", "JavaScript"}),
		NewQuestion("Insert project name", models.Prompt, []string{}),
		NewQuestion("Verify project structure", models.Verify, []string{}),
	}

	m := models.NewDefaultModel(questions)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}
}
