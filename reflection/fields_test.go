package reflection

import (
	"reflect"
	"testing"
)

func TestNonZeroFields(t *testing.T) {
	t.Parallel()

	type profile struct {
		Name   string `json:"name"`
		Age    int    `json:"age,omitempty"`
		Secret string `json:"-"`
		note   string
	}
	type Embedded struct{ Count int }
	alice := profile{Name: "Alice", Secret: "token", note: "private"}
	wantAlice := map[string]any{"name": "Alice"}
	tests := []struct {
		name  string
		input any
		want  map[string]any
	}{
		{"struct", alice, wantAlice},
		{"pointer", &alice, wantAlice},
		{"all zero", profile{}, map[string]any{}},
		{"unexported", profile{note: "private"}, map[string]any{}},
		{"empty tag name and non-nil empty slice",
			struct {
				Items []int `json:",omitempty"`
			}{Items: []int{}},
			map[string]any{"Items": []int{}},
		},
		{"embedded field", struct{ Embedded }{Embedded{Count: 1}},
			map[string]any{"Embedded": Embedded{Count: 1}}},
		{"dash with options is a name",
			struct {
				Value int `json:"-,omitempty"`
			}{Value: 1},
			map[string]any{"-": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NonZeroFields(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestNonZeroFieldsErrors(t *testing.T) {
	t.Parallel()

	type conflicting struct {
		Name  string
		Other string `json:"Name"`
	}
	value := 42
	ptr := &struct{}{}
	tests := []struct {
		name  string
		input any
	}{
		{"nil", nil},
		{"nil pointer", (*struct{})(nil)},
		{"pointer to pointer", &ptr},
		{"wrong type", value},
		{"pointer to wrong type", &value},
		{"duplicate name with zero fields", conflicting{}},
		{"duplicate name with nonzero field", conflicting{Other: "value"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, err := NonZeroFields(tt.input); err == nil {
				t.Errorf("got %#v, want an error", got)
			}
		})
	}
}
