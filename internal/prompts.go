package internal

import "log"

func InitPrompts() Queue {
	q := InitQueue()
	q.Enqueue(FrontendFrameworkType)

	UpdateProjectState(FrontendState)
	// // fmt.Println("Initiating the prompts proj stat: " + projectState.prompts[0].Question)
	// fmt.Println("INIT OF QEUUE")
	// fmt.Println(q)
	return q
}

func StateRouter(promptqueue *Queue, answer Prompt) Prompt {
	// first split decision (frontend into react or null)
	log.Println(projectState.prompts[0].Question)
	// decided to compare questions because they are unique (pray to God)
	if answer.Input == "yes" && answer.Question == FrontendFrameworkType.Question {
		// fmt.Println("Case 1")
		promptqueue.Enqueue(BackendFrameworkReact)
		promptqueue.Enqueue(StartingUI)
		promptqueue.Enqueue(WhichDB)
		// UpdateProjectState(WhichDBState)
		// fmt.Println("End of case 1")
		//end of config section
	} else if answer.Input == "no" && answer.Question == FrontendFrameworkType.Question { // the no react route
		// fmt.Println("Case 2")
		promptqueue.Enqueue(BackendFramework)
	} else if answer.Input == "None" && answer.Question == BackendFramework.Question { // no react no backend framework
		log.Println("Case 3")
		promptqueue.Enqueue(WhichLanguage)
		promptqueue.Enqueue(StartingUI)
		promptqueue.Enqueue(WhichDB)
		//end of config section
	} else {
		promptqueue.Enqueue(EndingState)
	}

	p := promptqueue.Dequeue()
	// fmt.Println(promptqueue.List)
	return p
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
