package internal

type ProjectType int
type Framework int
type Language int
type FrontLang int
type MobileLang int
type UserInterface int
type PromptType int

const (
	WebApp ProjectType = iota
	ClientServer
	Mobile
	Terminal
)

const (
	React Framework = iota
	Express
	None
)

const (
	C Language = iota
	Javascript
	Typescript
	Go
	CPlusPlus
	Java
)

const (
	Js FrontLang = iota
	Ts
)

const (
	Swift MobileLang = iota
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

type Project struct {
	projType   ProjectType
	framework  Framework
	frontLang  FrontLang
	language   Language
	mobileLang MobileLang
	ui         UserInterface
	projName   string
	filePath   string
}

type Prompt struct {
	title      string
	promptType PromptType
	questions  []string
}
