package gollection

import (
	"errors"
	"math/rand"
	"slices"
)

func (c *ReflectableCollection[T]) All() []T {
	return []T(*c)
}

func (c *ReflectableCollection[T]) Append(items ...T) *ReflectableCollection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, c.All()...)
	newSlice = append(newSlice, items...)
	*c = newSlice
	return c
}

func (c *ReflectableCollection[T]) At(index int) *T {
	items := c.All()
	if index < 0 || index >= c.Length() {
		return nil
	}
	return &items[index]
}

// func (c *ReflectableCollection[T]) After(item T) (*T, error) {
// 	items := c.All()

// 	for i, v := range items {
// 		if v == item {
// 			if i+1 >= c.Length() {
// 				return nil, errors.New("Last element")
// 			}
// 			return &items[i+1], nil
// 		}
// 	}
// 	return nil, errors.New("Not found")
// }

// func (c *ReflectableCollection[T]) Before(item T) (*T, error) {
// 	items := c.All()

// 	for i := range items {
// 		if items[i] == item {
// 			if i-1 < 0 {
// 				return nil, errors.New("First element")
// 			}
// 			return &items[i-1], nil
// 		}
// 	}
// 	return nil, errors.New("Not found")
// }

func (c *ReflectableCollection[T]) Chunk(size int) ([]ReflectableCollection[T], error) {

	chunks := make([]ReflectableCollection[T], 0)
	items := c.All()

	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	for i := 0; i < len(items); i += size {
		end := i + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, InitFromSliceReflectable(items[i:end]))
	}
	return chunks, nil
}

func (c *ReflectableCollection[T]) ChunkBy(fn func(T) bool) ([]ReflectableCollection[T], error) {
	items := c.All()
	chunks := make([]ReflectableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitReflectable(items[i]))
		}
	}
	return chunks, nil
}

func (c *ReflectableCollection[T]) ChunkUntil(fn func(T) bool) ([]ReflectableCollection[T], error) {
	items := c.All()
	chunks := make([]ReflectableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitReflectable(items[i]))
		}
	}
	return chunks, nil
}

func (c *ReflectableCollection[T]) ChunkWhile(fn func(T) bool) ([]ReflectableCollection[T], error) {
	items := c.All()
	chunks := make([]ReflectableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitReflectable(items[i]))
		}
	}
	return chunks, nil
}

// func (c *Collection[T]) Contains(item T) bool {
// 	return slices.Contains(c.All(), item)
// }

func (c *ReflectableCollection[T]) Count() int {
	return len(c.All())
}

// func (c *Collection[T]) CountBy(fn func(T) bool) int {
// 	count := 0
// 	for _, v := range c.All() {
// 		if fn(v) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func (c *Collection[T]) Counts() map[T]int {
// 	counts := make(map[T]int)
// 	for _, v := range c.All() {
// 		counts[v]++
// 	}
// 	return counts
// }

// O(n**2) time complexity
// func (c *Collection[T]) Diff(other Collection[T]) []T {
// 	diff := make([]T, 0)
// 	items := c.All()
// 	for i := range items {
// 		if !slices.Contains(other.All(), items[i]) {
// 			diff = append(diff, items[i])
// 		}
// 	}
// 	return diff
// }

// func (c *Collection[T]) DoesntContain(item T) bool {
// 	return !slices.Contains(c.All(), item)
// }

// func (c *Collection[T]) DoesntHave(item T) bool {
// 	return c.DoesntContain(item)
// }

// O(n**2) time complexity
// func (c *Collection[T]) Duplicates() map[T]int {
// 	duplicates := make(map[T]int)
// 	counts := c.Counts()
// 	for k, v := range counts {
// 		if v > 1 {
// 			duplicates[k] = v
// 		}
// 	}
// 	return duplicates
// }

func (c *ReflectableCollection[T]) Every(fn func(T) bool) bool {
	for _, v := range c.All() {
		if !fn(v) {
			return false
		}
	}
	return true
}

