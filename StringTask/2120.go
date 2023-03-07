package string_task

import "fmt"

func move(newX int, newY int, ch rune) (int, int) {
	switch ch {
	case 'R':
		newX++
	case 'L':
		newX--
	case 'U':
		newY++
	case 'D':
		newY--
	}
	return newX, newY
}

func ExecuteInstructions(n int, startPos []int, s string) []int {
	y, x := startPos[0], startPos[1]
	n--
	answer := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		X, Y := x, y
		for j := i; j < len(s); j++ {
			newX, newY := move(X, Y, rune(s[j]))
			if newX > n || newX < 0 ||
				newY > n || newY < 0 {
				fmt.Println(newX, newY)
				answer[i] = j
				break
			}
			X, Y = newX, newY
		}
	}
	return answer
}
