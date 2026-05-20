package arrays

func findThePrefixCommonArray(A []int, B []int) []int {
	C := make([]int, len(A))
	if A[0] == B[0] {
		C[0] = 1
	}

	for i := 1; i < len(A); i++ {
		C[i] += C[i-1]
		if A[i] == B[i] {
			C[i] += 1
			continue
		}

		for j := range i {
			if B[j] == A[i] {
				C[i] += 1
			}
			if A[j] == B[i] {
				C[i] += 1
			}
		}
	}
	return C
}
