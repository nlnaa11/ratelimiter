package closer

import "sync"

var DefaultCloser = MustCloser()

type closer struct {
	once    sync.Once
	closeCh chan struct{}
}

func NewCloser() (*closer, error) {
	return &closer{
		closeCh: make(chan struct{}),
		once:    sync.Once{},
	}, nil
}

func MustCloser() *closer {
	c, err := NewCloser()
	if err != nil {
		panic(err)
	}

	return c
}

// Closed возвращает канал, который сигнализирует об остановке rateLimmiter`а.
// Используется в select statements
func (c *closer) Closed() <-chan struct{} {
	return c.closeCh
}

// Close закрывает канал
func (c *closer) Close() {
	c.once.Do(func() {
		close(c.closeCh)
	})
}
