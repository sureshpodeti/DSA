package bitwiseXOR

/*
//ar := []int{1, 4, 2, 1, 3, 2, 3}

	ar := []int{7, 9, 7}

	num := FindSingle(ar)

	fmt.Printf("Single num: %d\n", num)
*/

func FindSingle(ar []int) int {
	x := ar[0]

	for i := 1; i < len(ar); i++ {
		x ^= ar[i]
	}
	return x
}
