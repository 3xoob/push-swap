package CFunc

import (
	"fmt"
	"strconv"
)

type NumStack struct {
	Elements []int
}

// Check if the stack is sorted in ascending order
func (s *NumStack) IsSortedAscending() bool {
	for i := range s.Elements[:len(s.Elements)-1] {
		if s.Elements[i] > s.Elements[i+1] {
			return false
		}
	}
	return true
}

// Parse string input and convert to an integer stack
func ParseInputValues(args []string) (*NumStack, error) {
	stack := &NumStack{}
	for _, strVal := range args {
		num, err := strconv.Atoi(strVal)
		if err != nil {
			return nil, fmt.Errorf("Error")
		}
		stack.Elements = append(stack.Elements, num)
	}
	return stack, nil
}
