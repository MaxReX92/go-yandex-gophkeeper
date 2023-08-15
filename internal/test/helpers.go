package test

func ArrayToChan[T any](items []T) <-chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for _, item := range items {
			result <- item
		}
	}()

	return result
}
