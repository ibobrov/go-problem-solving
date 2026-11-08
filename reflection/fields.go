package reflection

import (
	"fmt"
	"reflect"
	"strings"
)

func NonZeroFields(v any) (map[string]any, error) {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return nil, fmt.Errorf("expected a struct or a pointer to a struct, got nil")
	}

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil, fmt.Errorf("expected a non-nil pointer, got %T", v)
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct or a pointer to a struct, got %T", v)
	}

	typ := value.Type()
	result := make(map[string]any)
	names := make(map[string]struct{})
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if name == "" {
			name = field.Name
		}

		if _, exists := names[name]; exists {
			return nil, fmt.Errorf("duplicate field name %q", name)
		}
		names[name] = struct{}{}

		item := value.Field(i)
		if !item.IsZero() {
			result[name] = item.Interface()
		}
	}
	return result, nil
}
