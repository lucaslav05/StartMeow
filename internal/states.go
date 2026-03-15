package internal

type state struct {
	answers map[string]state
	prompts []Prompt
}

func UpdateProjectState(newState state) {
	projectState = newState
}

// Comparison state
var projectState = state{
	prompts: []Prompt{},
	answers: make(map[string]state),
}

var WebappState = state{
	prompts: []Prompt{AppType},
	answers: map[string]state{
		"yes": FrontendState,
	},
}

// 2 options state
var FrontendState = state{
	prompts: []Prompt{FrontendFrameworkType},
	answers: map[string]state{
		"yes": BackendFrameworkReactState,
		"no":  BackendFrameworkState,
	},
}

var BackendFrameworkReactState = state{
	prompts: []Prompt{BackendFrameworkReact},
	answers: map[string]state{
		"React":   StarterUIState,
		"Nextjs":  StarterUIState,
		"Express": StarterUIState,
	},
}

var StarterUIState = state{
	prompts: []Prompt{StartingUI},
	answers: map[string]state{
		"Store":    WhichDBState,
		"Download": WhichDBState,
		"Blog":     WhichDBState,
		"Empty":    WhichDBState,
	},
}

// start of 'no' frontend framework path
var BackendFrameworkState = state{
	prompts: []Prompt{WhichLanguage, StartingUI},
	answers: map[string]state{
		"Node":    StarterUIState,
		"Express": StarterUIState,
		"None":    WhichLanguageState,
	},
}

var WhichLanguageState = state{
	prompts: []Prompt{WhichLanguage},
	answers: map[string]state{
		"C":          StarterUIState,
		"Javascript": StarterUIState,
		"Typescript": StarterUIState,
		"Go":         StarterUIState,
		"C++":        StarterUIState,
		"Java":       StarterUIState,
		"Swift":      StarterUIState,
		"Kotlin":     StarterUIState,
		"CSharp":     StarterUIState,
		"Jsp":        StarterUIState,
	},
}

var WhichDBState = state{
	prompts: []Prompt{WhichDB},
	answers: map[string]state{
		"MongoDB": emptyState,
		"SQLite":  emptyState,
		"None":    emptyState,
	},
}

var emptyState = state{}
