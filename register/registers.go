package register

import (
	"fmt"
	"github.com/mengri/utils/untyped"
)

type Registers[T any] interface {
	Register(name string, T any) error
	ALl() []T
	Get(name string) (T, bool)
}

type imlRegisters[T any] struct {
	objects untyped.Untyped[string, T]
}

func NewRegisters[T any]() Registers[T] {

	return &imlRegisters[T]{
		objects: untyped.BuildUntyped[string, T](),
	}
}

func (rs *imlRegisters[T]) Register(name string, T any) error {
	_, has := rs.objects.Get(name)
	if has {
		return fmt.Errorf("%s exits", name)
	}
	rs.objects.Set(name, T)
	return nil
}

func (rs *imlRegisters[T]) ALl() []T {
	return rs.objects.List()
}

func (rs *imlRegisters[T]) Get(name string) (T, bool) {
	return rs.objects.Get(name)
}
