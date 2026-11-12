package context

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	printCancellation(os.Stdout)
}

func printCancellation(output io.Writer) {
	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()
	first, cancelFirst := context.WithTimeout(parent, time.Minute)
	defer cancelFirst()
	second, cancelSecond := context.WithCancel(parent)
	defer cancelSecond()

	// Явная отмена первого потомка оставляет родителя активным.
	cancelFirst()
	fmt.Fprintln(output, first.Err())
	fmt.Fprintln(output, parent.Err())

	// Отмена родителя распространяется на второго потомка.
	cancelParent()
	<-second.Done()
	fmt.Fprintln(output, second.Err())

	cancelSecond()
}
