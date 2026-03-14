package main

import (
	"fmt"
	// "os/exec"
	"StartMeow/internal"
)

func main() {
	// fmt.Println("Hello, World!")

	// cmd := exec.Command("git", "branch", "--show-current")
	// stdout, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(string(stdout))

	selections := []string{"webapp", "express", "homepage", "storepage", "aboutus"}
	err := internal.GenerateManifest(selections, "manifest.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Manifest created at manifest.json")

}
