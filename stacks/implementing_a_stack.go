package main

import (
	"errors"
	"fmt"
)

// stack our representation
type Stack struct {
	Elements []int
}

// add -method push
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// remove top elements on stack and return error if tack is empty
func (s *Stack) Pop() (int, error) {

	var lastelement int
	if len(s.Elements) == 0 {
		return 0, errors.New("Stack is empty")
	} else {
		/*
			// there is no any remove but we get last element whit this
				lastelementIndex := len(s.Elements) - 1
				lastelement = s.Elements[lastelementIndex]
		*/

		lastElementIndex := len(s.Elements) - 1
		lastelement, s.Elements = s.Elements[lastElementIndex], s.Elements[:lastElementIndex]
	}
	return lastelement, nil

}

func (s *Stack) Peek() (int, error) {
	var lastelement int
	if len(s.Elements) == 0 {
		return 0, errors.New("stack is clear-empty")
	} else {
		lastelementIndex := len(s.Elements) - 1
		lastelement = s.Elements[lastelementIndex]

	}

	return lastelement, nil
}

func (s *Stack) Length() int {
	return len(s.Elements)
}

func (s *Stack) isEmpty() bool {
	return s.Length() == 0
}

func main() {

	fmt.Println("Stack Section")
	fmt.Println("###################################################################")

	stack := Stack{}
	fmt.Println("###################################################################")

	fmt.Println(stack.isEmpty())
	if stack.isEmpty() {
		for i := 0; i <= 100; i++ {
			stack.Push(i)
		}
	}
	fmt.Println("###################################################################")

	fmt.Println(stack)
	fmt.Println("###################################################################")

	if !(stack.isEmpty()) {
		for i := 0; i <= 50; i++ {

			stack.Pop()
		}
	}
	fmt.Println("###################################################################")
	fmt.Println("Now our stack pop 50 elements our size : ", stack.Length())
	fmt.Println(stack)
	fmt.Println("###################################################################")
	elem, _ := stack.Peek()
	fmt.Println(" Last element on stack is :", elem)
	fmt.Println("###################################################################")

	popedelem, _ := stack.Pop()
	fmt.Println("Last deleted element is ", popedelem)
	fmt.Println("###################################################################")

}
