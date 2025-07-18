package utils

import "sort"

type sorter[T any] struct {
	list []T
	less func(a, b T) bool
}

func (s *sorter[T]) Len() int {
	return len(s.list)
}

func (s *sorter[T]) Less(i, j int) bool {
	return s.less(s.list[i], s.list[j])
}

func (s *sorter[T]) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}

func Sort[T any](list []T, less func(a, b T) bool) {
	i := &sorter[T]{
		list: list,
		less: less,
	}
	sort.Sort(i)

}

type sorterIndex[T any] struct {
	list []T
	less func(a, b int) bool
}

func (s *sorterIndex[T]) Len() int {
	return len(s.list)
}

func (s *sorterIndex[T]) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *sorterIndex[T]) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}

func SortI[T any](list []T, less func(a, b int) bool) {
	i := &sorterIndex[T]{
		list: list,
		less: less,
	}
	sort.Sort(i)

}
