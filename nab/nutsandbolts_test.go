package nab

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {

	var a []*Hardware
	for i := 0; i < 10; i++ {
		a = append(a, &Hardware{Type: "nut", Size: i})
	}
	for i := 0; i < 10; i++ {
		a = append(a, &Hardware{Type: "bolt", Size: i})
	}

	for i := 0; i < 10000; i++ {
		i, j := rand.Intn(len(a)), rand.Intn(len(a))
		if i != j {
			a[i], a[j] = a[j], a[i]
		}
	}

	for _, h := range a {
		fmt.Println(h.Type, h.Size)
	}

	Sort(a)
	fmt.Println()

	for _, h := range a {
		fmt.Println(h.Type, h.Size)
	}

}
