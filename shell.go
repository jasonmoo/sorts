package sorts

func Shell(a []int) {

	h := 1
	for h < len(a) {
		h = (h * 3) + 1
	}

	for h > 0 {
		h /= 3

		for k := 0; k < h; k++ {
			Insertion(a[k:h])
			Insertion(a[h:])
		}
	}

}
