package _map

import (
	"maps"
	"testing"
)

func TestFrequencies(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{
			name:  "nil",
			words: nil,
			want:  map[string]int{},
		},
		{
			name:  "empty",
			words: []string{},
			want:  map[string]int{},
		},
		{
			name:  "unique words",
			words: []string{"go", "java", "rust"},
			want:  map[string]int{"go": 1, "java": 1, "rust": 1},
		},
		{
			name:  "repeated words",
			words: []string{"go", "java", "go", "rust", "java", "go"},
			want:  map[string]int{"go": 3, "java": 2, "rust": 1},
		},
		{
			name:  "empty word",
			words: []string{"", "go", ""},
			want:  map[string]int{"": 2, "go": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Frequencies(tt.words)

			if !maps.Equal(got, tt.want) {
				t.Errorf("Frequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMostFrequent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		freq      map[string]int
		wantWord  string
		wantCount int
		wantOK    bool
	}{
		{
			name: "nil",
			freq: nil,
		},
		{
			name: "empty",
			freq: map[string]int{},
		},
		{
			name:      "single word",
			freq:      map[string]int{"go": 3},
			wantWord:  "go",
			wantCount: 3,
			wantOK:    true,
		},
		{
			name:      "unique maximum",
			freq:      map[string]int{"go": 3, "java": 2, "rust": 1},
			wantWord:  "go",
			wantCount: 3,
			wantOK:    true,
		},
		{
			name:      "equal maximum",
			freq:      map[string]int{"rust": 3, "go": 3, "java": 1},
			wantWord:  "go",
			wantCount: 3,
			wantOK:    true,
		},
		{
			name:      "all frequencies equal",
			freq:      map[string]int{"rust": 2, "java": 2, "go": 2},
			wantWord:  "go",
			wantCount: 2,
			wantOK:    true,
		},
		{
			name:      "empty word wins tie",
			freq:      map[string]int{"go": 2, "": 2},
			wantWord:  "",
			wantCount: 2,
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			before := maps.Clone(tt.freq)

			// Повторяем вызов, чтобы повысить вероятность обнаружения
			// зависимости результата от порядка обхода map.
			for attempt := 0; attempt < 100; attempt++ {
				word, count, ok := MostFrequent(tt.freq)

				if word != tt.wantWord ||
					count != tt.wantCount ||
					ok != tt.wantOK {
					t.Fatalf(
						"MostFrequent() = (%q, %d, %t), want (%q, %d, %t)",
						word, count, ok,
						tt.wantWord, tt.wantCount, tt.wantOK,
					)
				}

				if !maps.Equal(tt.freq, before) {
					t.Fatal("MostFrequent() changed the input map")
				}
			}
		})
	}
}
