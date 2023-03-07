package string_task

func MinSteps(s string, t string) int {

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
		count[int(t[i]-'a')]--
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			ans += count[i]
		}
	}
	return ans
}
