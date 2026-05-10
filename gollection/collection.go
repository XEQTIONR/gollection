package gollection

import "cmp"

type Collection[T comparable] struct {
	Items []T
}

type OrderableCollection[T cmp.Ordered] struct {
	Items []T
}

func Init[T comparable](items ...T) Collection[T] {
	return Collection[T]{
		Items: items,
	}
}

func Collect[T comparable](items ...T) Collection[T] {
	return Init(items...)
}

func InitFromArray[T comparable](items []T) Collection[T] {
	return Collection[T]{
		Items: items,
	}
}
