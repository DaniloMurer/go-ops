package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	goOpsArgs := os.Args
	filePath := goOpsArgs[1]
	oldVersion := goOpsArgs[2]
	newVersion := goOpsArgs[3]

	buffer, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// split the file content to lines, so we can only modify the image line
	lines := strings.Split(string(buffer), "\n")
	for index, line := range lines {
		// if the current line contains the old image version, we want to replace it
		if strings.Contains(line, oldVersion) {
			// replace old version of the file with new
			lines[index] = strings.Replace(line, oldVersion, newVersion, 1)
		}
	}
	// write content with updated image version to file
	os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0660)
	// todo: make the git commit with commit message containing "bumped to version {newVersion}"
	// i think we can use the exec.Command() library here. gotta keep it simple
}
