package find

func findDuplicatesInColumn(i int, input []string) [][2]int {
	buf := make(map[string][][2]int)
	for index, value := range input {
		buf[value] = append(buf[value], [2]int{i, index})
	}

	var res [][2]int
	for _, v := range buf {
		if len(v) > 1 {
			res = append(res, v...)
		}
	}

	return res
}

// todo
func findMinLen(input [][]string) int {
	l := len(input[0])
	for _, strings := range input {
		if l > len(strings) {
			l = len(strings)
		}
	}

	return l
}

func Duplicates(input [][]string) [][2]int {

	minLen := findMinLen(input)
	var result [][2]int
	for i := 0; i < minLen; i++ {

		var column []string
		for _, i3 := range input {
			column = append(column, i3[i])
		}

		duplicatesInColumn := findDuplicatesInColumn(i, column)
		result = append(result, duplicatesInColumn...)

	}

	return result

}
