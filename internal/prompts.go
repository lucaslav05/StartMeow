package internal

//go:generate stringer -type=UserInterface

import (
	"fmt"
	"log"
)

func InitPrompts() Queue {
	q := InitQueue()
	q.Enqueue(FrontendFrameworkType)

	UpdateProjectState(FrontendState)
	// // fmt.Println("Initiating the prompts proj stat: " + projectState.prompts[0].Question)
	// fmt.Println("INIT OF QEUUE")
	// fmt.Println(q)
	return q
}

func StateRouter(promptqueue *Queue, answer Prompt, pStruct *Project) Prompt {
	// first split decision (frontend into react or null)
	log.Println(projectState.prompts[0].Question)
	// decided to compare questions because they are unique (pray to God)
	if answer.Input == "yes" && answer.Question == FrontendFrameworkType.Question {
		// fmt.Println("Case 1")
		UpdateProjectState(WhichDBState)
		promptqueue.Enqueue(BackendFrameworkReact)
		promptqueue.Enqueue(StartingUI)
		promptqueue.Enqueue(WhichDB)
		// fmt.Println("End of case 1")
		//end of config section
	} else if answer.Input == "no" && answer.Question == FrontendFrameworkType.Question { // the no react route
		// fmt.Println("Case 2")
		promptqueue.Enqueue(BackendFramework)
	} else if answer.Input == "None" && answer.Question == BackendFramework.Question { // no react no backend framework
		// log.Println("Case 3")
		promptqueue.Enqueue(WhichLanguage)
		promptqueue.Enqueue(StartingUI)
		promptqueue.Enqueue(WhichDB)
		//end of config section
	} else if answer.Question == BackendFramework.Question {

		promptqueue.Enqueue(StartingUI)
		promptqueue.Enqueue(WhichDB)
	} else {
		promptqueue.Enqueue(EndingState)
	}

	UpdateProjectStruct(answer, pStruct)
	log.Print(*pStruct)

	p := promptqueue.Dequeue()
	// fmt.Println(promptqueue.List)
	return p
}

func UpdateProjectStruct(answer Prompt, pStruct *Project) {
	switch answer.Question {
	case FrontendFrameworkType.Question:
		t, _ := ResolveFramework(answer.Input)
		pStruct.FrontFrame = t
	case BackendFramework.Question:
		t, _ := ResolveFramework(answer.Input)
		pStruct.FrontFrame = t
	case WhichLanguage.Question:
		t, _ := ResolveLanguage(answer.Input)
		pStruct.BackLang = t
	case StartingUI.Question:
		t, _ := ResolveUserInterface(answer.Input)
		pStruct.Ui = t
	case WhichDB.Question:
		t, _ := ResolveDatabase(answer.Input)
		pStruct.Database = t
	}
}

// first question
var AppType = Prompt{
	Question:   "Pick an Application",
	PromptType: Select,
	Options: []string{
		"Web app",
		"Client/Server",
		"Mobile",
		"Terminal",
	},
}

// second question
var FrontendFrameworkType = Prompt{
	Question:   "Do you want a React frontend framework?",
	PromptType: Select,
	Options: []string{
		"yes",
		"no",
	},
}

// if it's yes queue these prompts
var BackendFrameworkReact = Prompt{
	Question:   "Choose a backend React framework",
	PromptType: Select,
	Options: []string{
		"Nextjs",
		"ReactRouter",
		"Express",
	},
}

// if no, queue the non framework prompts
var BackendFramework = Prompt{
	Question:   "Choose a backend framework",
	PromptType: Select,
	Options: []string{
		"Express",
		"Node",
		"None",
	},
}

// this queues if the no react backend framework option was "none"
var WhichLanguage = Prompt{
	Question:   "Choose a language for your backend",
	PromptType: Select,
	Options: []string{
		"C",
		"Javascript",
		"Typescript",
		"Go",
		"C++",
		"Java",
		"Swift",
		"Kotlin",
		"CSharp",
		"Jsx",
	},
}

// WhichLanguage, BackendFramework, BackendFrameworkReact all queue this prompt after
var StartingUI = Prompt{
	Question:   "Pick a starting UI",
	PromptType: Select,
	Options: []string{
		"Store",
		"Download",
		"Blog",
		"Empty",
	},
}

// this comes after starting UI prompt
var WhichDB = Prompt{
	Question:   "end of cycle into new section",
	PromptType: Select,
	Options: []string{
		"MongoDB",
		"SQLite",
		"None",
	},
}

var EndingState = Prompt{
	Question:   "Confirm project configuration",
	PromptType: Info,
	Options: []string{
		"Y",
		"n",
	},
}

func ResolveProjectType(answer string) (ProjectType, error) {
	switch answer {
	case "Web app":
		return WebApp, nil
	case "Client/Server":
		return ClientServer, nil
	case "Mobile":
		return Mobile, nil
	case "Terminal":
		return Terminal, nil
	default:
		return 0, fmt.Errorf("unknown ProjectType answer: %q", answer)
	}
}

func ResolveFramework(answer string) (Framework, error) {
	switch answer {
	case "yes":
		return React, nil
	case "Nextjs":
		return NextJS, nil
	case "ReactRouter":
		return ReactRouter, nil
	case "Express":
		return ExpressJs, nil
	case "Node":
		return NodeJS, nil
	default:
		return 0, fmt.Errorf("unknown Framework answer: %q", answer)
	}
}

func ResolveLanguage(answer string) (Language, error) {
	switch answer {
	case "C":
		return C, nil
	case "Javascript":
		return Javascript, nil
	case "Typescript":
		return Typescript, nil
	case "Go":
		return Go, nil
	case "C++":
		return CPlusPlus, nil
	case "Java":
		return Java, nil
	case "Swift":
		return Swift, nil
	case "Kotlin":
		return Kotlin, nil
	case "CSharp":
		return CSharp, nil
	case "Jsx":
		return Jsx, nil
	default:
		return 0, fmt.Errorf("unknown Language answer: %q", answer)
	}
}

func ResolveUserInterface(answer string) (UserInterface, error) {
	switch answer {
	case "Store":
		return Store, nil
	case "Download":
		return Download, nil
	case "Blog":
		return Blog, nil
	case "Empty":
		return Landing, nil
	default:
		return 0, fmt.Errorf("unknown UserInterface answer: %q", answer)
	}
}

func ResolveDatabase(answer string) (Database, error) {
	switch answer {
	case "MongoDB":
		return MongoDB, nil
	case "SQLite":
		return SQLite, nil
	case "None":
		return None, nil
	default:
		return 0, fmt.Errorf("unknown Database answer: %q", answer)
	}
}
