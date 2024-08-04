package PSFunc

func SortStacks(primaryStack, secondaryStack *[]int) []string {
	var operationsLog []string

	if len(*primaryStack) == 2 {
		for !IsSorted(*primaryStack) {
			swapPrimaryTopTwo(primaryStack)
			operationsLog = append(operationsLog, "sa")
		}
	} else if len(*primaryStack) == 3 {
		minIdx := findMinIndex(*primaryStack)
		maxIdx := findMaxIndex(*primaryStack)

		if (minIdx == 0 && maxIdx == 1) || (minIdx == 2 && maxIdx == 0) {
			if minIdx == 0 && maxIdx == 1 {
				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")

				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")
			} else if minIdx == 2 && maxIdx == 0 {
				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")

				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")
			}
		} else {
			if minIdx == 1 && maxIdx == 0 {
				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")
			} else if minIdx == 1 && maxIdx == 2 {
				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")
			} else if minIdx == 2 && maxIdx == 1 {
				rotatePrimaryDown(primaryStack)
				operationsLog = append(operationsLog, "rra")
			}
		}
	} else {
		for !IsSorted(*primaryStack) {
			minIdx := findMinIndex(*primaryStack)

			if minIdx == 0 {
				pushToSecondaryStack(primaryStack, secondaryStack)
				operationsLog = append(operationsLog, "pb")
			} else if minIdx == 1 {
				swapPrimaryTopTwo(primaryStack)
				operationsLog = append(operationsLog, "sa")
			} else if minIdx <= len(*primaryStack)/2 {
				rotatePrimaryUp(primaryStack)
				operationsLog = append(operationsLog, "ra")
			} else {
				rotatePrimaryDown(primaryStack)
				operationsLog = append(operationsLog, "rra")
			}
		}
	}

	for len(*secondaryStack) > 0 {
		pushToPrimaryStack(primaryStack, secondaryStack)
		operationsLog = append(operationsLog, "pa")
	}

	return operationsLog
}

func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}
