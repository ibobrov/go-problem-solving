package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	t.Parallel()

	err, ok := handle().(*ValidationError)
	if !ok || err == nil || err.Field != "email" || err.Reason != "empty" {
		t.Fatalf("handle() = %#v, want a validation error for an empty email", err)
	}
	if got, want := err.Error(), "invalid email: empty"; got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestValidationErrorChain(t *testing.T) {
	t.Parallel()

	original := &ValidationError{Field: "email", Reason: "empty"}
	wrapped := fmt.Errorf("register user: %w", original)
	other := &ValidationError{Field: original.Field, Reason: original.Reason}

	for _, tt := range []struct {
		name   string
		target error
		want   bool
	}{
		{name: "original instance", target: original, want: true},
		{name: "same fields in another instance", target: other},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := errors.Is(wrapped, tt.target); got != tt.want {
				t.Errorf("errors.Is() = %t, want %t", got, tt.want)
			}
		})
	}

	var target *ValidationError
	if !errors.As(wrapped, &target) || target != original {
		t.Fatalf("errors.As() found %#v, want the original error", target)
	}
	if target.Field != "email" || target.Reason != "empty" {
		t.Errorf("unexpected validation details: %+v", target)
	}
}

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		email   string
		wantErr bool
	}{
		{name: "nonempty email", email: "alice@example.com"},
		{name: "empty email", wantErr: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := registerUser(tt.email)
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("registerUser() = %v, want nil", err)
				}
				return
			}

			cause, ok := errors.Unwrap(err).(*ValidationError)
			if !ok || cause == nil || cause.Field != "email" || cause.Reason != "empty" {
				t.Fatalf("registerUser() = %v, want a wrapped validation error for an empty email", err)
			}
			if got, want := err.Error(), "register user: invalid email: empty"; got != want {
				t.Errorf("Error() = %q, want %q", got, want)
			}
		})
	}
}
