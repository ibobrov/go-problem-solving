package _select

import (
	"io"
	"testing"
	"testing/synctest"
	"time"
)

func TestReceive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		messages      []string
		closed        bool
		timeout       time.Duration
		want          string
		wantErr       string
		wantEOF       bool
		wantRemaining int
	}{
		{name: "готовое сообщение", messages: []string{"привет"}, timeout: time.Second, want: "привет"},
		{name: "пустое сообщение", messages: []string{""}, timeout: time.Second},
		{name: "сообщение в закрытом канале", messages: []string{"последнее"}, closed: true, timeout: time.Second, want: "последнее"},
		{name: "закрытый пустой канал", closed: true, timeout: time.Second, wantEOF: true},
		{name: "тайм-аут на открытом канале", messages: []string{}, timeout: time.Second, wantErr: "timeout"},
		{name: "тайм-аут на nil-канале", timeout: time.Second, wantErr: "timeout"},
		{name: "нулевой тайм-аут", messages: []string{"привет"}, wantErr: "incorrect timeout", wantRemaining: 1},
		{name: "отрицательный тайм-аут", messages: []string{"привет"}, timeout: -time.Second, wantErr: "incorrect timeout", wantRemaining: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				var messages chan string
				if tt.messages != nil || tt.closed {
					messages = make(chan string, len(tt.messages))
					for _, message := range tt.messages {
						messages <- message
					}
					if tt.closed {
						close(messages)
					}
				}

				got, err := receive(messages, tt.timeout)
				if got != tt.want {
					t.Errorf("receive() вернула %q, ожидалось %q", got, tt.want)
				}
				switch {
				case tt.wantEOF:
					if err != io.EOF {
						t.Errorf("receive() вернула ошибку %v, ожидалась io.EOF", err)
					}
				case tt.wantErr != "":
					if err == nil || err.Error() != tt.wantErr {
						t.Errorf("receive() вернула ошибку %v, ожидалась %q", err, tt.wantErr)
					}
				default:
					if err != nil {
						t.Errorf("receive() вернула неожиданную ошибку %v", err)
					}
				}
				if remaining := len(messages); remaining != tt.wantRemaining {
					t.Errorf("в канале осталось %d сообщений, ожидалось %d", remaining, tt.wantRemaining)
				}
			})
		})
	}
}
