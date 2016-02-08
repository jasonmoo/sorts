package nab

type (
	Hardware struct {
		Type string
		Size int
	}
)

func Sort(a []*Hardware) {

	i, j, mid := 0, len(a)-1, len(a)/2

	// presort
	for {
		for a[i].Type == "nut" {
			i++
		}
		for a[j].Type == "bolt" {
			j--
		}
		if i >= mid {
			break
		}
		a[i], a[j] = a[j], a[i]
	}

	// pair
	for j >= 0 {
		ii := i
		for a[j].Size != a[ii].Size {
			ii++
		}
		if ii != i {
			a[ii], a[i] = a[i], a[ii]
		}
		i++
		j--
	}

}
