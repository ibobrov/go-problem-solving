package context

import (
	"context"
	"time"
)

func process(ctx context.Context, jobs <-chan int) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case _, ok := <-jobs:
			if !ok {
				return ctx.Err()
			}
		}

		timer := time.NewTimer(200 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
}
