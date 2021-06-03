package DataStruct

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push("3")
	fmt.Println(s.Len())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop(),s.Pop(),s.Pop())
	fmt.Println(s.IsEmpty())
}