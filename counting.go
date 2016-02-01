package sorts

func Counting(a []uint8) {

	var m [255]int

	for _, n := range a {
		m[int(n)]++
	}

	a = a[:0]
	for n, ct := range m {
		for i := 0; i < ct; i++ {
			a = append(a, uint8(n))
		}
	}

}
