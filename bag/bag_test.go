package bag

import "testing"

func TestBag(t *testing.T) {
	tests := []struct {
		name         string
		values       []string
		check        string
		wantContains bool
		wantLen      int
	}{
		{
			name:         "empty bag",
			values:       nil,
			check:        "Go",
			wantContains: false,
			wantLen:      0,
		},
		{
			name:         "first addition",
			values:       []string{"Go"},
			check:        "Go",
			wantContains: true,
			wantLen:      1,
		},
		{
			name:         "duplicate addition",
			values:       []string{"Go", "Go"},
			check:        "Go",
			wantContains: true,
			wantLen:      1,
		},
		{
			name:         "empty string",
			values:       []string{""},
			check:        "",
			wantContains: true,
			wantLen:      1,
		},
		{
			name:         "multiple values",
			values:       []string{"Go", "Rust", "Java"},
			check:        "Rust",
			wantContains: true,
			wantLen:      3,
		},
		{
			name:         "missing value",
			values:       []string{"Go", "Rust"},
			check:        "Python",
			wantContains: false,
			wantLen:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bag

			for _, value := range tt.values {
				b.Add(value)
			}

			if got := b.Contains(tt.check); got != tt.wantContains {
				t.Errorf(
					"Contains(%q) = %v, want %v",
					tt.check, got, tt.wantContains,
				)
			}

			if got := b.Len(); got != tt.wantLen {
				t.Errorf(
					"Len() = %d, want %d",
					got, tt.wantLen,
				)
			}
		})
	}
}