func (c *ReflectableCollection[T]) Filter(fn func(T) bool) []T {
	items := c.All()
	filtered := make([]T, 0)
	for i := range items {
		if fn(items[i]) {
			filtered = append(filtered, items[i])
		}
	}
	return filtered
}

func (c *ReflectableCollection[T]) First() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[0], nil
}

// func (c *Collection[T]) Has(item T) bool {
// 	return c.Contains(item)
// }

// func (c *Collection[T]) HasAny(items ...T) bool {
// 	cMap := make(map[T]bool)

// 	for _, item := range c.All() {
// 		cMap[item] = true
// 	}

// 	for _, item := range items {
// 		if _, ok := cMap[item]; ok {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (c *Collection[T]) HasAll(items ...T) bool {
// 	cMap := make(map[T]bool)

// 	for _, item := range c.All() {
// 		cMap[item] = true
// 	}
// 	for _, item := range items {
// 		if _, ok := cMap[item]; !ok {
// 			return false
// 		}
// 	}
// 	return true
// }

// // O(n**2) time complexity
// func (c *Collection[T]) Intersect(other Collection[T]) Collection[T] {
// 	items := c.All()
// 	intersect := make([]T, 0)
// 	for i := range items {
// 		if slices.Contains(other.All(), items[i]) {
// 			intersect = append(intersect, items[i])
// 		}
// 	}
// 	return Collection[T](intersect)
// }

func (c *ReflectableCollection[T]) Last() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[len(items)-1], nil
}

func (c *ReflectableCollection[T]) Length() int {
	return c.Count()
}

func (c *ReflectableCollection[T]) Map(fn func(T) T) *ReflectableCollection[T] {
	items := c.All()
	mapped := make([]T, 0)
	for i := range items {
		mapped = append(mapped, fn(items[i]))
	}

	*c = InitFromSliceReflectable(mapped)
	return c
}

// func (c *Collection[T]) Mode() *T {
// 	items := c.Counts()

// 	var mode *T
// 	highest := -1

// 	for k, v := range items {
// 		if v > highest {
// 			*mode = k
// 		}
// 	}
// 	return mode
// }

func (c *ReflectableCollection[T]) Multiply(multiplier int) *ReflectableCollection[T] {
	items := c.All()
	if len(items) == 0 {
		return c
	}

	for i := len(items) - 1; i >= 0; i-- {
		repeat := slices.Repeat([]T{items[i]}, multiplier-1)
		items = slices.Insert(items, i, repeat...)
	}

	return c
}

func (c *ReflectableCollection[T]) Nth(n int) (*T, error) {
	items := c.All()
	if n < 0 {
		return nil, errors.New("n must be positive")
	}

	if n >= len(items) {
		return nil, errors.New("n is out of range")
	}
	return &items[n], nil
}

func (c *ReflectableCollection[T]) PadLeft(length int, value T) *ReflectableCollection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append(newSlice, c.All()...)

	*c = InitFromSliceReflectable(appended)

	return c

}

func (c *ReflectableCollection[T]) PadRight(length int, value T) *ReflectableCollection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append([]T{}, c.All()...)
	appended = append(appended, newSlice...)

	*c = InitFromSliceReflectable[T](appended)

	return c
}

func (c *ReflectableCollection[T]) Prepend(items ...T) *ReflectableCollection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, items...)
	newSlice = append(newSlice, c.All()...)
	*c = newSlice
	return c
}

func (c *ReflectableCollection[T]) Pop() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	ret := c.At(c.Length() - 1)

	*c = c.All()[:c.Length()-1]
	return ret, nil
}

func (c *ReflectableCollection[T]) Push(items ...T) *ReflectableCollection[T] {
	*c = append(c.All(), items...)
	return c
}

func (c *ReflectableCollection[T]) Random() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	return c.At(rand.Intn(c.Length())), nil
}

func (c *ReflectableCollection[T]) Range(start int, end int) *ReflectableCollection[T] {
	collection := InitFromSliceReflectable(c.All()[start:end])

	return &collection
}

