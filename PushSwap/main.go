package main

import (
	PSFunc "PS/PushSwap/Functions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]
	elements := strings.Split(input, " ")
	var primaryStack []int

	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("ERROR: Invalid input, not an integer")
			return
		}
		primaryStack = append(primaryStack, num)
	}

	var secondaryStack []int

	if PSFunc.ContainsDuplicates(primaryStack) {
		fmt.Println("ERROR: Duplicate numbers found")
		return
	} else if PSFunc.IsSorted(primaryStack) {
		fmt.Println("ERROR: Stack is already sorted")
		return
	}

	operationsLog := PSFunc.SortStacks(&primaryStack, &secondaryStack)

	for _, operation := range operationsLog {
		fmt.Println(operation)
	}
}
