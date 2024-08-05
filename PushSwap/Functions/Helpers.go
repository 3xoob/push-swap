package PSFunc

func findMaxIndex(stack []int) int {
	if len(stack) == 0 {
		return -1
	}

	maxIndex := 0
	for i, num := range stack {
		if num > stack[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func findMinIndex(stack []int) int {
	if len(stack) == 0 {
		return -1
	}

	minIndex := 0
	for i, num := range stack {
		if num < stack[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func ContainsDuplicates(stack []int) bool {
	seen := make(map[int]struct{})
	for _, num := range stack {
		if _, found := seen[num]; found {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
