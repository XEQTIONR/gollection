package gollection

import (
	"errors"
	"math/rand"
	"slices"
	"sort"
)

func (c *Collection[T]) After(item T) (*T, error) {
	for i := range c.Items {
		if c.Items[i] == item {
			if i+1 >= len(c.Items) {
				return nil, errors.New("Last element")
			}
			return &c.Items[i+1], nil
		}
	}
	return nil, errors.New("Not found")
}

// func (c *Collection[T]) Average() (*T, error) {
// 	if len(c.Items) == 0 {
// 		return nil, errors.New("empty collection")
// 	}
// 	return &c.Items[len(c.Items)/2], nil
// }

func (c *Collection[T]) All(item T) []T {
	return c.Items
}

func (c *Collection[T]) Before(item T) (*T, error) {
	for i := range c.Items {
		if c.Items[i] == item {
			if i-1 < 0 {
				return nil, errors.New("First element")
			}
			return &c.Items[i-1], nil
		}
	}
	return nil, errors.New("Not found")
}

func (c *Collection[T]) Chunk(size int) ([]Collection[T], error) {

	chunks := make([]Collection[T], 0)

	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	for i := 0; i < len(c.Items); i += size {
		end := i + size
		if end > len(c.Items) {
			end = len(c.Items)
		}
		chunks = append(chunks, Collection[T]{Items: c.Items[i:end]})
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkBy(fn func(T) bool) ([]Collection[T], error) {
	chunks := make([]Collection[T], 0)
	for i := range c.Items {
		if fn(c.Items[i]) {
			chunks = append(chunks, Collection[T]{Items: []T{c.Items[i]}})
		}
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkUntil(fn func(T) bool) ([]Collection[T], error) {
	chunks := make([]Collection[T], 0)
	for i := range c.Items {
		if fn(c.Items[i]) {
			chunks = append(chunks, Collection[T]{Items: []T{c.Items[i]}})
		}
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkWhile(fn func(T) bool) ([]Collection[T], error) {
	chunks := make([]Collection[T], 0)
	for i := range c.Items {
		if fn(c.Items[i]) {
			chunks = append(chunks, Collection[T]{Items: []T{c.Items[i]}})
		}
	}
	return chunks, nil
}

func (c *Collection[T]) Contains(item T) bool {
	return slices.Contains(c.Items, item)
}

func (c *Collection[T]) Count() int {
	return len(c.Items)
}

func (c *Collection[T]) CountBy(fn func(T) bool) int {
	count := 0
	for i := range c.Items {
		if fn(c.Items[i]) {
			count++
		}
	}
	return count
}

func (c *Collection[T]) Counts() map[T]int {
	counts := make(map[T]int)
	for i := range c.Items {
		counts[c.Items[i]]++
	}
	return counts
}

// O(n**2) time complexity
func (c *Collection[T]) Diff(other Collection[T]) []T {
	diff := make([]T, 0)
	for i := range c.Items {
		if !slices.Contains(other.Items, c.Items[i]) {
			diff = append(diff, c.Items[i])
		}
	}
	return diff
}

func (c *Collection[T]) DoesntContain(item T) bool {
	return !slices.Contains(c.Items, item)
}

func (c *Collection[T]) DoesntHave(item T) bool {
	return c.DoesntContain(item)
}

// O(n**2) time complexity
func (c *Collection[T]) Duplicates() map[T]int {
	duplicates := make(map[T]int)
	counts := c.Counts()
	for k, v := range counts {
		if v > 1 {
			duplicates[k] = v
		}
	}
	return duplicates
}

func (c *Collection[T]) Every(fn func(T) bool) bool {
	for i := range c.Items {
		if !fn(c.Items[i]) {
			return false
		}
	}
	return true
}

func (c *Collection[T]) Filter(fn func(T) bool) []T {
	filtered := make([]T, 0)
	for i := range c.Items {
		if fn(c.Items[i]) {
			filtered = append(filtered, c.Items[i])
		}
	}
	return filtered
}

func (c *Collection[T]) First() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[0], nil
}

func (c *Collection[T]) Has(item T) bool {
	return c.Contains(item)
}

// O(n**2) time complexity
func (c *Collection[T]) HasAny(items ...T) bool {
	for _, item := range items {
		if c.Contains(item) {
			return true
		}
	}
	return false
}

// O(n**2) time complexity
func (c *Collection[T]) HasAll(items ...T) bool {
	for _, item := range items {
		if !c.Contains(item) {
			return false
		}
	}
	return true
}

// O(n**2) time complexity
func (c *Collection[T]) Intersect(other Collection[T]) Collection[T] {
	intersect := make([]T, 0)
	for i := range c.Items {
		if slices.Contains(other.Items, c.Items[i]) {
			intersect = append(intersect, c.Items[i])
		}
	}
	return Collection[T]{Items: intersect}
}

func (c *Collection[T]) Last() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[len(c.Items)-1], nil
}

func (c *Collection[T]) Length() int {
	return len(c.Items)
}

func (c *Collection[T]) Map(fn func(T) T) *Collection[T] {
	mapped := make([]T, 0)
	for i := range c.Items {
		mapped = append(mapped, fn(c.Items[i]))
	}
	return &Collection[T]{Items: mapped}
}

func (c *Collection[T]) Max() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[len(c.Items)-1], nil
}

func (c *Collection[T]) Median() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[len(c.Items)/2], nil
}

func (c *Collection[T]) Min() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[0], nil
}

func (c *Collection[T]) Mode() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[0], nil
}

