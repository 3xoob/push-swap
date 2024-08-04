package PSFunc

func pushToPrimaryStack(primaryStack, secondaryStack *[]int) {
	if len(*secondaryStack) > 0 {
		*primaryStack = append([]int{(*secondaryStack)[0]}, *primaryStack...)
		*secondaryStack = (*secondaryStack)[1:]
	}
}

func pushToSecondaryStack(primaryStack, secondaryStack *[]int) {
	if len(*primaryStack) > 0 {
		*secondaryStack = append([]int{(*primaryStack)[0]}, *secondaryStack...)
		*primaryStack = (*primaryStack)[1:]
	}
}

func swapPrimaryTopTwo(primaryStack *[]int) {
	if len(*primaryStack) > 1 {
		(*primaryStack)[0], (*primaryStack)[1] = (*primaryStack)[1], (*primaryStack)[0]
	}
}

func swapSecondaryTopTwo(secondaryStack *[]int) {
	if len(*secondaryStack) > 1 {
		(*secondaryStack)[0], (*secondaryStack)[1] = (*secondaryStack)[1], (*secondaryStack)[0]
	}
}

func swapTopTwoInBoth(primaryStack, secondaryStack *[]int) {
	swapPrimaryTopTwo(primaryStack)
	swapSecondaryTopTwo(secondaryStack)
}

func rotatePrimaryUp(primaryStack *[]int) {
	if len(*primaryStack) > 1 {
		*primaryStack = append((*primaryStack)[1:], (*primaryStack)[0])
	}
}

func rotateSecondaryUp(secondaryStack *[]int) {
	if len(*secondaryStack) > 1 {
		*secondaryStack = append((*secondaryStack)[1:], (*secondaryStack)[0])
	}
}

func rotateBothUp(primaryStack, secondaryStack *[]int) {
	rotatePrimaryUp(primaryStack)
	rotateSecondaryUp(secondaryStack)
}

func rotatePrimaryDown(primaryStack *[]int) {
	if len(*primaryStack) > 1 {
		*primaryStack = append([]int{(*primaryStack)[len(*primaryStack)-1]}, (*primaryStack)[:len(*primaryStack)-1]...)
	}
}

func rotateSecondaryDown(secondaryStack *[]int) {
	if len(*secondaryStack) > 1 {
		*secondaryStack = append([]int{(*secondaryStack)[len(*secondaryStack)-1]}, (*secondaryStack)[:len(*secondaryStack)-1]...)
	}
}

func rotateBothDown(primaryStack, secondaryStack *[]int) {
	rotatePrimaryDown(primaryStack)
	rotateSecondaryDown(secondaryStack)
}
