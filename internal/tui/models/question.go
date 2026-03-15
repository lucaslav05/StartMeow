package internal

import internal "StartMeow/internal"

type QuestionType int

const (
	Select QuestionType = iota
	Prompt
	Verify
)

type Question struct {
	PromptType  internal.PromptType
	Prompt      internal.Prompt
	OptionIndex int
	Answer      string
}

func (q *Question) Next() {
	if q.OptionIndex < len(q.Prompt.questions)-1 {
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
