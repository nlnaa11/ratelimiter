package config

import "time"

func WithRequestLimit(limit int, interval time.Duration) Opt {
	return func(c *config) error {
		c.refreshInterval =
			time.Duration(interval.Nanoseconds() / int64(limit))

		return nil
	}
}

func WithMaxSpike(spike int) Opt {
	return func(c *config) error {
		c.maxSpike = spike

		return nil
	}
}
