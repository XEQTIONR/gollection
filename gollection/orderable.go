package gollection

import (
	"cmp"
	"errors"
	"math/rand"
	"slices"
)

func (c *OrderableCollection[T]) After(item T) (*T, error) {
	items := c.All()

	for i, v := range items {
		if v == item {
			if i+1 >= c.Length() {
				return nil, errors.New("Last element")
			}
			return &items[i+1], nil
		}
	}
	return nil, errors.New("Not found")
}

func (c *OrderableCollection[T]) All() []T {
	return []T(*c)
}

func (c *OrderableCollection[T]) Append(items ...T) *OrderableCollection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, c.All()...)
	newSlice = append(newSlice, items...)
	*c = newSlice
	return c
}

func (c *OrderableCollection[T]) At(index int) *T {
	items := c.All()
	if index < 0 || index >= c.Length() {
		return nil
	}
	return &items[index]
}

func (c *OrderableCollection[T]) Before(item T) (*T, error) {
	items := c.All()

	for i := range items {
		if items[i] == item {
			if i-1 < 0 {
				return nil, errors.New("First element")
			}
			return &items[i-1], nil
		}
	}
	return nil, errors.New("Not found")
}

func (c *OrderableCollection[T]) Chunk(size int) ([]OrderableCollection[T], error) {

	chunks := make([]OrderableCollection[T], 0)
	items := c.All()

	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	for i := 0; i < len(items); i += size {
		end := i + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, InitFromSliceOrderable(items[i:end]))
	}
	return chunks, nil
}

func (c *OrderableCollection[T]) ChunkBy(fn func(T) bool) ([]OrderableCollection[T], error) {
	items := c.All()
	chunks := make([]OrderableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitOrderable(items[i]))
		}
	}
	return chunks, nil
}

func (c *OrderableCollection[T]) ChunkUntil(fn func(T) bool) ([]OrderableCollection[T], error) {
	items := c.All()
	chunks := make([]OrderableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitOrderable(items[i]))
		}
	}
	return chunks, nil
}

func (c *OrderableCollection[T]) ChunkWhile(fn func(T) bool) ([]OrderableCollection[T], error) {
	items := c.All()
	chunks := make([]OrderableCollection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, InitOrderable(items[i]))
		}
	}
	return chunks, nil
}

func (c *OrderableCollection[T]) Contains(item T) bool {
	return slices.Contains(c.All(), item)
}

func (c *OrderableCollection[T]) Count() int {
	return len(c.All())
}

func (c *OrderableCollection[T]) CountBy(fn func(T) bool) int {
	count := 0
	for _, v := range c.All() {
		if fn(v) {
			count++
		}
	}
	return count
}

func (c *OrderableCollection[T]) Counts() map[T]int {
	counts := make(map[T]int)
	for _, v := range c.All() {
		counts[v]++
	}
	return counts
}

// O(n**2) time complexity
func (c *OrderableCollection[T]) Diff(other OrderableCollection[T]) []T {
	diff := make([]T, 0)
	items := c.All()
	for i := range items {
		if !slices.Contains(other.All(), items[i]) {
			diff = append(diff, items[i])
		}
	}
	return diff
}

func (c *OrderableCollection[T]) DoesntContain(item T) bool {
	return !slices.Contains(c.All(), item)
}

func (c *OrderableCollection[T]) DoesntHave(item T) bool {
	return c.DoesntContain(item)
}

// O(n**2) time complexity
func (c *OrderableCollection[T]) Duplicates() map[T]int {
	duplicates := make(map[T]int)
	counts := c.Counts()
	for k, v := range counts {
		if v > 1 {
			duplicates[k] = v
		}
	}
	return duplicates
}

func (c *OrderableCollection[T]) Every(fn func(T) bool) bool {
	for _, v := range c.All() {
		if !fn(v) {
			return false
		}
	}
	return true
}

func (c *OrderableCollection[T]) Filter(fn func(T) bool) []T {
	items := c.All()
	filtered := make([]T, 0)
	for i := range items {
		if fn(items[i]) {
			filtered = append(filtered, items[i])
		}
	}
	return filtered
}

func (c *OrderableCollection[T]) First() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[0], nil
}

func (c *OrderableCollection[T]) Has(item T) bool {
	return c.Contains(item)
}

