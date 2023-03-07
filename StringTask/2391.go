package string_task

func GarbageCollection(garbage []string, travel []int) int {

	newGarbage := make([]map[rune]int, len(garbage))

	for i := 0; i < len(garbage); i++ {
		for _, v := range garbage[i] {
			if newGarbage[i] == nil {
				newGarbage[i] = make(map[rune]int)
			}
			newGarbage[i][v]++
		}
	}

	total := 0
	for _, v := range "MPG" {
		garbageCount := 0
		totalCount := 0
		travelCount := 0

		val, ok := newGarbage[0][v]
		if ok {
			garbageCount = val
		}
		for i := 1; i < len(newGarbage); i++ {
			val, ok = newGarbage[i][v]
			totalCount += travel[i-1]
			if ok {
				garbageCount += val
				travelCount = totalCount
			}
		}
		total += garbageCount + travelCount
	}
	return total
}
