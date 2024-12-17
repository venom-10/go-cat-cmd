package main

import (
	"fmt"
	"go-tutorial/cmd"
	"os"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// Append adds a new element with the given value to the end of the list.
func (l *List[T]) Append(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
	println(current.next)
}

// Insert inserts a new element with the given value after the first occurrence of a target value.
func (l *List[T]) Insert(after T, value T) bool {
	current := l
	for current != nil {
		if current.val == after {
			nextNode := current.next
			current.next = &List[T]{val: value, next: nextNode}
			return true
		}
		current = current.next
	}
	return false
}

// Delete removes the first occurrence of the specified value from the list.
func (l *List[T]) Delete(value T) bool {
	if l == nil {
		return false
	}

	if l.val == value {
		*l = *l.next
		return true
	}

	prev, current := l, l.next
	for current != nil {
		if current.val == value {
			prev.next = current.next
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

// Display prints all the elements of the list.
func (l *List[T]) Display() []T {
	var result []T
	current := l
	for current != nil {
		result = append(result, current.val)
		fmt.Printf("Current List: %+v\n", *current)
		current = current.next
	}
	return result
}

// func main() {
// list := &List[int]{val: 1}
// fmt.Printf("Address of list: %p\n", list) // Print the address of the list instance

// list.Append(2)
// list.Append(7)
// list.Append(888888888888888888)
// list.Insert(2, 5)
// list.Delete(3)

// elements := list.Display()
// for _, v := range elements {
// 	println(v)
// }
// ptr1 := uintptr(0xc000118000)
// ptr2 := uintptr(0xc000118110)

// Calculate the difference
// diff := ptr2 - ptr1
// fmt.Printf("Difference in bytes: %d\n", diff)

// head := pkg.Parent(2)
// head.AppendRight(3)
// head.AppendRight(4)
// head.AppendRight(4)
// head.AppendRight(5)
// head = head.AppendLeft(1)
// head.TraverseForward()

// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycat <filename>")
		os.Exit(1)
	}

	// Get filename from command-line arguments
	filename := os.Args[1]

	// Read and prsint file contents
	err := cmd.CatFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
