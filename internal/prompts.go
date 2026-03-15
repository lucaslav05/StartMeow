package internal

func InitPrompts() Queue {
	q := InitPrompts()
	q.Enqueue(AppType)

	return q
}

func StateRouter(promptqueue *Queue, answer Prompt) Prompt {

	return promptqueue.Dequeue()
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
		"c",
		"Javascript",
		"Typescript",
		"Go",
		"c++",
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
	Question:   "Pick a starting UI",
	PromptType: Select,
	Options: []string{
		"MongoDB",
		"Firebase",
		"SQLite",
	},
}
