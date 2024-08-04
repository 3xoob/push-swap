package PSFunc

func ContainsDuplicates(stack []int) bool {
	seenNumbers := make(map[int]bool)
	for _, number := range stack {
		if seenNumbers[number] {
			return true
		}
		seenNumbers[number] = true
	}
	return false
}

func findMinIndex(stack []int) int {
	minIdx := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

func findMaxIndex(stack []int) int {
	maxIdx := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] > stack[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
