package string_task

import (
	"strconv"
)

func MinPartitions(n string) (result int) {
	maxCh := rune(n[0])
	for _, v := range n {
		if maxCh < v {
			maxCh = v
		}
	}
	result, _ = strconv.Atoi(string(maxCh))
	return
}
