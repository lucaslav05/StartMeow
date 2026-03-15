package main

import (
	"StartMeow/internal"
	models "StartMeow/internal/tui/models"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

//go:embed all:templates
var templateFS embed.FS

func main() {
	walk := func(fsys embed.FS) *internal.Templates {
		t := template.New("")

		err := fs.WalkDir(fsys, "templates", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".tmpl" {
				contents, err := fsys.ReadFile(path)
				if err != nil {
					return err
				}
				template.Must(t.New(path).Parse(string(contents)))
			}
			return nil
		})
		if err != nil {
			log.Fatalf("failed to parse templates: %v", err)
		}

		return &internal.Templates{T: t}
	}

	tmpl := walk(templateFS)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	q := internal.InitPrompts()
	initialP := q.Dequeue()

	state := false
	pStruct := internal.Project{}

	questions := []models.Question{models.NewQuestion(initialP.Question, initialP.PromptType, initialP.Options)}

	m := models.NewDefaultModel(questions, q, &pStruct, &state)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(pStruct)
	if state {
		internal.BuildProject(&pStruct, *tmpl)
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
