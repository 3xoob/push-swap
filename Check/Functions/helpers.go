package CFunc

import (
	"fmt"
	"strconv"
)

type IntegerStack struct {
	Values []int
}

func (stack *IntegerStack) IsSortedAscending() bool {
	for i := 0; i < len(stack.Values)-1; i++ {
		if stack.Values[i] > stack.Values[i+1] {
			return false
		}
	}
	return true
}

func ParseInputValues(args []string) (*IntegerStack, error) {
	stackA := &IntegerStack{}
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Invalid integer '%v'", err)
		}
		stackA.PushBack(value)
	}
	return stackA, nil
}
