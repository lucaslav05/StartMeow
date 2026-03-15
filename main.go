package main

import (
	"fmt"
	"log"
	"os"

	"StartMeow/internal"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// init prompt queue
	q := internal.InitPrompts()

	questions := []models.Question{
		models.NewQuestion("Select the language of the project", internal.Select, []string{"C", "C#", "Java", "JavaScript"}),
		models.NewQuestion("Insert project name", internal.Field, []string{}),
		models.NewQuestion("Verify project structure", internal.Info, []string{}),
	}

	m := models.NewDefaultModel(questions)

	// Evins intervention
	// m := models.NewDefaultModelQ(q)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}
}
