package generator

import (
    "os"
    "fmt"
    "path/filepath"
    "text/template"
)

type Context struct {
    ProjectName string
    Force bool
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

func GenerateWebapp(ctx Context) error {
    // Check if directory exists
    if _, err := os.Stat(ctx.ProjectName); err == nil && ctx.Force {
        return fmt.Errorf("directory %s already exists (use --force to overwrite)")
    }

    // The location of the templates for a webapp
    base := "templates/webapp"

    // Map all the template files to the file they will create
    files := map[string]string{
        "public/index.html.tmpl":   "public/index.html",
        "src/index.js.tmpl":        "src/index.js",
        "gitignore.tmpl":           ".gitignore",
        "package.json.tmpl":        "package.json",
        "README.md.tmpl":           "README.md",
    }

    for src, dest := range files {
        srcPath := filepath.Join(base, src)
        destPath := filepath.Join(ctx.ProjectName, dest)

        if err := RenderTemplate(srcPath, destPath, ctx); err != nil{
            return err
        }
    }

    return nil
}