package stack

import (
	"errors"
)

type data interface{}

type queue struct {
	data      []data
	size      int
	start     int
	end       int
	queueSize int
}

var ErrQueueOverflow = errors.New("queue_overflow")
var ErrQueueUnderflow = errors.New("queue_underflow")

func NewQueue(queueSize int) *queue {
	data := make([]data, queueSize)

	return &queue{
		queueSize: queueSize,
		size:      0,
		start:     0,
		end:       0,
		data:      data,
	}
}

func (q *queue) Push(d data) error {
	if q.size >= q.queueSize {
		return ErrQueueOverflow
	}
	q.data[q.end] = d
	q.end++
	q.size++

	return nil
}

func (q *queue) Pop() (d data, err error) {
	if q.size <= 0 {
		err = ErrQueueUnderflow
		return
	}
	d = q.data[q.start]
	q.data[q.start] = nil
	q.start++
	q.size--

	return
}
