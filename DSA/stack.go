package main

import (
	"errors"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is Empty")
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]

	s.elements = s.elements[:topIndex]
	return topElement, nil

}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// peek returns top element without removing it

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

/*
func main() {
	stack := Stack{}

	// Push elements to the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Peek at the top element
	top, _ := stack.Peek()
	fmt.Println("Top element:", top)

	// Pop elements from the stack
	for !stack.IsEmpty() {
		element, _ := stack.Pop()
		fmt.Println("Popped element:", element)
	}
}
*/
