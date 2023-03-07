package string_task

func letterCasePermutation(s string) []string {
	variant := []byte(s)
	for i := 0; i < len(variant); i++ {
		if 'A' <= variant[i] && variant[i] <= 'Z' {
			variant[i] += ('a' - 'A')
		}
	}

	answer := [][]byte{variant}
	for i := 0; i < len(variant); i++ {
		tempArr := [][]byte{}
		for j := 0; 'a' <= variant[i] && variant[i] <= 'z' && j < len(answer); j++ {
			word := make([]byte, len(variant))
			copy(word, answer[j])
			word[i] -= ('a' - 'A')
			tempArr = append(tempArr, word)
		}
		answer = append(answer, tempArr...)
	}
	answerStr := make([]string, 0, len(answer))
	for i := 0; i < len(answer); i++ {
		answerStr = append(answerStr, string(answer[i]))
	}

	return answerStr
}
