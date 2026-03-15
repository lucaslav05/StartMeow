package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Context struct {
	ProjectName string
	Force       bool
	Template    string
}

type Manifest struct {
	Files map[string]string `json: "files"`
}

// Create the map from the manifest json file
func loadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func RenderTemplate(src, dest string, ctx Context) error {
	// If the template given by user exist, load it, otherwise err and return
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		return err
	}

	//
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	// overwrites file if it exists, f is created file
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, ctx)
}

func GenerateProject(ctx Context, templates Templates) error {
	// Load the generated Manifest file
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	manifest, err := loadManifest(pwd + "/manifest.json")
	if err != nil {
		return err
	}

	// Raise error if the project directory already exists and --force flag is not set
	if _, err := os.Stat(ctx.ProjectName); err == nil && !ctx.Force {
		return fmt.Errorf("directory %s already exists (use --force to overwrite)", ctx.ProjectName)
	}

	for src, dest := range manifest.Files {

		srcPath := src
		destPath := filepath.Join(pwd, dest)

		// Use a buffer as the io.Writer so output is actually captured
		var buf bytes.Buffer
		if err := templates.T.ExecuteTemplate(&buf, srcPath, ctx); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", srcPath, err)
		}

		// Write the rendered buffer to the destination file
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create directories for %s: %w", destPath, err)
		}
		if err := os.WriteFile(destPath, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", destPath, err)
		}
	}

	return nil
}

func GenerateManifest(fileMap map[string]string, outputPath string) error {
	manifest := Manifest{Files: make(map[string]string)}

	for src, dest := range fileMap {
		manifest.Files[src] = dest
	}

	// Convert manifest to JSON
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	// Write manifest.json to the output path
	fmt.Printf("Writing manifest to %s \n", outputPath)
	return os.WriteFile(outputPath, data, 0644)
}
