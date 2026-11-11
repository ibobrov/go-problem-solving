package _select

import (
	"errors"
	"io"
	"time"
)

func receive(messages <-chan string, timeout time.Duration) (string, error) {
	if timeout <= 0 {
		return "", errors.New("incorrect timeout")
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case message, ok := <-messages:
		if !ok {
			return "", io.EOF
		}
		return message, nil
	case <-timer.C:
	}

	return "", errors.New("timeout")
}
