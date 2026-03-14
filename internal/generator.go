package internal

import (
    "os"
    "fmt"
    "encoding/json"
    "strings"
    "io/fs"
    "path/filepath"
    "text/template"
)

type Context struct {
    ProjectName string
    Force bool
    Template string
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

func GenerateProject(ctx Context) error {
    templateDir := filepath.Join("templates", ctx.Template)

    // Load manifest.json for that template
    manifestPath := filepath.Join(templateDir, "manifest.json")
    manifest, err := loadManifest(manifestPath)
    if err != nil {
        return err
    }

    // Check if directory exists
    if _, err := os.Stat(ctx.ProjectName); err == nil && !ctx.Force {
        return fmt.Errorf("directory %s already exists (use --force to overwrite)")
    }

    // Debug
    fmt.Println("Using template:", ctx.Template)
    fmt.Println("Template directory:", templateDir)
    fmt.Println("Manifest path:", manifestPath)

    // Render all the template files
    for src, dest := range manifest.Files {
        srcPath := filepath.Join(templateDir, src)
        destPath := filepath.Join(ctx.ProjectName, dest)

        if err := RenderTemplate(srcPath, destPath, ctx); err != nil{
            return err
        }
    }

    return nil
}

func GenerateManifest(selections []string, outputPath string) error {
    manifest := Manifest{Files: make(map[string]string)}

    for _, sel := range selections {
        templateDir := filepath.Join("templates/test", sel)

        err := filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
            if err != nil {
                return err
            }
            if d.IsDir() {
                return nil
            }
            if !strings.HasSuffix(path, ".tmpl") {
                return nil
            }

            rel, err := filepath.Rel(templateDir, path)
            if err != nil {
                return err
            }

            dest := strings.TrimSuffix(rel, ".tmpl")
            manifest.Files[path] = dest
            return nil
        })

        if err != nil {
            return err
        }
    }

    // Convert manifest to JSON
    data, err := json.MarshalIndent(manifest, "", "  ")
    if err != nil {
        return err
    }

    // Write manifest.json to the output path
    return os.WriteFile(outputPath, data, 0644)
}
