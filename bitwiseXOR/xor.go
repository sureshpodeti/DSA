package bitwiseXOR

/*
ar := []int{1, 5, 2, 6, 4}

	mis := missing(ar)

	fmt.Println(mis)
*/
func Missing(ar []int) int {
	n := len(ar)
	x := 1
	for i := 2; i <= n+1; i++ {
		x ^= i
	}

	y := ar[0]
	for i := 2; i < n; i++ {
		y ^= i
	}

	return x ^ y
}
