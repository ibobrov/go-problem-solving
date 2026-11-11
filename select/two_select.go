package _select

import (
	"fmt"
	"io"
	"os"
)

func printAll(first, second <-chan int) {
	printAllTo(os.Stdout, first, second)
}

func printAllTo(output io.Writer, first, second <-chan int) {
	// Закрытый канал отключаем через nil и продолжаем читать через select.
	for first != nil || second != nil {
		select {
		case value, ok := <-first:
			if !ok {
				first = nil
				continue
			}
			fmt.Fprintln(output, "first:", value)
		case value, ok := <-second:
			if !ok {
				second = nil
				continue
			}
			fmt.Fprintln(output, "second:", value)
		}
	}
}
