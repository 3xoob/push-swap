package PSFunc

func SortStacks(primaryStack, secondaryStack *[]int) []string {
	var operationsLog []string

	if len(*primaryStack) == 2 {
		// Sorting a two-element stack
		if (*primaryStack)[0] > (*primaryStack)[1] {
			swapPrimaryTopTwo(primaryStack)
			operationsLog = append(operationsLog, "sa")
		}
	} else if len(*primaryStack) == 3 {
		// Sorting a three-element stack
		for !IsSorted(*primaryStack) {
			switch {
			case (*primaryStack)[0] > (*primaryStack)[1] && (*primaryStack)[0] > (*primaryStack)[2]:
				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")
			case (*primaryStack)[1] > (*primaryStack)[0] && (*primaryStack)[1] > (*primaryStack)[2]:
				rotatePrimaryDown(primaryStack)
				operationsLog = append(operationsLog, "rra")
			default:
				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")
			}
		}
	} else {
		// Sorting a stack with more than three elements
		for !IsSorted(*primaryStack) {
			minIdx := findMinIndex(*primaryStack)

			switch {
			case minIdx == 0:
				pushToSecondaryStack(primaryStack, secondaryStack)
				operationsLog = append(operationsLog, "pb")
			case minIdx == 1:
				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")
			case minIdx <= len(*primaryStack)/2:
				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")
			default:
				rotatePrimaryDown(primaryStack)
				operationsLog = append(operationsLog, "rra")
			}
		}
	}

	// Moving all elements back to the primary stack
	for len(*secondaryStack) > 0 {
		pushToPrimaryStack(primaryStack, secondaryStack)
		operationsLog = append(operationsLog, "pa")
	}

	return operationsLog
}

func IsSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}
