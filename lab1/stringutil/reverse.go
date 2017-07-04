package stringutil

// package stringutil terkandung utility fungsi untuk bekerja dengan string
// Reverse returns its argument string Reversed run-wise left to right

func Reverse(s string) string { // Reverse adalah public fungsi
	return reverseTwo(s) // reverseTwo adalah private fungsi jadi hanya tersedia didalam package
}

func reverseTwo(s string) string { // reverseTwo adalah private fungsi jadi hanya tersedia didalam package
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// this demonstrates how an unexported function
// can be used by an exported function in the same package
