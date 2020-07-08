package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}

	fmt.Println("Stack created with size ", stack.Size())
	fmt.Println("Empty? ", stack.IsEmpty())

	stack.Insert("Go")
	stack.Insert(2009)
	stack.Insert(3.14)
	stack.Insert("End")

	fmt.Println("Stack after insert four items with size ", stack.Size())
	fmt.Println("Empty? ", stack.IsEmpty())

	for !stack.IsEmpty() {
		item, _ := stack.Remove()
		fmt.Println("Removing ", item)
		fmt.Println("Size ", stack.Size())
		fmt.Println("Empty? ", stack.IsEmpty())
	}

	_, err := stack.Remove()
	
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

type Stack struct {
	items []interface{}
}

func (stack Stack) Size() int {
	return len(stack.items)
}

func (stack Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Insert(item interface{}) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Remove() (interface{}, error) {
	
	if stack.IsEmpty() {
		return nil, errors.New("Empty stack!")
	}

	item := stack.items[stack.Size() - 1]
	stack.items = stack.items[:stack.Size() - 1]
	
	return item, nil
	
}
