package internal

type ProjectType int
type Framework int
type Language int
type UserInterface int
type PromptType int
type Database int

const (
	WebApp ProjectType = iota
	ClientServer
	Mobile
	Terminal
)

const (
	React Framework = iota
	ExpressJs
	NodeJS
	NextJS
	ReactRouter
)

const (
	C Language = iota
	Javascript
	Typescript
	Go
	CPlusPlus
	Java
	Swift
	Kotlin
	CSharp
	Jsx
)

const (
	Store UserInterface = iota
	Download
	Blog
	Empty
)

const (
	Select PromptType = iota
	Field
	Info
)

const (
	MongoDB Database = iota
	Firebase
	SQLite
)

type Project struct {
	projType   ProjectType
	frontFrame Framework
	backFrame  Framework
	frontLang  Language
	backLang   Language
	ui         UserInterface
	database   Database
	projName   string
	filePath   string
}

type Prompt struct {
	title      string
	promptType PromptType
	questions  []string
}

var templatePaths = map[string]string{"expressServer": "src/server.js"}
