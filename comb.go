package sorts

func Comb(a []int) {

	const shrink = 1.3

	gap := float64(len(a))
	swapped := true

	for gap > 1 && swapped {

		gap /= shrink

		// trim and floor the float
		g := int(gap)
		if g < 1 {
			g = 1
		}
		gap = float64(g)

		swapped = false

		for i := 0; i+g < len(a); i++ {
			if a[i+g] < a[i] {
				a[i], a[i+g] = a[i+g], a[i]
				swapped = true
			}
		}

	}

}
