package limiter

import (
	"github.com/nlnaa11/ratelimiter"
)

func WithConfig(c ratelimiter.Config) Opt {
	return func(rl *rateLimiter) error {
		rl.config = c

		return nil
	}
}

func WithCloser(c ratelimiter.Closer) Opt {
	return func(rl *rateLimiter) error {
		rl.closer = c

		return nil
	}
}
