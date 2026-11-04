package slice

import (
	"slices"
	"testing"
)

func TestRemoveValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []int
		value int
		want  []int
	}{
		{
			name:  "nil",
			items: nil,
			value: 1,
			want:  nil,
		},
		{
			name:  "empty",
			items: []int{},
			value: 1,
			want:  []int{},
		},
		{
			name:  "value absent",
			items: []int{1, 2, 3},
			value: 5,
			want:  []int{1, 2, 3},
		},
		{
			name:  "repeated value",
			items: []int{1, 2, 3, 2, 4},
			value: 2,
			want:  []int{1, 3, 4},
		},
		{
			name:  "remove all",
			items: []int{2, 2, 2},
			value: 2,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			original := tt.items
			got := RemoveValue(tt.items, tt.value)

			checkFilterResult(t, original, got, tt.want)
		})
	}
}

func TestKeepPositive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{
			name:  "nil",
			items: nil,
			want:  nil,
		},
		{
			name:  "empty",
			items: []int{},
			want:  []int{},
		},
		{
			name:  "mixed values",
			items: []int{-1, 3, 0, 2},
			want:  []int{3, 2},
		},
		{
			name:  "remove all",
			items: []int{-3, 0, -1},
			want:  []int{},
		},
		{
			name:  "keep all",
			items: []int{1, 3, 2},
			want:  []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			original := tt.items
			got := KeepPositive(tt.items)

			checkFilterResult(t, original, got, tt.want)
		})
	}
}

func checkFilterResult(t *testing.T, original, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Fatalf("result = %v, want %v", got, want)
	}

	if (got == nil) != (want == nil) {
		t.Errorf("result nil = %t, want %t", got == nil, want == nil)
	}

	// Проверяем очистку хвоста через исходный срез.
	for i := len(got); i < len(original); i++ {
		if original[i] != 0 {
			t.Errorf("tail[%d] = %d, want 0", i, original[i])
		}
	}

	// Проверяем переиспользование массива, включая удаление всех элементов.
	if len(original) > 0 {
		if cap(got) == 0 {
			t.Fatal("result lost the original backing array")
		}

		if &got[:1][0] != &original[0] {
			t.Error("result uses a different backing array")
		}
	}
}
