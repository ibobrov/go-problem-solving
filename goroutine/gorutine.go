package goroutine

import (
	"fmt"
	"sync"
)

func printNumbers() {
	var wg sync.WaitGroup

	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		wg.Add(1)

		// каждая горутина независимая и порядок заранее не определен
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(val)
	}

	// sleep может ждать дольше чем надо или завершиться слишком рано
	wg.Wait()
}
