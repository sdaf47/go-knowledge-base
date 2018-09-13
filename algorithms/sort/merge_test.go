package sort

import (
	"testing"
	"math/rand"
	"sort"
)

var testIntArray = rendomIntArray(10000)

func TestMergeSort(t *testing.T) {
	res := MergeSort(testIntArray)

	if !sort.IntsAreSorted(res) {
		t.Fatalf("array does not sorted")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	MergeSort(testIntArray)
}

func rendomIntArray(size int) (result []int) {
	rendomizer := rand.NewSource(42)

	for i := 0; i < size; i++ {
		result = append(result, int(rendomizer.Int63()))
	}

	return
}
