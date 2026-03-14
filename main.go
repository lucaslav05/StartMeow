package main

import (
	"fmt"
	// 	tea "charm.land/bubbletea/v2"
	"StartMeow/queue"
	"os/exec"
)

func usage(){
	fmt.Println("Usage: StartMeow <project-name> [--force]")
}

func main() {
	q := queue.InitQueue()

	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	q.Enqueue("4")

	fmt.Println("Initial", q)


	q.Dequeue()
	q.Dequeue()
	fmt.Println("After", q)

	q.Enqueue("100")
	fmt.Println("After", q)



	cmd := exec.Command("git", "branch", "--show-current")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