func (c *Collection[T]) Multiply(multiplier int) *Collection[T] {
	if len(c.Items) == 0 {
		return c
	}

	for i := len(c.Items) - 1; i >= 0; i-- {
		repeat := slices.Repeat([]T{c.Items[i]}, multiplier-1)
		c.Items = slices.Insert(c.Items, i, repeat...)
	}

	return c
}

func (c *Collection[T]) Nth(n int) (*T, error) {
	if n < 0 {
		return nil, errors.New("n must be positive")
	}

	if n >= len(c.Items) {
		return nil, errors.New("n is out of range")
	}
	return &c.Items[n], nil
}

func (c *Collection[T]) PadLeft(length int, value T) *Collection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	c = &Collection[T]{Items: append(newSlice, c.Items...)}
	return c
}

func (c *Collection[T]) PadRight(length int, value T) *Collection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	c = &Collection[T]{Items: append(c.Items, newSlice...)}

	return c
}

func (c *Collection[T]) Percentage(fn func(T) bool) float64 {
	count := c.CountBy(fn)
	return float64(count) / float64(len(c.Items)) * 100
}

func (c *Collection[T]) Prepend(item ...T) *Collection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, c.Items...)
	newSlice = append(newSlice, item...)
	c.Items = newSlice
	return c
}

func (c *Collection[T]) Pop() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[len(c.Items)-1], nil
}

func (c *Collection[T]) Push(item ...T) *Collection[T] {
	c.Items = append(c.Items, item...)
	return c
}

func (c *Collection[T]) Random() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &c.Items[rand.Intn(len(c.Items))], nil
}

func (c *Collection[T]) Range(start int, end int) *Collection[T] {

	return &Collection[T]{Items: c.Items[start:end]}
}

// func (c *Collection[T]) Reduce (fn func(accumulator any, item T) T) *Collection[T] {
// 	reduced := c.Items[0]
// 	for i := 1; i < len(c.Items); i++ {
// 		reduced = fn(reduced, c.Items[i])
// 	}

// 	return &Collection[T]{Items:reduced}
// }

func (c *Collection[T]) Reject(fn func(T) bool) *Collection[T] {
	rejected := make([]T, 0)
	for i := range c.Items {
		if !fn(c.Items[i]) {
			rejected = append(rejected, c.Items[i])
		}
	}
	return &Collection[T]{Items: rejected}
}

