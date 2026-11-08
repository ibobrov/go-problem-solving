package _interface

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "invalid field: " + e.Field
}

func ValidateName(name string) error {
	if name == "" {
		return &ValidationError{Field: "name"}
	}
	return nil
}
