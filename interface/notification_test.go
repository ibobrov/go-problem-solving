package _interface

import (
	"errors"
	"testing"
)

type recordingSender struct {
	calls []Message
	err   error
}

var _ Sender = (*recordingSender)(nil)

func (s *recordingSender) Send(recipient, text string) error {
	s.calls = append(s.calls, Message{Recipient: recipient, Text: text})
	return s.err
}

func TestNewNotificationService(t *testing.T) {
	t.Parallel()

	var pointer *MemorySender
	tests := []struct {
		name    string
		sender  Sender
		wantErr bool
	}{
		{name: "valid sender", sender: &MemorySender{}},
		{name: "nil interface", wantErr: true},
		{name: "typed nil", sender: pointer},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service, err := NewNotificationService(tt.sender)
			if tt.wantErr {
				if service != nil || err == nil {
					t.Fatalf("got (%v, %v), want (nil, error)", service, err)
				}
				return
			}
			if err != nil || service == nil {
				t.Fatalf("got (%v, %v), want a service and no error", service, err)
			}
		})
	}
}

func TestNotify(t *testing.T) {
	t.Parallel()

	sendError := errors.New("delivery failed")
	tests := []struct {
		name    string
		message Message
		sendErr error
		invalid bool
	}{
		{name: "success", message: Message{Recipient: " Alice ", Text: " Привет "}},
		{name: "sender error", message: Message{Recipient: "Alice", Text: "Hi"}, sendErr: sendError},
		{name: "empty recipient", message: Message{Text: "Hi"}, invalid: true},
		{name: "empty text", message: Message{Recipient: "Alice"}, invalid: true},
		{name: "both empty", invalid: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sender := &recordingSender{err: tt.sendErr}
			service := newTestService(t, sender)
			err := service.Notify(tt.message.Recipient, tt.message.Text)

			if tt.invalid {
				if err == nil || len(sender.calls) != 0 {
					t.Fatalf("got error %v and calls %v, want a validation error and no calls", err, sender.calls)
				}
				return
			}
			if err != tt.sendErr {
				t.Errorf("got error %v, want the exact sender result %v", err, tt.sendErr)
			}
			if len(sender.calls) != 1 || sender.calls[0] != tt.message {
				t.Errorf("got calls %v, want exactly one call with %+v", sender.calls, tt.message)
			}
		})
	}
}

func TestMemorySender(t *testing.T) {
	t.Parallel()

	sender := &MemorySender{}
	service := newTestService(t, sender)
	messages := []Message{
		{Recipient: "Alice", Text: "First"},
		{Recipient: "Bob", Text: "Second"},
	}
	for _, message := range messages {
		if err := service.Notify(message.Recipient, message.Text); err != nil {
			t.Fatal(err)
		}
	}
	if len(sender.Messages) != len(messages) {
		t.Fatalf("stored %d messages, want %d", len(sender.Messages), len(messages))
	}
	for i, want := range messages {
		if got := sender.Messages[i]; got != want {
			t.Errorf("message %d = %+v, want %+v", i, got, want)
		}
	}
}

func newTestService(t *testing.T, sender Sender) *NotificationService {
	t.Helper()
	service, err := NewNotificationService(sender)
	if err != nil || service == nil {
		t.Fatalf("got (%v, %v), want a service and no error", service, err)
	}
	return service
}
