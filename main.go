package main

import (
	"fmt"
	"os"

// 	tea "charm.land/bubbletea/v2"
	"StartMeow/generator"
)

func usage(){
    fmt.Println("Usage: StartMeow <project-name> [--force]")
    return
}

func main() {
// 	fmt.Println("Hello, World!")
	if len(os.Args) < 2 {
	    usage()
	    return
	}

    projectName := os.Args[1]
    force := false

    // Check for --force flag
    if len(os.Args) > 2 && os.Args[2] == "--force" {
        force = true
    }

    ctx := generator.Context{
        ProjectName: projectName,
        Force: force,
    }

    fmt.Println("Generating barebones webapp: ", projectName)

    if err := generator.GenerateWebapp(ctx); err != nil {
        fmt.Println("Error: ", err)
        return
    }

    fmt.Println(projectName, " is ready!")
}
