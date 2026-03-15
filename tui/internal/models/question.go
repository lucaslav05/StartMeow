package models

type QuestionType int

const (
	Select QuestionType = iota
	Prompt
	Verify
)

type Question struct {
	QuestionType QuestionType
	Question     string
	Options      []string
	OptionIndex  int
	Answer       string
}

func (q *Question) Next() {
	if q.OptionIndex < len(q.Options)-1 {
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

func NewQuestion(question string, qtype QuestionType, options []string) Question {
	return Question{QuestionType: qtype, Question: question, Options: options, OptionIndex: 0}
}
