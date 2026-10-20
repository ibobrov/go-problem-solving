package accamulator

import "testing"

func TestAccumulator(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "positive and negative",
			input: []int{2, -1, 5},
			want:  []int{2, 1, 6},
		},
		{
			name:  "zero",
			input: []int{0, 0},
			want:  []int{0, 0},
		},
		{
			name:  "negative values",
			input: []int{-5, -3},
			want:  []int{-5, -8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Accumulator()

			for i, value := range tt.input {
				got := acc(value)

				if got != tt.want[i] {
					t.Errorf(
						"call %d: got %d, want %d",
						i, got, tt.want[i],
					)
				}
			}
		})
	}
}

func TestAccumulator_Independent(t *testing.T) {
	a := Accumulator()
	b := Accumulator()

	if got := a(10); got != 10 {
		t.Errorf("a(10) = %d, want 10", got)
	}

	if got := b(5); got != 5 {
		t.Errorf("b(5) = %d, want 5", got)
	}

	if got := a(2); got != 12 {
		t.Errorf("a(2) = %d, want 12", got)
	}

	if got := b(3); got != 8 {
		t.Errorf("b(3) = %d, want 8", got)
	}
}

func TestReaders(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "two values",
			values: []int{10, 20},
			want:   []int{10, 20},
		},
		{
			name:   "duplicate values",
			values: []int{5, 5},
			want:   []int{5, 5},
		},
		{
			name:   "nil slice",
			values: nil,
			want:   []int{},
		},
		{
			name:   "empty slice",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readers := Readers(tt.values)

			if len(readers) != len(tt.want) {
				t.Fatalf(
					"len = %d, want %d",
					len(readers), len(tt.want),
				)
			}

			for i, reader := range readers {
				if got := reader(); got != tt.want[i] {
					t.Errorf(
						"reader[%d]() = %d, want %d",
						i, got, tt.want[i],
					)
				}
			}
		})
	}
}

func TestReaders_InputMutation(t *testing.T) {
	values := []int{10, 20}
	readers := Readers(values)

	values[0] = 100
	values[1] = 200

	if got := readers[0](); got != 10 {
		t.Errorf("reader[0]() = %d, want 10", got)
	}

	if got := readers[1](); got != 20 {
		t.Errorf("reader[1]() = %d, want 20", got)
	}
}
