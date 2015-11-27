package pla

import (
	"errors"
	"fmt"
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
	"strings"
)

func LoadTargets(filename string) ([]Target, error) {
	data, error := ioutil.ReadFile(filename)
	if error != nil {
		panic(error)
	}

	yaml, error := simpleyaml.NewYaml(data)
	if error != nil {
		panic(error)
	}

	m, err := yaml.Map()
	if err != nil {
		panic(err)
	}

	var targets []Target

	for key, value := range m {
		value := value.([]interface{})
		commands := make([]Runnable, len(value))
		targetName := simplifyTargetName(key.(string))
		targetParams, targetParamErr := findParametersInTargetName(key.(string))
		if targetParamErr != nil {
			panic(targetParamErr)
		}

		for i := range value {
			commandString := value[i].(string)
			if commandString[0] == "="[0] {
				panic("Subtargets not implemented yet!")
			}
			commands[i] = Command{commandString}
			// commands[i] = Command{value[i].(string)}
		}
		targets = append(targets, Target{targetName, commands, targetParams})
	}

	// for targetIndex := range targets {
	// 	for commandIndex := range targets[targetIndex].Commands {
	// 		currentCommand := targets[targetIndex].Commands[commandIndex].(string)
	// 		if currentCommand[0] != "="[0] {
	// 			continue
	// 		}

	// 		subTargetName := currentCommand[1:]
	// 		subTarget, err := FindTargetByTargetName(subTargetName, targets)
	// 		if err != nil {
	// 			fmt.Printf("Error: Using non-existent target \"%v\" as subtarget.\n", subTargetName)
	// 			fmt.Println("Valid targets are:")
	// 			for tIndex := range targets {
	// 				fmt.Println("  -", targets[tIndex].Name)
	// 			}
	// 			return nil, err
	// 		}

	// 		targets[targetIndex].Commands[commandIndex] = subTarget
	// 	}
	// }

	return targets, nil
}

func findParametersInTargetName(rawName string) ([]string, error) {
	paramStartIndex := strings.Index(rawName, "[")
	paramEndIndex := strings.Index(rawName, "]")
	if paramStartIndex == -1 {
		return make([]string, 0), nil
	}

	if paramEndIndex == -1 {
		return nil, errors.New(fmt.Sprintf("Incorrect syntax for target %v", rawName))
	}

	return strings.Split(rawName[paramStartIndex+1:paramEndIndex], ","), nil
}

func simplifyTargetName(rawName string) string {
	if rawName[0] == "="[0] {
		rawName = rawName[1:]
	}

	paramStartIndex := strings.Index(rawName, "[")
	if paramStartIndex == -1 {
		return rawName
	}
	return rawName[:paramStartIndex]
}
