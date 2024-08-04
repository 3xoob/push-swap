package main

import (
	CFunc "PS/Check/Functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	stackA, err := CFunc.ParseInputValues(strings.Split(os.Args[1], " "))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	stackB := &CFunc.IntegerStack{}

	var instructions []string
	for {
		var instruction string
		if _, err := fmt.Scanln(&instruction); err != nil {
			break
		}
		instructions = append(instructions, instruction)
	}

	for _, instruction := range instructions {
		if err := CFunc.PerformInstruction(instruction, stackA, stackB); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	if stackA.IsSortedAscending() && len(stackB.Values) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
