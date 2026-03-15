package models

type Question struct {
	QuestionType int
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
