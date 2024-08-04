package CFunc

import (
	"fmt"
)


func (stack *IntegerStack) PushFront(value int) {
	stack.Values = append([]int{value}, stack.Values...)
}

func (stack *IntegerStack) PushBack(value int) {
	stack.Values = append(stack.Values, value)
}

func (stack *IntegerStack) PopFront() (int, error) {
	if len(stack.Values) == 0 {
		return 0, fmt.Errorf("Error")
	}
	value := stack.Values[0]
	stack.Values = stack.Values[1:]
	return value, nil
}

func (stack *IntegerStack) SwapTopTwo() error {
	if len(stack.Values) < 2 {
		return fmt.Errorf("Error")
	}
	stack.Values[0], stack.Values[1] = stack.Values[1], stack.Values[0]
	return nil
}

func (stack *IntegerStack) RotateUp() error {
	if len(stack.Values) == 0 {
		return fmt.Errorf("Error")
	}
	value := stack.Values[0]
	copy(stack.Values, stack.Values[1:])
	stack.Values[len(stack.Values)-1] = value
	return nil
}

func (stack *IntegerStack) RotateDown() error {
	if len(stack.Values) == 0 {
		return fmt.Errorf("Error")
	}
	value := stack.Values[len(stack.Values)-1]
	stack.Values = append([]int{value}, stack.Values[:len(stack.Values)-1]...)
	return nil
}
