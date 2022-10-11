package main

import (
	"errors"
	"fmt"
)

// Representation of our set data structure

type Set struct {
	Elements map[string]struct{}
	// new values will stored in string
}

func NewSet() *Set {
	//var set Set
	//set.element = make(map[string]struct{})
	//return &set
	return &Set{
		Elements: map[string]struct{}{},
	}
}

// helper methods based on struct

// adding new element

func (set *Set) Add(ele string) {
	set.Elements[ele] = struct{}{}
}

// Delete - removes an element from our set if it exists
func (set *Set) Delete(elem string) error {
	if _, exists := set.Elements[elem]; !exists {
		return errors.New("element not present in Set")
	}
	delete(set.Elements, elem)
	return nil
}

// check if element is exist in the set

func (set *Set) Contains(elem string) bool {
	_, exists := set.Elements[elem]
	return exists
}

// displaying all elements of set

func (set *Set) List() {
	for k, _ := range set.Elements {
		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Set")
	mySet := NewSet()
	mySet.Add("Kiran")
	mySet.Add("Arun")
	mySet.Add("Kshirsagar")
	fmt.Println()
	mySet.List()
	mySet.Add("Kiran")
	fmt.Println()
	mySet.List()
	mySet.Add("temp")
	fmt.Println()
	mySet.List()
	mySet.Delete("temp")
	fmt.Println()
	mySet.List()
}
