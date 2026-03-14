package main

import (
	"fmt"
	"os"
	"flag"

// 	tea "charm.land/bubbletea/v2"
	"StartMeow/internal"
)

func usage(){
    fmt.Println("Usage: StartMeow <project-name> [--template type] [--force]")
    return
}

func main() {
// 	fmt.Println("Hello, World!")

    // Define flags
    templatePtr := flag.String("template", "webapp", "template of the project you want to generate")
    forcePtr := flag.Bool("force", false, "if true, overwrite the same named project")

    // Parse flags
    flag.Parse()

    args := flag.Args()

//     if len(args) == 0 && *template == "" {
//         // No args → launch Bubble Tea UI
//         ui.Run()
//         return
//     }

	if len(args) < 1 {
	    usage()
	    os.Exit(1)
	}

    projectName := args[0]

    ctx := generator.Context{
        ProjectName: projectName,
        Template: *templatePtr,
        Force: *forcePtr,
    }

    if err := generator.GenerateProject(ctx); err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

    fmt.Println(projectName, " is ready!")
}
