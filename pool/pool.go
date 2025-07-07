package pool

import "sync"

type Pool[T any] interface {
	Get() T
	PUT(t T)
}

type _Pool[T any] struct {
	pool sync.Pool
}

func (p *_Pool[T]) Get() T {
	v := p.pool.Get()
	return v.(T)
}

func (p *_Pool[T]) PUT(t T) {
	p.pool.Put(t)
}
func New[T any](new func() T) Pool[T] {
	return &_Pool[T]{
		pool: sync.Pool{New: func() any { return new() }},
	}
}
