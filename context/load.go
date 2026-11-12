package context

import "context"

func load(ctx context.Context) (string, error) {
	return loadWith(ctx, slowOperation)
}

func loadWith(ctx context.Context, operation func() string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	// Единственный результат помещается в буфер даже после ухода получателя.
	// Сама операция завершится самостоятельно, поскольку не принимает контекст.
	results := make(chan string, 1)
	go func() {
		results <- operation()
	}()

	select {
	case result := <-results:
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func slowOperation() string {
	return "результат"
}
