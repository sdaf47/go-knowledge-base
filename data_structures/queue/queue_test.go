package stack

import (
	"testing"
)

func TestList_NewQueue(t *testing.T) {
	var testStack = NewQueue(100)

	for i := 0; i <= 100; i++ {
		err := testStack.Push(i)
		if err != nil && i < 100 {
			t.Fatal(err.Error())
		}
	}

	for i := 100; i <= 0; i-- {
		d, err := testStack.Pop()
		if err != nil && i != 0 {
			t.Fatal("err must be empty_stack")
		}
		if i != d {
			t.Fatal("i must be equil d")
		}
	}

}
