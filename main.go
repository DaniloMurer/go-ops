package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	goOpsArgs := os.Args

	//version := goOpsArgs[0]
	filePath := goOpsArgs[1]

	buffer, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//command := exec.Command("git", "status")
	//output, _ := command.CombinedOutput()
	lines := strings.Split(string(buffer), "\n")
	fmt.Println(len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println("Hello, World!")
}
