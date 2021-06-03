package DataStruct

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	q.Push(7)
	q.Push(8)
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.elements[i])
	}
	fmt.Println(q.len,q.IsEmpty(),q.Head(),q.Tail())
	fmt.Println(q.Pop(),q.Pop())
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.elements[i])
	}
	fmt.Println(q.len,q.IsEmpty(),q.Head(),q.Tail())
	fmt.Println(q.cap)
	q.Expansion()
	fmt.Println(q.cap)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	q.Push(7)
	q.Push(8)
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.elements[i])
	}
	fmt.Println(q.len,q.IsEmpty(),q.Head(),q.Tail())
	fmt.Println(q.Pop(),q.Pop())
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.elements[i])
	}
	fmt.Println(q.len,q.IsEmpty(),q.Head(),q.Tail())
}