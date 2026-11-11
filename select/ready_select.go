package _select

import (
	"fmt"
	"io"
	"os"
)

func main() {
	printReady(os.Stdout, 10)
}

func printReady(output io.Writer, iterations int) {
	ch := make(chan int, 1)

	for i := 0; i < iterations; i++ {
		select {
		case value := <-ch:
			fmt.Fprint(output, value, " ")
		case ch <- i:
		}
	}
}

// На чётных итерациях буфер пуст, поэтому готова только отправка.
// На нечётных буфер полон, поэтому готово только чтение.
// За десять итераций выведет "0 2 4 6 8 ".
// С небуферизованным каналом обе ветки заблокируются на первой итерации:
// другой горутины для обмена нет, поэтому Go обнаружит дедлок.
