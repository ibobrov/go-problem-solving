package context

import (
	"context"
	"errors"
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

func TestPrintCancellation(t *testing.T) {
	t.Parallel()

	var output strings.Builder
	printCancellation(&output)

	const want = "context canceled\n<nil>\ncontext canceled\n"
	if got := output.String(); got != want {
		t.Errorf("вывод = %q, ожидалось %q", got, want)
	}
}

func TestProcess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		jobs    int
		close   bool
		timeout time.Duration
		cancel  bool
		wantErr error
		elapsed time.Duration
	}{
		{name: "закрытый пустой канал", close: true},
		{name: "два задания до закрытия", jobs: 2, close: true, elapsed: 400 * time.Millisecond},
		{name: "дедлайн во время ожидания задания", timeout: time.Second, wantErr: context.DeadlineExceeded, elapsed: time.Second},
		{name: "дедлайн во время обработки", jobs: 2, close: true, timeout: 50 * time.Millisecond, wantErr: context.DeadlineExceeded, elapsed: 50 * time.Millisecond},
		{name: "заранее отменённый контекст", jobs: 1, close: true, cancel: true, wantErr: context.Canceled},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				ctx := context.Background()
				if tt.timeout > 0 {
					var cancel context.CancelFunc
					ctx, cancel = context.WithTimeout(ctx, tt.timeout)
					defer cancel()
				}
				if tt.cancel {
					var cancel context.CancelFunc
					ctx, cancel = context.WithCancel(ctx)
					cancel()
				}

				jobs := make(chan int, tt.jobs)
				for i := range tt.jobs {
					jobs <- i
				}
				if tt.close {
					close(jobs)
				}

				start := time.Now()
				if err := process(ctx, jobs); !errors.Is(err, tt.wantErr) {
					t.Errorf("process() вернула %v, ожидалось %v", err, tt.wantErr)
				}
				if elapsed := time.Since(start); elapsed != tt.elapsed {
					t.Errorf("process() заняла %v, ожидалось %v", elapsed, tt.elapsed)
				}
			})
		})
	}
}

func TestLoad(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		duration time.Duration
		timeout  time.Duration
		want     string
		wantErr  error
		elapsed  time.Duration
	}{
		{name: "результат до дедлайна", duration: time.Second, timeout: 2 * time.Second, want: "результат", elapsed: time.Second},
		{name: "дедлайн до результата", duration: 2 * time.Second, timeout: time.Second, wantErr: context.DeadlineExceeded, elapsed: time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
				defer cancel()

				finished := make(chan struct{})
				start := time.Now()
				got, err := loadWith(ctx, func() string {
					defer close(finished)
					time.Sleep(tt.duration)
					return "результат"
				})
				if got != tt.want || !errors.Is(err, tt.wantErr) {
					t.Errorf("loadWith() = (%q, %v), ожидалось (%q, %v)", got, err, tt.want, tt.wantErr)
				}
				if elapsed := time.Since(start); elapsed != tt.elapsed {
					t.Errorf("loadWith() заняла %v, ожидалось %v", elapsed, tt.elapsed)
				}
				// Даём операции завершиться. synctest обнаружит блокировку на отправке.
				<-finished
			})
		})
	}
}
