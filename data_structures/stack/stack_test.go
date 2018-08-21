package stack

import (
	"testing"
)

func TestList_NewStack(t *testing.T) {
	var testStack = NewStack(100)

	for i := 0; i <= 100; i++ {
		err := testStack.Push(i)
		if err != nil && i < 100 {
			t.Fatal(err.Error())
		}
	}

	head, _ := testStack.Peek()
	if head != 99 {
		t.Fatal("99 must be last element of stack")
	}

	for i := 99; i >= 0; i-- {
		d, err := testStack.Pop()
		if err != nil && i != 0 {
			t.Fatal("err must be empty_stack")
		}
		if i != d {
			t.Fatal("i must be equil d")
		}
	}

}
