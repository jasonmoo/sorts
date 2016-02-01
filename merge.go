package sorts

func Merge(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	aa := append([]int(nil), a[:len(a)/2]...)
	bb := append([]int(nil), a[len(a)/2:]...)

	aa = Merge(aa)
	bb = Merge(bb)

	return merge(aa, bb)

}

func merge(a, b []int) []int {

	r := make([]int, 0, len(a)+len(b))

	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			r = append(r, a[0])
			a = a[1:]
		} else {
			r = append(r, b[0])
			b = b[1:]
		}
	}
	r = append(r, a...)
	r = append(r, b...)

	return r

}
