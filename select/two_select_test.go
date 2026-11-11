package _select

import (
	"slices"
	"strings"
	"testing"
)

func TestPrintAll(t *testing.T) {
	t.Parallel()

	closedChannel := func(values ...int) <-chan int {
		ch := make(chan int, len(values))
		for _, value := range values {
			ch <- value
		}
		close(ch)
		return ch
	}

	tests := []struct {
		name   string
		first  <-chan int
		second <-chan int
		want   []string
	}{
		{
			name:   "значения в обоих каналах",
			first:  closedChannel(0, 1, 1),
			second: closedChannel(2, 3),
			want:   []string{"first: 0", "first: 1", "first: 1", "second: 2", "second: 3"},
		},
		{
			name:   "первый канал закрыт без значений",
			first:  closedChannel(),
			second: closedChannel(2, 3),
			want:   []string{"second: 2", "second: 3"},
		},
		{
			name:   "второй канал закрыт без значений",
			first:  closedChannel(1, 2),
			second: closedChannel(),
			want:   []string{"first: 1", "first: 2"},
		},
		{name: "оба канала закрыты без значений", first: closedChannel(), second: closedChannel()},
		{name: "первый канал nil", second: closedChannel(2), want: []string{"second: 2"}},
		{name: "второй канал nil", first: closedChannel(1), want: []string{"first: 1"}},
		{name: "оба канала nil"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var output strings.Builder
			printAllTo(&output, tt.first, tt.second)

			got := slices.Collect(strings.Lines(output.String()))
			for i, line := range got {
				got[i] = strings.TrimSuffix(line, "\n")
			}
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("printAll() вывела %q, ожидались строки %q в любом порядке", output.String(), tt.want)
			}
		})
	}
}
