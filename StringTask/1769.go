package string_task



func MinOperations(boxes string) []int {
	n := len(boxes)
	right := make([]int, n)
	left := make([]int, n)

	var beforeN int

	if boxes[0] == '1' {
		beforeN = 1
	}

	for i := 1; i < n; i++ {
		right[i] = right[i-1] + beforeN
		if boxes[i] == '1' {
			beforeN++
		}
	}

	beforeN = 0
	if boxes[n-1] == '1' {
		beforeN = 1
	}

	for i := n - 2; i > -1; i-- {
		left[i] = left[i+1] + beforeN
		if boxes[i] == '1' {
			beforeN++
		}
	}

	for i := 0; i < n; i++ {
		right[i] += left[i]
	}

	return right

}
