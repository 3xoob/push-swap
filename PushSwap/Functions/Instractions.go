package PSFunc

func pushToPrimaryStack(mainStack, auxStack *[]int) {
	if len(*auxStack) == 0 {
		return
	}
	topAux := (*auxStack)[0]
	*mainStack = append([]int{topAux}, *mainStack...)
	*auxStack = (*auxStack)[1:]
}

func pushToSecondaryStack(mainStack, auxStack *[]int) {
	if len(*mainStack) == 0 {
		return
	}
	topMain := (*mainStack)[0]
	*auxStack = append([]int{topMain}, *auxStack...)
	*mainStack = (*mainStack)[1:]
}

func swapPrimaryTopTwo(mainStack *[]int) {
	if len(*mainStack) < 2 {
		return
	}
	firstElem, secondElem := (*mainStack)[0], (*mainStack)[1]
	(*mainStack)[0], (*mainStack)[1] = secondElem, firstElem
}

func swapSecondaryTopTwo(auxStack *[]int) {
	if len(*auxStack) < 2 {
		return
	}
	firstElem, secondElem := (*auxStack)[0], (*auxStack)[1]
	(*auxStack)[0], (*auxStack)[1] = secondElem, firstElem
}

func swapTopTwoInBoth(mainStack, auxStack *[]int) {
	swapPrimaryTopTwo(mainStack)
	swapSecondaryTopTwo(auxStack)
}

func rotatePrimaryUp(mainStack *[]int) {
	if len(*mainStack) <= 1 {
		return
	}
	firstElem := (*mainStack)[0]
	*mainStack = append((*mainStack)[1:], firstElem)
}

func rotateSecondaryUp(auxStack *[]int) {
	if len(*auxStack) <= 1 {
		return
	}
	firstElem := (*auxStack)[0]
	*auxStack = append((*auxStack)[1:], firstElem)
}

func rotateBothUp(mainStack, auxStack *[]int) {
	rotatePrimaryUp(mainStack)
	rotateSecondaryUp(auxStack)
}

func rotatePrimaryDown(mainStack *[]int) {
	if len(*mainStack) <= 1 {
		return
	}
	lastElem := (*mainStack)[len(*mainStack)-1]
	*mainStack = append([]int{lastElem}, (*mainStack)[:len(*mainStack)-1]...)
}

func rotateSecondaryDown(auxStack *[]int) {
	if len(*auxStack) <= 1 {
		return
	}
	lastElem := (*auxStack)[len(*auxStack)-1]
	*auxStack = append([]int{lastElem}, (*auxStack)[:len(*auxStack)-1]...)
}

func rotateBothDown(mainStack, auxStack *[]int) {
	rotatePrimaryDown(mainStack)
	rotateSecondaryDown(auxStack)
}
