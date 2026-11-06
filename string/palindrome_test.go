package string

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", s: "", want: true},
		{name: "single ASCII", s: "a", want: true},
		{name: "single Cyrillic", s: "я", want: true},
		{name: "single emoji", s: "😀", want: true},

		{name: "odd ASCII", s: "level", want: true},
		{name: "even ASCII", s: "levvel", want: true},
		{name: "not palindrome", s: "hello", want: false},
		{name: "Go", s: "Go", want: false},
		{name: "inner mismatch", s: "abca", want: false},
		{name: "Cyrillic", s: "топот", want: true},
		{name: "even Cyrillic", s: "анна", want: true},
		{name: "emoji palindrome", s: "😀a😀", want: true},

		{name: "ASCII case matters", s: "Aa", want: false},
		{name: "Cyrillic case matters", s: "Аа", want: false},
		{name: "spaces matter", s: "а роза упала на лапу азора", want: false},
		{name: "leading space matters", s: " aa", want: false},
		{name: "symmetric spaces", s: "а а", want: true},
		{name: "punctuation matters", s: "а,а!", want: false},
		{name: "symmetric punctuation", s: "!а!", want: true},

		// Две одинаковые графемы é: [e, ◌́, e, ◌́].
		// Визуальная симметрия не означает палиндром по кодовым точкам.
		{name: "combining marks", s: "e\u0301e\u0301", want: false},
		// [e, ◌́, e] — палиндром именно по кодовым точкам.
		{name: "combining mark in center", s: "e\u0301e", want: true},
		// é и e + ◌́ не приравниваются: нормализация не выполняется.
		{name: "no normalization", s: "\u00e9e\u0301", want: false},

		{name: "invalid byte", s: "\xff", want: false},
		{name: "invalid byte in center", s: "a\xffa", want: false},
		{name: "truncated UTF-8", s: "\xe2\x82", want: false},
		{name: "unexpected continuation", s: "\x80", want: false},
		{name: "overlong encoding", s: "\xc0\xaf", want: false},
		{name: "encoded surrogate", s: "\xed\xa0\x80", want: false},
		{name: "above Unicode maximum", s: "\xf4\x90\x80\x80", want: false},

		// Валидный U+FFFD нельзя путать с ошибкой декодирования UTF-8.
		{name: "valid replacement character", s: "\uFFFD", want: true},
		{name: "replacement character palindrome", s: "\uFFFDa\uFFFD", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %t; want %t",
					tt.s, got, tt.want)
			}
		})
	}
}
