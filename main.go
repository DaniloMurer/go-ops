package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {

	goOpsArgs := os.Args
	filePath := goOpsArgs[1]
	oldVersion := goOpsArgs[2]
	newVersion := goOpsArgs[3]

	buffer, err := os.ReadFile(filePath)
	fmt.Println(path.Dir(filePath))
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
	if err := exec.Command("git", "add", filePath).Run(); err != nil {
		fmt.Print(err)
	}
	if err := exec.Command("git", "-C", path.Dir(filePath), "commit", "-m", "feat: bumped to version: "+newVersion).Run(); err != nil {
		fmt.Print(err)
	}
	//fixme: handle git credentials
	if err := exec.Command("git", "-C", path.Dir(filePath), "push").Run(); err != nil {
		fmt.Print(err)
	}
}
