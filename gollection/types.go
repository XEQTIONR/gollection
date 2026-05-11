package gollection

import "cmp"

type CollectionInterface[T any] interface {
	All() []T
	Length() int
}

type Collection[T comparable] []T

type OrderableCollection[T cmp.Ordered] []T

type ReflectableCollection[T map[string]any] []T
