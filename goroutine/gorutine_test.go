package goroutine

import (
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	tests := []struct {
		name  string
		procs int
	}{
		{name: "один поток выполнения Go", procs: 1},
		{name: "два потока выполнения Go", procs: 2},
		{name: "пять потоков выполнения Go", procs: 5},
	}
	want := []string{"1", "2", "3", "4", "5"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// GOMAXPROCS и stdout общие для всего процесса.
			previousProcs := runtime.GOMAXPROCS(tt.procs)
			t.Cleanup(func() { runtime.GOMAXPROCS(previousProcs) })

			path := filepath.Join(t.TempDir(), "stdout.txt")
			output, err := os.Create(path)
			if err != nil {
				t.Fatal(err)
			}

			previousStdout := os.Stdout
			os.Stdout = output
			t.Cleanup(func() {
				os.Stdout = previousStdout
				output.Close()
			})

			printNumbers()

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			got := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Errorf("printNumbers() вывела %q, ожидались строки %q в любом порядке", data, want)
			}
		})
	}
}
