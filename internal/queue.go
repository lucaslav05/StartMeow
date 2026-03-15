package internal

import (
	"fmt"
)

type Queue struct {
	List []Prompt
}

func InitQueue() Queue {

	return Queue{
		List: make([]Prompt, 0),
	}
}

func (q *Queue) Enqueue(item Prompt) {
	// fmt.Println("enqueue called: " + item.Question)

	q.List = append(q.List, item)

	fmt.Println("WHATS IN")
	fmt.Println(q.List)
}

func (q *Queue) Dequeue() Prompt {
	fmt.Print("dequeue called: ")

	if len(q.List) == 0 {
		fmt.Println("Queue Empty")
		return Prompt{
			Question: "ERROR, QUEUE EMPTY",
		}
	}

	var item = q.List[0]

	q.remove()
	q.resize()

	fmt.Println(item.Question)
	return item
}

func (q *Queue) remove() []Prompt {

	return append(q.List[:0], q.List[1:]...)
}

func (q *Queue) resize() {

	var new_size = len(q.List) - 1
	var new_list = q.List[0:new_size]

	q.List = new_list
}
