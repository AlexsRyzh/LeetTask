package string_task

func partitionLabels(s string) []int {
	last_index_char := map[byte]int{}
	for i := 0; i < len(s); i++ {
		last_index_char[s[i]] = i
	}
	start, end := 0, 0
	answers := []int{}
	for i := 0; i < len(s); i++ {
		v := last_index_char[s[i]]
		if v > end {
			end = v
		}
		if i == end {
			answers = append(answers, end-start+1)
			start = end + 1
		}
	}
	return answers
}
