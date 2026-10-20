package accamulator

func Accumulator() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func Readers(values []int) []func() int {
	result := make([]func() int, len(values))
	for i, value := range values {
		result[i] = func() int { return value }
	}
	return result
}
