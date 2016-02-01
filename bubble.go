package sorts

func Bubble(a []int) {

	for {
		var swapped bool
		for i := 1; i < len(a); i++ {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

}
