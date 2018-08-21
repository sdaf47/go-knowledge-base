package stack

import (
	"errors"
)

type data interface{}

type stack struct {
	data      []data
	size      int
	stackSize int
}

var ErrStackOverflow = errors.New("stack_overflow")
var ErrStackUnderflow = errors.New("stack_underflow")

func NewStack(stackSize int) *stack {
	data := make([]data, stackSize)

	return &stack{
		stackSize: stackSize,
		size:      0,
		data:      data,
	}
}

func (s *stack) push(d data) error {
	if s.size >= s.stackSize {
		return ErrStackOverflow
	}

	s.data[s.size] = d
	s.size++

	return nil
}

func (s *stack) pop() (d data, err error) {
	if s.size <= 0 {
		err = ErrStackUnderflow
		return
	}
	s.size--
	d = s.data[s.size]
	s.data[s.size] = nil

	return
}

func (s *stack) peek() (d data, err error) {
	if s.size <= 0 {
		err = ErrStackUnderflow
		return
	}
	d = s.data[s.size-1]

	return
}
