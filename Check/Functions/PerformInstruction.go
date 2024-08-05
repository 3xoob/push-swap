package CFunc

import "fmt"

func PerformInstruction(instruction string, stackA, stackB *NumStack) error {
	switch instruction {
	case "sa":
		return stackA.SwapTopTwo()
	case "sb":
		return stackB.SwapTopTwo()
	case "ss":
		if err := stackA.SwapTopTwo(); err != nil {
			return fmt.Errorf("Error")
		}
		return stackB.SwapTopTwo()
	case "pa":
		value, err := stackB.PopFront()
		if err != nil {
			return fmt.Errorf("Error")
		}
		stackA.PushFront(value)
	case "pb":
		value, err := stackA.PopFront()
		if err != nil {
			return fmt.Errorf("Error")
		}
		stackB.PushFront(value)
	case "ra":
		return stackA.RotateUp()
	case "rb":
		return stackB.RotateUp()
	case "rr":
		if err := stackA.RotateUp(); err != nil {
			return fmt.Errorf("Error")
		}
		return stackB.RotateUp()
	case "rra":
		return stackA.RotateDown()
	case "rrb":
		return stackB.RotateDown()
	case "rrr":
		if err := stackA.RotateDown(); err != nil {
			return fmt.Errorf("Error")
		}
		return stackB.RotateDown()
	default:
		return fmt.Errorf("Error")
	}
	return nil
}
