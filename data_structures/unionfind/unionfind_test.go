package unionfind

import (
	"fmt"
	"testing"
)

func TestUnionFind_Find(t *testing.T) {
	uf := NewUnionFind()

	assert := func(p, q int, expect bool) {
		if uf.Connected(p, q) != expect {
			t.Fatalf("connected %d %d: !%t", p, q, expect)
		} else {
			t.Logf("connected %d %d: %t", p, q, expect)
		}
	}

	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 4)
	uf.Union(5, 6)
	uf.Union(7, 8)
	uf.Union(9, 10)

	assert(2, 1, true)
	assert(4, 5, false)
	assert(1, 4, true)
	assert(4, 1, true)

	uf.Union(5, 4)
	fmt.Println(uf)
	assert(6, 1, true)
}
