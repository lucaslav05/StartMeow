package internal

//go:generate stringer -type=UserInterface

type UserInterface int

const (
	Store UserInterface = iota
	Download
	Blog
	Empty
)

type ProjectType int
type Framework int
type Language int
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
	ProjType   ProjectType
	FrontFrame Framework
	BackFrame  Framework
	FrontLang  Language
	BackLang   Language
	Ui         UserInterface
	Database   Database
	ProjName   string
	FilePath   string
}

type Prompt struct {
	Question   string
	PromptType PromptType
	Options    []string
	Input      string
}

var templatePaths = map[string]string{"expressServer": "src/server.js"}
