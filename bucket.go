package sorts

func Bucket(a []int, n int) {

	buckets := make([][]int, n)
	chunk := len(a) / n

	for _, v := range a {

		i := v / chunk

		if i > 0 {
			i--
		}

		buckets[i] = append(buckets[i], v)
	}

	a = a[:0]

	for _, bucket := range buckets {
		Insertion(bucket)
		a = append(a, bucket...)
	}

}

func BucketInsertion(a []int, n int) {

	buckets := make([][]int, n)
	chunk := len(a) / n

	for _, v := range a {

		i := v / chunk

		if i > 0 {
			i--
		}

		buckets[i] = append(buckets[i], v)
	}

	a = a[:0]

	for _, bucket := range buckets {
		a = append(a, bucket...)
	}

	Insertion(a)

}
