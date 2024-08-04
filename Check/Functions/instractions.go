package CFunc

import (
	"fmt"
)

func PerformInstruction(instruction string, stackA, stackB *IntegerStack) error {
	switch instruction {
	case "sa":
		return stackA.SwapTopTwo()
	case "sb":
		return stackB.SwapTopTwo()
	case "ss":
		if err := stackA.SwapTopTwo(); err != nil {
			return err
		}
		return stackB.SwapTopTwo()
	case "pa":
		value, err := stackB.PopFront()
		if err != nil {
			return fmt.Errorf("ERROR: Cannot pop from Stack B: %v", err)
		}
		stackA.PushFront(value)
	case "pb":
		value, err := stackA.PopFront()
		if err != nil {
			return fmt.Errorf("ERROR: Cannot pop from Stack A: %v", err)
		}
		stackB.PushFront(value)
	case "ra":
		return stackA.RotateUp()
	case "rb":
		return stackB.RotateUp()
	case "rr":
		if err := stackA.RotateUp(); err != nil {
			return err
		}
		return stackB.RotateUp()
	case "rra":
		return stackA.RotateDown()
	case "rrb":
		return stackB.RotateDown()
	case "rrr":
		if err := stackA.RotateDown(); err != nil {
			return err
		}
		return stackB.RotateDown()
	default:
		return fmt.Errorf("ERROR: Invalid instruction '%s'", instruction)
	}
	return nil
}

func (stack *IntegerStack) PushFront(value int) {
	stack.Values = append([]int{value}, stack.Values...)
}

func (stack *IntegerStack) PushBack(value int) {
	stack.Values = append(stack.Values, value)
}

func (stack *IntegerStack) PopFront() (int, error) {
	if len(stack.Values) == 0 {
		return 0, fmt.Errorf("ERROR: Stack underflow, cannot pop from empty stack")
	}
	value := stack.Values[0]
	stack.Values = stack.Values[1:]
	return value, nil
}

func (stack *IntegerStack) SwapTopTwo() error {
	if len(stack.Values) < 2 {
		return fmt.Errorf("ERROR: Not enough elements to swap")
	}
	stack.Values[0], stack.Values[1] = stack.Values[1], stack.Values[0]
	return nil
}

func (stack *IntegerStack) RotateUp() error {
	if len(stack.Values) == 0 {
		return fmt.Errorf("ERROR: Cannot rotate empty stack")
	}
	value := stack.Values[0]
	copy(stack.Values, stack.Values[1:])
	stack.Values[len(stack.Values)-1] = value
	return nil
}

func (stack *IntegerStack) RotateDown() error {
	if len(stack.Values) == 0 {
		return fmt.Errorf("ERROR: Cannot reverse rotate empty stack")
	}
	value := stack.Values[len(stack.Values)-1]
	stack.Values = append([]int{value}, stack.Values[:len(stack.Values)-1]...)
	return nil
}
