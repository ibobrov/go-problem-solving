package slice

import "testing"

func TestCompress(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "пустая строка", input: "", want: ""},
		{name: "один символ", input: "a", want: "a"},
		{name: "без повторов", input: "abc", want: "abc"},
		{name: "одна группа", input: "aaaa", want: "a4"},
		{name: "две группы", input: "aaabbb", want: "a3b3"},
		{name: "повтор в начале", input: "aabc", want: "a2bc"},
		{name: "повтор в середине", input: "abbc", want: "ab2c"},
		{name: "повтор в конце", input: "abcc", want: "abc2"},
		{name: "раздельные группы", input: "aaabaa", want: "a3ba2"},
		{name: "чередование", input: "ababab", want: "ababab"},
		{name: "разный регистр", input: "AAaa", want: "A2a2"},
		{name: "русские буквы", input: "яяяюю", want: "я3ю2"},
		{name: "эмодзи", input: "😀😀😀🙃", want: "😀3🙃"},
		{name: "двузначное количество", input: "aaaaaaaaaaaa", want: "a12"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compress(tt.input)

			if got != tt.want {
				t.Errorf("compress(%q) = %q; want %q",
					tt.input, got, tt.want)
			}
		})
	}
}
