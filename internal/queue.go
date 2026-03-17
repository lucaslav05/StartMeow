package internal

type Queue struct {
	List []Prompt
}

func InitQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(item Prompt) {
	q.List = append(q.List, item)
}

func (q *Queue) Dequeue() Prompt {
	if len(q.List) == 0 {
		return Prompt{
			Question: "ERROR, QUEUE EMPTY",
		}
	}

	item := q.List[0]
	q.List = q.List[1:]

	return item
}
