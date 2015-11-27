package pla

import (
	"errors"
	"fmt"
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

	target.Run(params, stopRunning)
	return nil
}

func FindTargetByTargetName(targetName string, targets []Target) (Target, error) {
	for targetIndex := range targets {
		if targets[targetIndex].Name == targetName {
			return targets[targetIndex], nil
		}
	}
	return Target{}, errors.New("failed to find target")
}