func (c *Collection[T]) Reverse() *Collection[T] {
	reversed := make([]T, 0)
	for i := len(c.Items) - 1; i >= 0; i-- {
		reversed = append(reversed, c.Items[i])
	}
	return &Collection[T]{Items: reversed}
}

func (c *Collection[T]) Search(item T) int {
	return slices.Index(c.Items, item)
}

func (c *Collection[T]) Shift() (*T, error) {
	if len(c.Items) == 0 {
		return nil, errors.New("empty collection")
	}

	ret := c.Items[0]
	c.Items = c.Items[1:]

	return &ret, nil
}

func (c *Collection[T]) Shuffle() *Collection[T] {
	shuffled := make([]T, 0)
	for i := range c.Items {
		shuffled = append(shuffled, c.Items[i])
	}
	return &Collection[T]{Items: shuffled}
}

func (c *Collection[T]) Skip(n int) *Collection[T] {
	return &Collection[T]{Items: c.Items[n:]}
}

func (c *Collection[T]) SkipUntil(fn func(T) bool) *Collection[T] {
	for i := range c.Items {
		if fn(c.Items[i]) {
			return &Collection[T]{Items: c.Items[i:]}
		}
	}
	return c
}

func (c *Collection[T]) SkipWhile(fn func(T) bool) *Collection[T] {
	for i := range c.Items {
		if !fn(c.Items[i]) {
			return &Collection[T]{Items: c.Items[i:]}
		}
	}
	return c
}

func (c *Collection[T]) Slice(start int, end int) *Collection[T] {
	return &Collection[T]{Items: c.Items[start:end]}
}

func (c *OrderableCollection[T]) Sort(desc bool) *OrderableCollection[T] {
	sorted := make([]T, 0)
	for i := range c.Items {
		sorted = append(sorted, c.Items[i])
	}
	sort.Slice(sorted, func(i, j int) bool {
		if desc {
			return sorted[i] > sorted[j]
		}
		return sorted[i] < sorted[j]
	})
	return &OrderableCollection[T]{Items: sorted}
}

func (c *Collection[T]) Splice(start int, take int, items ...T) *Collection[T] {
	return &Collection[T]{Items: append(c.Items[:start], append(items, c.Items[start+take:]...)...)}
}

func (c *Collection[T]) Split(n int) []Collection[T] {

	chunks := make([]Collection[T], 0)

	for i := 0; i < len(c.Items)/n; i += n {
		chunks = append(chunks, Collection[T]{Items: c.Items[i : i+n]})
	}
	return chunks
}

func (c *Collection[T]) SplitInto(n int) []Collection[T] {
	return c.Split(c.Length() / n)
}

func (c *OrderableCollection[T]) SortAsc() *OrderableCollection[T] {
	return c.Sort(false)
}

func (c *OrderableCollection[T]) SortDesc() *OrderableCollection[T] {
	return c.Sort(true)
}

func (c *Collection[T]) Take(n int) *Collection[T] {
	return &Collection[T]{Items: c.Items[:n]}
}

func (c *Collection[T]) TakeUntil(fn func(T) bool) *Collection[T] {
	for i := range c.Items {
		if fn(c.Items[i]) {
			return &Collection[T]{Items: c.Items[:i]}
		}
	}
	return c
}

func (c *Collection[T]) TakeWhile(fn func(T) bool) *Collection[T] {
	for i := range c.Items {
		if !fn(c.Items[i]) {
			return &Collection[T]{Items: c.Items[:i]}
		}
	}
	return c
}

func (c *Collection[T]) Union(other Collection[T]) []T {
	m := make(map[T]bool)
	for _, v := range c.Items {
		m[v] = true
	}
	for _, v := range other.Items {
		m[v] = true
	}

	union := make([]T, 0, len(m))
	for v := range m {
		union = append(union, v)
	}
	return union
}

func (c *Collection[T]) Unique() *Collection[T] {
	unique := make([]T, 0)
	for i := range c.Items {
		if slices.Index(unique, c.Items[i]) == -1 {
			unique = append(unique, c.Items[i])
		}
	}
	return &Collection[T]{Items: unique}
}
