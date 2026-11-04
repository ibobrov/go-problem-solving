package slice

func RemoveValue(items []int, value int) []int {
	return filter(items, func(item int) bool {
		return item != value
	})
}

func KeepPositive(items []int) []int {
	return filter(items, func(item int) bool {
		return item > 0
	})
}

func filter(items []int, keep func(int) bool) []int {
	write := 0

	// Инвариант: перед каждой итерацией items[:write] содержит
	// сохранённые элементы уже просмотренной части в исходном порядке.
	for _, item := range items {
		if !keep(item) {
			continue
		}
		items[write] = item
		write++
	}

	// Обнуляем хвост исходного массива, не затрагивая items[:write].
	clear(items[write:])

	// Возвращаем укороченный срез того же массива.
	// Вызывающий код должен сохранить результат: items = filter(...).
	return items[:write]
}
