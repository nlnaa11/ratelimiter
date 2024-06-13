package config

import "time"

var (
	defaultRefreshInterval = time.Duration(time.Second / 100)
	defaultMaxSpike        = 1
)

// Opt определяет тип функции для настройки конфига rate limiter`а
type Opt func(*config) error

type config struct {
	refreshInterval time.Duration
	maxSpike        int
}

func NewConfig(oo ...Opt) (*config, error) {
	c := &config{defaultRefreshInterval, defaultMaxSpike}

	for _, o := range oo {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func MustConfig(oo ...Opt) *config {
	c, err := NewConfig(oo...)
	if err != nil {
		panic(err)
	}

	return c
}

func (c config) RefreshInterval() time.Duration {
	return c.refreshInterval
}

func (c config) MaxSpike() int {
	return c.maxSpike
}
