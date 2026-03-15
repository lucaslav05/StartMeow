package internal

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
	Question:   "Frontend Framework Type",
	PromptType: Select,
	Options:    []string{"Do you want a React frontend framework?"},
}

// if it's yes queue these prompts
var BackendFramework = Prompt{
	Question:   "Backend Framework Type",
	PromptType: Select,
	Options: []string{
		"Choose a backend React framework",
		"Nextjs",
		"ReactRouter",
		"Express",
	},
}
