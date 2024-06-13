package ratelimiter

// Closer определяет механизм остановки rate limiter`а
type Closer interface {
	// Closed возвращает канал, который сигнализирует об остановке работы.
	// Используется в select statements
	Closed() <-chan struct{}
	// Close закрывает канал
	Close()
}
