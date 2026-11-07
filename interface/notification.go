package _interface

import "errors"

type Sender interface {
	Send(recipient, text string) error
}

type NotificationService struct {
	sender Sender
}

func NewNotificationService(sender Sender) (*NotificationService, error) {
	if sender == nil {
		return nil, errors.New("sender is nil")
	}

	return &NotificationService{
		sender: sender,
	}, nil
}

func (s *NotificationService) Notify(recipient, text string) error {
	if recipient == "" {
		return errors.New("recipient is empty")
	}
	if text == "" {
		return errors.New("text is empty")
	}
	return s.sender.Send(recipient, text)
}

type Message struct {
	Recipient string
	Text      string
}

type MemorySender struct {
	Messages []Message
}

var _ Sender = (*MemorySender)(nil)

func (s *MemorySender) Send(recipient, text string) error {
	s.Messages = append(s.Messages, Message{Recipient: recipient, Text: text})
	return nil
}
