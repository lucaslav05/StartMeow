package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello, World!")

	cmd := exec.Command("git", "branch", "--show-current")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))

}
