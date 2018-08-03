package linkedlist

import (
	"errors"
)

type entry struct {
	value interface{}
	next  *entry
	prev  *entry
}

type List struct {
	cnt  uint
	size uint
	head *entry
	tail *entry
}

var ErrInvalidIndex = errors.New("invalid_index")

func NewLinkedList(size uint) (list *List) {
	return new(List)
}

func (l *List) Add(v interface{}) {
	e := new(entry)
	e.value = v

	if l.head == nil {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		e.prev = l.tail
		l.tail = e
	}

	l.cnt++

	return
}

func (l *List) Get(key uint) (v interface{}, err error) {
	e, err := l.get(key)
	if err != nil {
		return
	}

	v = e.value
	return
}

func (l *List) Remove(key uint) (err error) {
	if key == 0 {
		e := l.head.next
		l.head = e
		e.prev = nil
	} else if key == l.cnt {
		e := l.tail.prev
		l.tail = e
		e.next = nil
	}

	e, err := l.get(key)
	if err != nil {
		return
	}

	next := e.next
	prev := e.prev

	prev.next = e.next
	next.prev = e.prev

	return
}

func (l *List) get(key uint) (e *entry, err error) {
	mid := l.cnt >> 1

	if key > l.cnt {
		err = ErrInvalidIndex
		return
	}

	if key < mid {
		e = l.head

		for i := uint(0); i < key; i++ {
			e = e.next
		}
	} else {
		e = l.tail

		for i := l.cnt - 1; i > key; i-- {
			e = e.prev
		}
	}

	return
}

func (l *List) Clear() {
	l.head = nil
	l.tail = nil
	l.cnt = 0
}

func (l *List) Count() (cnt uint) {
	return l.cnt
}

func (l *List) Contains(interface{}) (ok bool, err error) {
	return
}

func (l *List) Iterator() {

}
