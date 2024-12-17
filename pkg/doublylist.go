package pkg

import "fmt"

type DList[T comparable] struct {
	prev *DList[T]
	next *DList[T]
	val  T
}

func Parent[T comparable](value T) *DList[T] {
	return &DList[T]{
		val:  value,
		prev: nil,
		next: nil,
	}
}

func (l *DList[T]) AppendRight(value T) {
	end := l
	for end.next != nil {
		end = end.next
	}
	newNode := &DList[T]{val: value}
	newNode.prev = end
	end.next = newNode
}

func (l *DList[T]) AppendLeft(value T) *DList[T] {
	for l.prev != nil {
		l = l.prev
	}
	newNode := &DList[T]{val: value}
	fmt.Printf("Initial current node value: %v, %v\n", l.val, newNode.val)
	newNode.prev = nil
	newNode.next = l
	l.prev = newNode
	return newNode
}

// Forward Traversal
func (l *DList[T]) TraverseForward() []T {
	fmt.Println("Starting TraverseForward...")
	var result []T
	current := l

	fmt.Printf("Initial current node value: %v\n", current.val)

	// Go to the start of the list
	// fmt.Println("Finding start of the list...")
	// for current.prev != nil {
	// 	fmt.Printf("Moving to previous node. Current value: %v\n", current.val)
	// 	current = current.prev
	// }
	// fmt.Printf("Start of list found. Start node value: %v\n", current.val)

	// Traverse forward
	// fmt.Println("Traversing forward...")
	for current != nil {
		fmt.Printf("Adding node value to result: %v\n", current.val)
		result = append(result, current.val)
		current = current.next
	}

	fmt.Printf("Traversal complete. Result: %v\n", result)
	return result
}

// Backward Traversal
func (l *DList[T]) TraverseBackward() []T {
	fmt.Println("Starting TraverseBackward...")
	var result []T
	current := l

	fmt.Printf("Initial current node value: %v\n", current.val)

	// Go to the start of the list
	// fmt.Println("Finding start of the list...")
	// for current.prev != nil {
	// 	fmt.Printf("Moving to previous node. Current value: %v\n", current.val)
	// 	current = current.prev
	// }
	// fmt.Printf("Start of list found. Start node value: %v\n", current.val)

	// Traverse forward
	// fmt.Println("Traversing forward...")
	for current != nil {
		fmt.Printf("Adding node value to result: %v\n", current.val)
		result = append(result, current.val)
		current = current.prev
	}

	fmt.Printf("Traversal complete. Result: %v\n", result)
	return result
}

// Alternate Traversal with Callback
func (l *DList[T]) Traverse(callback func(T)) {
	current := l

	// Go to the start of the list
	for current.prev != nil {
		current = current.prev
	}

	// Traverse and apply callback
	for current != nil {
		callback(current.val)
		current = current.next
	}
}
