package _select

import (
	"strings"
	"testing"
)

func TestPrintReady(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		iterations int
		want       string
	}{
		{name: "без итераций", iterations: 0, want: ""},
		{name: "первая отправка", iterations: 1, want: ""},
		{name: "первое чтение", iterations: 2, want: "0 "},
		{name: "последнее значение остаётся в буфере", iterations: 5, want: "0 2 "},
		{name: "десять итераций из задания", iterations: 10, want: "0 2 4 6 8 "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var output strings.Builder
			printReady(&output, tt.iterations)

			if got := output.String(); got != tt.want {
				t.Errorf("printReady(%d) вывела %q, ожидалось %q", tt.iterations, got, tt.want)
			}
		})
	}
}