func (c *ReflectableCollection[T]) Reject(fn func(T) bool) *ReflectableCollection[T] {
	rejected := make([]T, 0)
	for _, v := range c.All() {
		if !fn(v) {
			rejected = append(rejected, v)
		}
	}

	collection := InitFromSliceReflectable(rejected)

	return &collection
}

func (c *ReflectableCollection[T]) Reverse() *ReflectableCollection[T] {
	reversed := make([]T, 0)
	for i := c.Length() - 1; i >= 0; i-- {
		reversed = append(reversed, *c.At(i))
	}

	collection := InitFromSliceReflectable(reversed)

	return &collection
}

// func (c *Collection[T]) Search(item T) int {
// 	return slices.Index(c.All(), item)
// }

func (c *ReflectableCollection[T]) Shift() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}

	ret := c.At(0)
	*c = c.All()[1:]

	return ret, nil
}

func (c *ReflectableCollection[T]) Shuffle() *ReflectableCollection[T] {
	shuffled := make([]T, 0)
	for i := range c.All() {
		shuffled = append(shuffled, *c.At(i))
	}

	collection := InitFromSliceReflectable(shuffled)

	return &collection
}

func (c *ReflectableCollection[T]) Skip(n int) *ReflectableCollection[T] {
	collection := InitFromSliceReflectable(c.All()[n:])
	return &collection
}

func (c *ReflectableCollection[T]) SkipUntil(fn func(T) bool) *ReflectableCollection[T] {
	for i := range c.All() {
		if fn(*c.At(i)) {
			collection := InitFromSliceReflectable(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *ReflectableCollection[T]) SkipWhile(fn func(T) bool) *ReflectableCollection[T] {
	for i := range c.All() {
		if !fn(*c.At(i)) {
			collection := InitFromSliceReflectable(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *ReflectableCollection[T]) Slice(start int, end int) *ReflectableCollection[T] {
	collection := InitFromSliceReflectable(c.All()[start:end])
	return &collection
}

func (c *ReflectableCollection[T]) Splice(start int, take int, items ...T) *ReflectableCollection[T] {
	collection := InitFromSliceReflectable(append(c.All()[:start], append(items, c.All()[start+take:]...)...))
	return &collection
}

func (c *ReflectableCollection[T]) Split(n int) []ReflectableCollection[T] {

	chunks := make([]ReflectableCollection[T], 0)

	for i := 0; i < c.Length()/n; i += n {
		chunks = append(chunks, InitFromSliceReflectable(c.All()[i:i+n]))
	}
	return chunks
}

func (c *ReflectableCollection[T]) SplitInto(n int) []ReflectableCollection[T] {
	return c.Split(c.Length() / n)
}

func (c *ReflectableCollection[T]) Take(n int) *ReflectableCollection[T] {
	collection := InitFromSliceReflectable(c.All()[:n])
	return &collection
}

func (c *ReflectableCollection[T]) TakeUntil(fn func(T) bool) *ReflectableCollection[T] {
	for _, v := range c.All() {
		if fn(v) {
			collection := InitFromSliceReflectable([]T{v})
			return &collection
		}
	}
	return c
}

func (c *ReflectableCollection[T]) TakeWhile(fn func(T) bool) *ReflectableCollection[T] {
	for _, v := range c.All() {
		if !fn(v) {
			collection := InitFromSliceReflectable([]T{v})
			return &collection
		}
	}
	return c
}

// func (c *Collection[T]) Union(other Collection[T]) *Collection[T] {
// 	m := make(map[T]bool)
// 	for _, v := range c.All() {
// 		m[v] = true
// 	}
// 	for _, v := range other.All() {
// 		m[v] = true
// 	}

// 	union := make([]T, 0, len(m))
// 	for v := range m {
// 		union = append(union, v)
// 	}

// 	collection := InitFromSlice(union)

// 	return &collection
// }

// func (c *Collection[T]) Unique() *Collection[T] {
// 	counts := c.Counts()
// 	unique := make([]T, 0)
// 	for k, v := range counts {
// 		if v == 1 {
// 			unique = append(unique, k)
// 		}
// 	}

// 	collection := InitFromSlice(unique)
// 	return &collection
// }
