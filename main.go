package main

import (
	"StartMeow/internal"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"
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

	// f, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// questions := []models.Question{
	// 	models.NewQuestion("Select a language", models.Select, []string{"C", "C#", "Java", "JavaScript"}),
	// 	models.NewQuestion("Insert project name", models.Prompt, []string{}),
	// 	models.NewQuestion("Verify project structure", models.Verify, []string{}),
	// }

	// m := models.NewDefaultModel(questions)

	// p := tea.NewProgram(m)
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Error occurred: %v", err)
	// 	os.Exit(1)
	// }

	testProj := internal.Project{
		ProjType:   internal.WebApp,
		FrontFrame: internal.React,
		BackFrame:  internal.NextJS,
		FrontLang:  internal.Javascript,
		BackLang:   internal.Typescript,
		Ui:         internal.Download,
		Database:   internal.MongoDB,
		ProjName:   "test-project",
	}

	internal.BuildProject(&testProj, *tmpl)

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
