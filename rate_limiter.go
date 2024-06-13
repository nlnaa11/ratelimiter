package ratelimiter

// RateLimiter определяет контракт для работы с rate limiter`ом
type RateLimiter interface {
	// Allowed возвращает канал, сигнализирующий о разрешение на выполнение запроса.
	// Используется в select statements и в диапазонном цикле
	Allowed() <-chan struct{}
	// Allow возвращает ответ, может ли быть выполнен запрос
	Allow() bool

	// Stop останавливает работу rate limiter`а
	Stop()
}