// O(n**2) time complexity
func (c *OrderableCollection[T]) HasAny(items ...T) bool {
	cMap := make(map[T]bool)

	for _, item := range c.All() {
		cMap[item] = true
	}

	for _, item := range items {
		if _, ok := cMap[item]; ok {
			return true
		}
	}

	return false
}

// O(n**2) time complexity
func (c *OrderableCollection[T]) HasAll(items ...T) bool {
	cMap := make(map[T]bool)

	for _, item := range c.All() {
		cMap[item] = true
	}
	for _, item := range items {
		if _, ok := cMap[item]; !ok {
			return false
		}
	}
	return true
}

// O(n**2) time complexity
func (c *OrderableCollection[T]) Intersect(other OrderableCollection[T]) OrderableCollection[T] {
	items := c.All()
	intersect := make([]T, 0)
	for i := range items {
		if slices.Contains(other.All(), items[i]) {
			intersect = append(intersect, items[i])
		}
	}
	return OrderableCollection[T](intersect)
}

func (c *OrderableCollection[T]) Last() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[len(items)-1], nil
}

func (c *OrderableCollection[T]) Length() int {
	return c.Count()
}

func (c *OrderableCollection[T]) Map(fn func(T) T) *OrderableCollection[T] {
	items := c.All()
	mapped := make([]T, 0)
	for i := range items {
		mapped = append(mapped, fn(items[i]))
	}

	*c = InitFromSliceOrderable(mapped)
	return c
}

func (c *OrderableCollection[T]) Max() (*T, error) {
	items := c.All()
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	highest := items[0]
	for _, v := range items {
		if cmp.Compare(v, highest) > 0 {
			highest = v
		}
	}

	return &highest, nil
}

func (c *OrderableCollection[T]) Median() (*T, error) {
	items := c.All()
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[len(items)/2], nil
}

func (c *OrderableCollection[T]) Min() (*T, error) {
	items := c.All()
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	lowest := items[0]
	for _, v := range items {
		if cmp.Compare(v, lowest) < 0 {
			lowest = v
		}
	}

	return &lowest, nil
}

func (c *OrderableCollection[T]) Mode() (*T, int) {
	items := c.Counts()

	var mode *T
	highest := -1

	for k, v := range items {
		if v > highest {
			*mode = k
		}
	}
	return mode, highest
}

func (c *OrderableCollection[T]) Multiply(multiplier int) *OrderableCollection[T] {
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

func (c *OrderableCollection[T]) Nth(n int) (*T, error) {
	items := c.All()
	if n < 0 {
		return nil, errors.New("n must be positive")
	}

	if n >= len(items) {
		return nil, errors.New("n is out of range")
	}
	return &items[n], nil
}

func (c *OrderableCollection[T]) PadLeft(length int, value T) *OrderableCollection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append(newSlice, c.All()...)

	*c = InitFromSliceOrderable(appended)

	return c

}

func (c *OrderableCollection[T]) PadRight(length int, value T) *OrderableCollection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append([]T{}, c.All()...)
	appended = append(appended, newSlice...)

	*c = InitFromSliceOrderable(appended)

	return c
}

func (c *OrderableCollection[T]) Prepend(item ...T) *OrderableCollection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, item...)
	newSlice = append(newSlice, c.All()...)
	*c = newSlice
	return c
}

func (c *OrderableCollection[T]) Pop() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	ret := c.At(c.Length() - 1)

	*c = c.All()[:c.Length()-1]
	return ret, nil
}

func (c *OrderableCollection[T]) Push(item ...T) *OrderableCollection[T] {
	*c = append(c.All(), item...)
	return c
}

func (c *OrderableCollection[T]) Random() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	return c.At(rand.Intn(c.Length())), nil
}

func (c *OrderableCollection[T]) Range(start int, end int) *OrderableCollection[T] {
	collection := InitFromSliceOrderable(c.All()[start:end])

	return &collection
}

func (c *OrderableCollection[T]) Reject(fn func(T) bool) *OrderableCollection[T] {
	rejected := make([]T, 0)
	for _, v := range c.All() {
		if !fn(v) {
			rejected = append(rejected, v)
		}
	}

	collection := InitFromSliceOrderable(rejected)

	return &collection
}

