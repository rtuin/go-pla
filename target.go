package pla

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Runnable interface {
	Run(params []string, stopRunning bool) bool
}

type Target struct {
	Name       string
	Commands   []Runnable
	Parameters []string
}

type Command struct {
	Command string
}

func (t Target) Run(params []string, stopRunning bool) bool {
	for i := range t.Commands {
		stopRunning = t.Commands[i].Run(params, stopRunning)
	}
	return stopRunning
}

func (c Command) Run(params []string, stopRunning bool) bool {
	if stopRunning {
		fmt.Printf("\x1b[37;2m    . %v\x1b[0m\n", c.Command)
		return true
	}

	commandString := c.Command
	// if len(params) > 0 {
	// 	for index := range params {
	// 		commandString = strings.Replace(commandString, fmt.Sprintf("%%%v%%", target.Parameters[index]), params[index], -1)
	// 	}
	// }

	fmt.Printf("    ⌛ %v", c.Command)

	cmd := exec.Command("sh", "-c", commandString)
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\033[2K\r\x1b[31;1m    ✘ %v\x1b[0m\n", c.Command)
		strErrLines := strings.Split(stdErr.String(), "\n")
		if len(stdErr.String()) == 0 {
			strErrLines = []string{"[no output]"}
		}

		for lineIndex := range strErrLines {
			fmt.Printf("\x1b[31;2m        %s\x1b[0m\n", strErrLines[lineIndex])
		}
		return true
	}
	fmt.Printf("\033[2K\r\x1b[32m    ✔ %v\x1b[0m\n", c.Command)

	return false
}
