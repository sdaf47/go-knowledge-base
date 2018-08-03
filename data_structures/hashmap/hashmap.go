package hashmap

import (
	"errors"
	"encoding/gob"
	"bytes"
	"hash/fnv"
)

type HashFunc func(blockSize int, key Key) int

type Key interface{}

type Statistic struct {
	collisionCnt  int
	allocationCnt int
}

var (
	ErrInvalidKey = errors.New("key_error")
	ErrCollision  = errors.New("collision_error")
)

type HashMapper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Statistic() Statistic
	Count() int
}

type entry struct {
	key   Key
	value interface{}
	next  *entry
}

type hashMapper struct {
	data      []*entry
	hash      HashFunc
	blockSize int
	count     int

	stat Statistic
}

func hash(blockSize int, key Key) int {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		panic(err)
	}

	h := fnv.New32a()
	_, err = h.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	return int(h.Sum32()) % blockSize
}

func (hm *hashMapper) malloc(size int) {
	hm.blockSize += size
	od := hm.data
	nd := make([]*entry, hm.blockSize)

	for _, el := range od {
		for oitem := el; oitem != nil; oitem = oitem.next {
			i := hm.hash(hm.blockSize, oitem.key)
			if nitem := nd[i]; nitem != nil {
				for ; nitem != nil; nitem = nitem.next {
					if nitem.next == nil {
						nitem.next = &entry{oitem.key, oitem.value, nil}
						break
					}
				}
			} else {
				nd[i] = &entry{oitem.key, oitem.value, nil}
			}
		}
	}

	hm.data = nd

	hm.stat.allocationCnt++
}

func NewHashMap(blockSize int, fn HashFunc) (hm HashMapper) {
	if fn == nil {
		fn = hash
	}

	m := &hashMapper{
		blockSize: blockSize,
		data:      make([]*entry, blockSize),
		hash:      fn,
	}

	hm = m

	return
}

func (hm *hashMapper) Set(key Key, value interface{}) (err error) {
	i := hm.hash(hm.blockSize, key)
	if stack := hm.data[i]; stack != nil {
		var item *entry
		for item = stack; item != nil; item = item.next {
			if item.key == key {
				item.value = value
				break
			}
			if item.next == nil {
				hm.stat.collisionCnt++
				item.next = &entry{key, value, nil}
				hm.count++
				break
			}
		}
		return
	}
	hm.data[i] = &entry{key, value, nil}
	hm.count++

	if hm.count >= hm.blockSize {
		hm.malloc(hm.blockSize * 2)
	}

	return
}

func (hm *hashMapper) Get(key Key) (value interface{}, err error) {
	i := hm.hash(hm.blockSize, key)
	if hm.data[i] == nil {
		return
	}
	for item := hm.data[i]; item != nil; item = item.next {
		if key == item.key {
			value = item.value
			return
		}
	}
	err = ErrInvalidKey

	return
}

func (hm *hashMapper) Unset(key Key) (err error) {
	i := hm.hash(hm.blockSize, key)
	if hm.data[i] == nil {
		return
	}
	if hm.data[i].key == key {
		hm.data[i] = nil
		hm.count--
		return
	}
	for item := hm.data[i]; item.next != nil; item = item.next {
		if item.next.key == key {
			item.next = item.next.next
			hm.count--
			return
		}
	}
	err = ErrInvalidKey

	return
}

func (hm *hashMapper) Count() (cnt int) {
	return hm.count
}

func (hm *hashMapper) Statistic() Statistic {
	return hm.stat
}
