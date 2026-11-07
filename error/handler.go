package error

import "fmt"

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return "invalid " + e.Field + ": " + e.Reason
}

func handle() error {
	return validateEmail("")
}

func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Reason: "empty"}
	}
	return nil
}

func registerUser(email string) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("register user: %w", err)
	}
	return nil
}
