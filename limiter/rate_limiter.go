package limiter

import (
	"time"

	"github.com/nlnaa11/ratelimiter"
	"github.com/nlnaa11/ratelimiter/closer"
	"github.com/nlnaa11/ratelimiter/config"
)

// Opt определяет тип функции для настройки rate limiter`а
type Opt func(*rateLimiter) error

type rateLimiter struct {
	tokensCh chan struct{}
	closer   ratelimiter.Closer
	config   ratelimiter.Config
}

func NewRateLimiter(oo ...Opt) (*rateLimiter, error) {
	rl := &rateLimiter{
		config: config.DefaultConfig,
		closer: closer.DefaultCloser,
	}

	for _, o := range oo {
		if err := o(rl); err != nil {
			return nil, err
		}
	}

	tokensCh := make(chan struct{}, rl.config.MaxSpike())
	for i := 0; i < rl.config.MaxSpike(); i++ {
		tokensCh <- struct{}{}
	}

	rl.tokensCh = tokensCh

	go rl.runRefreshTokens()

	return rl, nil
}

func MustRateLimiter(oo ...Opt) *rateLimiter {
	rl, err := NewRateLimiter(oo...)
	if err != nil {
		panic(err)
	}

	return rl
}

func (rl *rateLimiter) runRefreshTokens() {
	ticker := time.NewTicker(rl.config.RefreshInterval())
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.tokensCh <- struct{}{}
		case <-rl.closer.Closed():
			close(rl.tokensCh)
			return
		}
	}
}

func (rl *rateLimiter) Allowed() <-chan struct{} {
	return rl.tokensCh
}

func (rl *rateLimiter) Allow() bool {
	select {
	case <-rl.Allowed():
		return true
	default:
		return false
	}
}

func (rl *rateLimiter) Stop() {
	rl.closer.Close()
}
