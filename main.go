package main

import (
	"fmt"

	// 	tea "charm.land/bubbletea/v2"
	"StartMeow/queue"
)

func usage(){
    fmt.Println("Usage: StartMeow <project-name> [--force]")
}

func main() {
	q := queue.Queue {
		List: make([]any, 0),
	}

	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	q.Enqueue("4")

	fmt.Println("Initial", q)


	q.Dequeue()
	q.Dequeue()
	fmt.Println("After", q)

}
