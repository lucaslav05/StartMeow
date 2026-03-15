<<<<<<< HEAD
package internal

import internal "StartMeow/internal"
=======
package models
>>>>>>> dev

type QuestionType int

const (
	Select QuestionType = iota
	Prompt
	Verify
)

type Question struct {
<<<<<<< HEAD
	PromptType  internal.PromptType
	Prompt      internal.Prompt
	OptionIndex int
	Answer      string
}

func (q *Question) Next() {
	if q.OptionIndex < len(q.Prompt.questions)-1 {
=======
	QuestionType QuestionType
	Question     string
	Options      []string
	OptionIndex  int
	Answer       string
}

func (q *Question) Next() {
	if q.OptionIndex < len(q.Options)-1 {
>>>>>>> dev
		q.OptionIndex++
	} else {
		q.OptionIndex = 0
	}
}

func (q *Question) Prev() {
	if q.OptionIndex > 0 {
		q.OptionIndex--
	} else {
		q.OptionIndex = len(q.Options) - 1
	}
}
<<<<<<< HEAD
=======

func NewQuestion(question string, qtype QuestionType, options []string) Question {
	return Question{QuestionType: qtype, Question: question, Options: options, OptionIndex: 0}
}
>>>>>>> dev
