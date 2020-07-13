package main

import "fmt"

type GenericList []interface{}

func (list *GenericList) RemoveByIndex(index int) interface{} {
	l := *list
	removed := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return removed
}

func (list *GenericList) RemoveFromStart() interface{} {
	return list.RemoveByIndex(0)
}

func (list *GenericList) RemoveFromEnd() interface{} {
	return list.RemoveByIndex(len(*list)-1)
}

func main() {
	list := GenericList{
		1,
		"coffee",
		42,
		true,
		23,
		"ball",
		3.14,
		false,
	}

	fmt.Printf("Original list:\n%v\n\n", list)

	fmt.Printf("Removing from start: %v, new list:\n%v\n", list.RemoveFromStart(), list)

	fmt.Printf("Removing from end: %v, new list:\n%v\n", list.RemoveFromEnd(), list)

	fmt.Printf("Removing by index 3: %v, new list:\n%v\n", list.RemoveByIndex(3), list)

	fmt.Printf("Removing by index 0: %v, new list:\n%v\n", list.RemoveByIndex(0), list)

	fmt.Printf("Removing by last index: %v, new list:\n%v\n", list.RemoveByIndex(len(list)-1), list)
}
