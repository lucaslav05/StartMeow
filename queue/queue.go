package queue

import (
	"fmt"
)

type Queue struct {
	List []any;
}

func InitQueue() Queue {

	return Queue {
		List: make([]any, 0),
	}
}

func (q *Queue) Enqueue(item any) {
	fmt.Println("enqueue called")

	q.List = append(q.List, item)
}


func (q *Queue) Dequeue() any {
	fmt.Println("dequeue called")

	if (len(q.List) == 0) {
		fmt.Println("Queue Empty")
		return nil
	}
	
	var item = q.List[0]

	q.remove()
	q.resize()

	return item
}

func (q *Queue)remove() []any {

	return append(q.List[:0], q.List[1:]...)
}

func (q *Queue) resize() {

	var new_size = len(q.List) - 1
	var new_list = q.List[0:new_size]

	q.List = new_list
}
