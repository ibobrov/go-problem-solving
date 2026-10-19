package temperature

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestParseCelsius(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int16
		wantErr error
	}{
		{
			name:  "zero",
			input: "0",
			want:  0,
		},
		{
			name:  "max int16",
			input: "32767",
			want:  32767,
		},
		{
			name:  "min int16",
			input: "-32768",
			want:  -32768,
		},
		{
			name:  "negative",
			input: "-25",
			want:  -25,
		},
		{
			name:    "positive overflow",
			input:   "32768",
			wantErr: strconv.ErrRange,
		},
		{
			name:    "negative overflow",
			input:   "-32769",
			wantErr: strconv.ErrRange,
		},
		{
			name:    "invalid string",
			input:   "abc",
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCelsius(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf(
					"ParseCelsius(%q) error = %v, want %v",
					tt.input, err, tt.wantErr,
				)
			}

			if got != tt.want {
				t.Errorf(
					"ParseCelsius(%q) = %d, want %d",
					tt.input, got, tt.want,
				)
			}
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	tests := []struct {
		name  string
		input int16
		want  float64
	}{
		{
			name:  "zero",
			input: 0,
			want:  32,
		},
		{
			name:  "positive",
			input: 25,
			want:  77,
		},
		{
			name:  "negative",
			input: -10,
			want:  14,
		},
		{
			name:  "fractional result",
			input: 21,
			want:  69.8,
		},
		{
			name:  "min int16",
			input: -32768,
			want:  -58950.4,
		},
		{
			name:  "max int16",
			input: 32767,
			want:  59012.6,
		},
	}

	const epsilon = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToFahrenheit(tt.input)

			if math.Abs(got-tt.want) > epsilon {
				t.Errorf(
					"ToFahrenheit(%d) = %f, want %f",
					tt.input, got, tt.want,
				)
			}
		})
	}
}
