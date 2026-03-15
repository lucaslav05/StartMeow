package internal

import internal "StartMeow/internal"

type Question struct {
	Prompt      internal.Prompt
	OptionIndex int
}

func (q *Question) Next() {
	if q.OptionIndex < len(q.Prompt.Question)-1 {
		q.OptionIndex++
	} else {
		q.OptionIndex = 0
	}
}

func (q *Question) Prev() {
	if q.OptionIndex > 0 {
		q.OptionIndex--
	} else {
		q.OptionIndex = len(q.Prompt.Options) - 1
	}
}

func NewQuestion(question string, qtype internal.PromptType, options []string) Question {
	p := internal.Prompt{
		Question:   question,
		PromptType: qtype,
		Options:    options,
	}

	return Question{Prompt: p, OptionIndex: 0}
}
