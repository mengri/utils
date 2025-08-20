package register

import (
	"fmt"
	"github.com/mengri/utils/untyped"
)

type Registers[T any] interface {
	Register(name string, t T) error
	All() map[string]T
	Get(name string) (T, bool)
}

type imlRegisters[T any] struct {
	untyped.Untyped[string, T]
}

func NewRegisters[T any]() Registers[T] {

	return &imlRegisters[T]{
		Untyped: untyped.BuildUntyped[string, T](),
	}
}

func (rs *imlRegisters[T]) Register(name string, t T) error {
	_, has := rs.Untyped.Get(name)
	if has {
		return fmt.Errorf("%s exits", name)
	}
	rs.Set(name, t)
	return nil
}