func (c *OrderableCollection[T]) Reverse() *OrderableCollection[T] {
	reversed := make([]T, 0)
	for i := c.Length() - 1; i >= 0; i-- {
		reversed = append(reversed, *c.At(i))
	}

	collection := InitFromSliceOrderable(reversed)

	return &collection
}

func (c *OrderableCollection[T]) Search(item T) int {
	return slices.Index(c.All(), item)
}

func (c *OrderableCollection[T]) Shift() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}

	ret := c.At(0)
	*c = c.All()[1:]

	return ret, nil
}

func (c *OrderableCollection[T]) Shuffle() *OrderableCollection[T] {
	shuffled := make([]T, 0)
	for i := range c.All() {
		shuffled = append(shuffled, *c.At(i))
	}

	collection := InitFromSliceOrderable(shuffled)

	return &collection
}

func (c *OrderableCollection[T]) Skip(n int) *OrderableCollection[T] {
	collection := InitFromSliceOrderable(c.All()[n:])
	return &collection
}

func (c *OrderableCollection[T]) SkipUntil(fn func(T) bool) *OrderableCollection[T] {
	for i := range c.All() {
		if fn(*c.At(i)) {
			collection := InitFromSliceOrderable(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *OrderableCollection[T]) SkipWhile(fn func(T) bool) *OrderableCollection[T] {
	for i := range c.All() {
		if !fn(*c.At(i)) {
			collection := InitFromSliceOrderable(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *OrderableCollection[T]) Slice(start int, end int) *OrderableCollection[T] {
	collection := InitFromSliceOrderable(c.All()[start:end])
	return &collection
}

func (c *OrderableCollection[T]) Splice(start int, take int, items ...T) *OrderableCollection[T] {
	collection := InitFromSliceOrderable(append(c.All()[:start], append(items, c.All()[start+take:]...)...))
	return &collection
}

func (c *OrderableCollection[T]) Split(n int) []OrderableCollection[T] {

	chunks := make([]OrderableCollection[T], 0)

	for i := 0; i < c.Length()/n; i += n {
		chunks = append(chunks, InitFromSliceOrderable(c.All()[i:i+n]))
	}
	return chunks
}

func (c *OrderableCollection[T]) SplitInto(n int) []OrderableCollection[T] {
	return c.Split(c.Length() / n)
}

func (c *OrderableCollection[T]) Sort(desc bool) *OrderableCollection[T] {

	items := c.All()
	slices.SortFunc(items, func(a, b T) int {

		if desc {
			return cmp.Compare(b, a) * (-1)
		}
		return cmp.Compare(a, b)
	})

	*c = InitFromSliceOrderable(items)
	return c
}

func (c *OrderableCollection[T]) SortAsc() *OrderableCollection[T] {
	return c.Sort(false)
}

func (c *OrderableCollection[T]) SortDesc() *OrderableCollection[T] {
	return c.Sort(true)
}

func (c *OrderableCollection[T]) Take(n int) *OrderableCollection[T] {
	*c = InitFromSliceOrderable(c.All()[:n])
	return c
}

func (c *OrderableCollection[T]) TakeUntil(fn func(T) bool) *OrderableCollection[T] {
	for _, v := range c.All() {
		if fn(v) {
			collection := InitOrderable(v)
			return &collection
		}
	}
	return c
}

func (c *OrderableCollection[T]) TakeWhile(fn func(T) bool) *OrderableCollection[T] {
	for _, v := range c.All() {
		if !fn(v) {
			collection := InitOrderable(v)
			return &collection
		}
	}
	return c
}

func (c *OrderableCollection[T]) Union(other OrderableCollection[T]) *OrderableCollection[T] {
	m := make(map[T]bool)
	for _, v := range c.All() {
		m[v] = true
	}
	for _, v := range other.All() {
		m[v] = true
	}

	union := make([]T, 0, len(m))
	for v := range m {
		union = append(union, v)
	}

	collection := InitFromSliceOrderable(union)

	return &collection
}

func (c *OrderableCollection[T]) Unique() *OrderableCollection[T] {
	counts := c.Counts()
	unique := make([]T, 0)
	for k, v := range counts {
		if v == 1 {
			unique = append(unique, k)
		}
	}

	collection := InitFromSliceOrderable(unique)
	return &collection
}
