package _interface

import (
	"errors"
	"testing"
)

func TestTypedNilValidationError(t *testing.T) {
	t.Parallel()

	var pointer *ValidationError
	var err error = pointer
	if err == nil {
		t.Fatal("expected an error interface with a dynamic type")
	}

	var actual *ValidationError
	if !errors.As(err, &actual) || actual != nil {
		t.Fatal("expected the interface to hold a nil *ValidationError")
	}
}

func TestValidateNameSuccess(t *testing.T) {
	t.Parallel()

	for _, name := range []string{"Alice", " "} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if err := ValidateName(name); err != nil {
				t.Fatalf("ValidateName(%q) returned a non-nil error interface of type %T", name, err)
			}
		})
	}
}

func TestValidateNameEmpty(t *testing.T) {
	t.Parallel()

	err := ValidateName("")
	if err == nil {
		t.Fatal("expected an error for an empty name")
	}

	var validationErr *ValidationError
	if !errors.As(err, &validationErr) || validationErr == nil {
		t.Fatalf("expected a non-nil *ValidationError, got %T", err)
	}
	if validationErr.Field != "name" {
		t.Errorf("got field %q, want %q", validationErr.Field, "name")
	}
}
