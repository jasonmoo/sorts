package sorts

func Selection(a []int) {

	for i := 1; i < len(a); i++ {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[k] {
				k = j
			}
		}
		if k != i {
			a[k], a[i] = a[i], a[k]
		}
	}

}
