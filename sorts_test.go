package sorts

import (
	"fmt"
	"testing"
)

var a = []int{2, 7, 9, 0, 6, 8, 1, 5, 4, 3}

func TestBubble(t *testing.T) {

	a := append([]int(nil), a...)

	Bubble(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestBucket(t *testing.T) {

	a := append([]int(nil), a...)

	Bucket(a, 3)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestBucketInsertion(t *testing.T) {

	a := append([]int(nil), a...)

	Bucket(a, 3)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestComb(t *testing.T) {

	a := append([]int(nil), a...)

	Comb(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestCounting(t *testing.T) {

	b := make([]byte, len(a))
	for i, n := range a {
		b[i] = byte(n)
	}

	Counting(b)

	for i, n := range b {
		if i != int(n) {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(b)

}

func TestHeap(t *testing.T) {
	t.Skip("unimplemented")

	a := append([]int(nil), a...)

	// Heap(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestQuick(t *testing.T) {
	t.Skip("unimplemented")

	a := append([]int(nil), a...)

	// Quick(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestQuick3(t *testing.T) {
	t.Skip("unimplemented")

	a := append([]int(nil), a...)

	// Quick3(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestRadix(t *testing.T) {
	t.Skip("unimplemented")

	a := append([]int(nil), a...)

	// Radix(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestInsertion(t *testing.T) {

	a := append([]int(nil), a...)

	Insertion(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestShell(t *testing.T) {

	a := append([]int(nil), a...)

	Shell(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}

func TestMerge(t *testing.T) {

	a := append([]int(nil), a...)

	a = Merge(a)

	for i, n := range a {
		if i != n {
			t.Errorf("Expected %d, got %d", i, n)
		}
	}

	fmt.Println(a)

}
