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

	q := internal.InitPrompts()
	initialP := q.Dequeue()

	pStruct := internal.Project{}

	questions := []models.Question{models.NewQuestion(initialP.Question, initialP.PromptType, initialP.Options)}

	m := models.NewDefaultModel(questions, q, &pStruct)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(pStruct)
	internal.BuildProject(&pStruct)

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
