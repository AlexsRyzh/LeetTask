package string_task

func MinAddToMakeValid(s string) int {
	count := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
			continue
		}
		count--

		if count == -1 {
			ans++
			count++
		}
	}
	return count + ans
}
