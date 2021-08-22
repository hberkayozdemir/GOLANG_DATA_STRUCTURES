package main

import (
	"errors"
	"fmt"
)

// set - our representation of a set data structure

type Set struct {
	Elements map[string]struct{}
}

// new elements will be store on this string map
// empty struct some thing like a boolean check exist

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// for create a new set

// Add- adds an element to our Set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

func (s *Set) Delete(elem string) error {
	if _, exist := s.Elements[elem]; !exist {
		return errors.New("Elements not represents in this set")
	}
	delete(s.Elements, elem)

	return nil
}

func (s *Set) Contains(elem string) bool {
	_, exist := s.Elements[elem]
	return exist
}

func (s *Set) List() {

	for k, _ := range s.Elements {

		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Sets Implementation Tutorial")
	mySet := NewSet()
	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.Add("Jupiter")

	mySet.List()
	fmt.Println("------ After mySet.Delete('Venus')------- ")
	mySet.Delete("Venus")
	mySet.List()

	exist := mySet.Contains("Earth")
	if !exist {
		fmt.Println("Set dont have this element")
	} else {
		fmt.Println("Set contains this element")
	}
}
