package main

import (
	"fmt"
	"github.com/rtuin/go-pla"
	"os"
)

func main() {
	fmt.Println("GoPla master by Richard Tuin - Coder's simplest workflow automation tool.\n")

	targets, err := pla.LoadTargets("Plafile.yml")
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	calledTarget := "all"
	if len(args) > 0 {
		calledTarget = args[0]
	}

	var params []string
	if len(args) > 1 {
		params = args[1:]
	}

	pla.RunTargetByName(calledTarget, targets, false, params)
}
