package gollection

import "cmp"

type Collection[T comparable] []T

type OrderableCollection[T cmp.Ordered] []T

type ReflectableCollection[T map[string]any] []T

func Init[T comparable](items ...T) Collection[T] {
	return items
}

func InitOrderable[T cmp.Ordered](items ...T) OrderableCollection[T] {
	return items
}

func InitReflectable[T map[string]any](items ...T) ReflectableCollection[T] {
	return items
}

func Collect[T comparable](items ...T) Collection[T] {
	return Init(items...)
}

func CollectOrderable[T cmp.Ordered](items ...T) OrderableCollection[T] {
	return InitOrderable(items...)
}

func CollectReflectable[T map[string]any](items ...T) ReflectableCollection[T] {
	return InitReflectable(items...)
}

func InitFromSlice[T comparable](items []T) Collection[T] {
	return items
}

func InitFromSliceOrderable[T cmp.Ordered](items []T) OrderableCollection[T] {
	return items
}

func InitFromSliceReflectable[T map[string]any](items []T) ReflectableCollection[T] {
	return items
}
