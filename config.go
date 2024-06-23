package ratelimiter

import "time"

// Config определяет конфигурацию rate limiter`а
type Config interface {
	// RefreshInterval возвращает интервал времени между обновлением токенов
	RefreshInterval() time.Duration
	// MaxSpike возвращает максимально допустимый всплекс запросов
	MaxSpike() int
}
