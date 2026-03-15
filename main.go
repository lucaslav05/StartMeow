package main

import (
	internal "StartMeow/internal"
	models "StartMeow/internal/tui/models"
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	questions := []models.Question{
		models.NewQuestion("Select a language", internal.Select, []string{"C", "C#", "Java", "JavaScript"}),
		models.NewQuestion("Insert project name", internal.Field, []string{}),
		models.NewQuestion("Verify project structure", internal.Info, []string{}),
	}

	m := models.NewDefaultModel(questions)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occurred: %v", err)
		os.Exit(1)
	}

	// testProj := internal.Project{
	// 	ProjType:   internal.WebApp,
	// 	FrontFrame: internal.React,
	// 	BackFrame:  internal.NextJS,
	// 	FrontLang:  internal.Javascript,
	// 	BackLang:   internal.Typescript,
	// 	Ui:         internal.Download,
	// 	Database:   internal.MongoDB,
	// 	ProjName:   "test-project",
	// }

	// internal.BuildProject(&testProj)

	// 	type Project struct {
	// 	ProjType   ProjectType
	// 	FrontFrame Framework
	// 	BackFrame  Framework
	// 	FrontLang  Language
	// 	BackLang   Language
	// 	Ui         UserInterface
	// 	Database   Database
	// 	ProjName   string
	// 	FilePath   string
	// }

}
