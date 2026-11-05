package slice

import (
	"strconv"
	"strings"
)

func Compress(str string) string {
	runes := []rune(str)
	if len(runes) <= 1 {
		return str
	}

	var result strings.Builder
	result.Grow(len(str))

	writeGroup := func(char rune, count int) {
		result.WriteRune(char)
		if count > 1 {
			result.WriteString(strconv.Itoa(count))
		}
	}

	groupStart := 0
	count := 1

	for read := 1; read < len(runes); read++ {
		if runes[read] == runes[groupStart] {
			count++
			continue
		}

		writeGroup(runes[groupStart], count)
		groupStart = read
		count = 1
	}

	writeGroup(runes[groupStart], count)
	return result.String()
}
