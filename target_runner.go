package pla

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func RunTargetByName(targetName string, targets []Target, stopRunning bool, params []string) error {

	target, error := FindTargetByTargetName(targetName, targets)
	if error != nil {
		err := fmt.Sprintf("Error: Invalid value: Target \"%v\" not present in Plafile.yml.\n", targetName)
		fmt.Printf(err)
		fmt.Println("Valid targets are:")
		for tIndex := range targets {
			fmt.Println("  -", targets[tIndex].Name)
		}
		return errors.New(err)
	}

	if len(params) < len(target.Parameters) {
		missingParameters := target.Parameters[len(params):]
		err := fmt.Sprintf("\x1b[31;2mCannot run \"%v\": Parameter \x1b[31;1m\x1b[31;4m%v\x1b[0m\x1b[31;2m not provided.\x1b[0m\n", target.Name, strings.Join(missingParameters, ", "))
		fmt.Printf(err)
		return errors.New(err)
	}

	fmt.Printf("Running target \"%v\":\n", targetName)

	runTargetCommands(target, stopRunning, params)
	return nil
}

func runTargetCommands(target Target, stopRunning bool, params []string) bool {
commandLoop:
	for commandIndex := range target.Commands {
		switch commandType := target.Commands[commandIndex].(type) {
		case Target:
			stopRunning = runTargetCommands(commandType, stopRunning, params)
			continue commandLoop
		}

		rawCommandString := target.Commands[commandIndex].(string)
		if stopRunning {
			fmt.Printf("\x1b[37;2m    . %v\x1b[0m\n", target.Commands[commandIndex])
			continue
		}

		commandString := rawCommandString
		if len(params) > 0 {
			for index := range params {
				commandString = strings.Replace(commandString, fmt.Sprintf("%%%v%%", target.Parameters[index]), params[index], -1)
			}
			// rawCommandString = commandString
		}

		fmt.Printf("    ⌛ %v", rawCommandString)
		fmt.Println("cmd:", commandString)

		cmd := exec.Command("sh", "-c", commandString)
		var stdErr bytes.Buffer
		cmd.Stderr = &stdErr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\033[2K\r\x1b[31;1m    ✘ %v\x1b[0m\n", rawCommandString)
			strErrLines := strings.Split(stdErr.String(), "\n")
			for lineIndex := range strErrLines {
				fmt.Printf("\x1b[31;2m        %s\x1b[0m\n", strErrLines[lineIndex])
			}
			stopRunning = true
			continue
		}
		fmt.Printf("\033[2K\r\x1b[32m    ✔ %v\x1b[0m\n", rawCommandString)
	}
	return stopRunning
}

func FindTargetByTargetName(targetName string, targets []Target) (Target, error) {
	for targetIndex := range targets {
		if targets[targetIndex].Name == targetName {
			return targets[targetIndex], nil
			break
		}
	}
	return Target{}, errors.New("failed to find target")
}
