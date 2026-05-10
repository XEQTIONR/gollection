package gollection

import "cmp"

type Collection[T comparable] struct {
	items []T
}

type OrderableCollection[T cmp.Ordered] struct {
	items []T
}

func Init[T comparable](items ...T) Collection[T] {
	return Collection[T]{
		items: items,
	}
}

func InitOrderable[T cmp.Ordered](items ...T) OrderableCollection[T] {
	return OrderableCollection[T]{
		items: items,
	}
}

func Collect[T comparable](items ...T) Collection[T] {
	return Init(items...)
}

func CollectOrderable[T cmp.Ordered](items ...T) OrderableCollection[T] {
	return InitOrderable(items...)
}

func InitFromSlice[T comparable](items []T) Collection[T] {
	return Collection[T]{
		items: items,
	}
}

func InitFromSliceOrderable[T cmp.Ordered](items []T) OrderableCollection[T] {
	return OrderableCollection[T]{
		items: items,
	}
}
