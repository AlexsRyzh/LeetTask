package string_task

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	number := make([]rune, 0, len(num)-k)
	number = append(number, rune(num[len(num)-1]))
	for i := len(num) - 2; i < len(num); i++ {
	}
	return "ds"
}
