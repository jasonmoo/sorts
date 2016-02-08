package sorts

import "sort"

func Heap(a []int) {

	for len(a) > 3 {
		heapify(a)
		a = a[1:]
	}

	heapify(a)

	a = a[0:cap(a)]

	var r []int
	for _, n := range a {
		r = append(r, n)
		a = a[1:]
		heapify(a)
	}

}

func up(h sort.Interface, left int) {
	for {
		parent := (left - 1) / 2 // parent
		if parent == left || !h.Less(left, parent) {
			break
		}
		h.Swap(parent, left)
		left = parent
	}
}

func down(h sort.Interface, parent, n int) {
	for {
		left := 2*parent + 1
		if left >= n || left < 0 { // left < 0 after int overflow
			break
		}
		larger := left // left child
		if right := left + 1; right < n && !h.Less(left, right) {
			larger = right // = 2*parent + 2  // right child
		}
		if !h.Less(larger, parent) {
			break
		}
		h.Swap(parent, larger)
		parent = larger
	}
}
