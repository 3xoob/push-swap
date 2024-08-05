package CFunc

import (
	"fmt"
)

// Push an element to the front of the stack
func (s *NumStack) PushFront(value int) {
	s.Elements = append([]int{value}, s.Elements...)
}

// Push an element to the back of the stack
func (s *NumStack) PushBack(value int) {
	s.Elements = append(s.Elements, value)
}

// Pop an element from the front of the stack
func (s *NumStack) PopFront() (int, error) {
	if len(s.Elements) == 0 {
		return 0, fmt.Errorf("Error")
	}
	value := s.Elements[0]
	s.Elements = s.Elements[1:]
	return value, nil
}

// Swap the top two elements of the stack
func (s *NumStack) SwapTopTwo() error {
	if len(s.Elements) < 2 {
		return fmt.Errorf("Error")
	}
	topTwo := s.Elements[:2]
	s.Elements[0], s.Elements[1] = topTwo[1], topTwo[0]
	return nil
}

// Rotate the stack upwards (move the first element to the last position)
func (s *NumStack) RotateUp() error {
	if len(s.Elements) == 0 {
		return fmt.Errorf("Error")
	}
	firstValue := s.Elements[0]
	s.Elements = append(s.Elements[1:], firstValue)
	return nil
}

// Rotate the stack downwards (move the last element to the first position)
func (s *NumStack) RotateDown() error {
	if len(s.Elements) == 0 {
		return fmt.Errorf("Error")
	}
	lastValue := s.Elements[len(s.Elements)-1]
	s.Elements = append([]int{lastValue}, s.Elements[:len(s.Elements)-1]...)
	return nil
}
