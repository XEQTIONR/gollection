package gollection

import (
	"errors"
	"math/rand"
	"slices"
)

func (c *Collection[T]) After(item T) (*T, error) {
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

func (c *Collection[T]) All() []T {
	return []T(*c)
}

func (c *Collection[T]) Append(items ...T) *Collection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, c.All()...)
	newSlice = append(newSlice, items...)
	*c = newSlice
	return c
}

func (c *Collection[T]) At(index int) *T {
	items := c.All()
	if index < 0 || index >= c.Length() {
		return nil
	}
	return &items[index]
}

func (c *Collection[T]) Before(item T) (*T, error) {
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

func (c *Collection[T]) Chunk(size int) ([]Collection[T], error) {

	chunks := make([]Collection[T], 0)
	items := c.All()

	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	for i := 0; i < len(items); i += size {
		end := i + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkBy(fn func(T) bool) ([]Collection[T], error) {
	items := c.All()
	chunks := make([]Collection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, Init(items[i]))
		}
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkUntil(fn func(T) bool) ([]Collection[T], error) {
	items := c.All()
	chunks := make([]Collection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, Init(items[i]))
		}
	}
	return chunks, nil
}

func (c *Collection[T]) ChunkWhile(fn func(T) bool) ([]Collection[T], error) {
	items := c.All()
	chunks := make([]Collection[T], 0)
	for i := range items {
		if fn(items[i]) {
			chunks = append(chunks, Init(items[i]))
		}
	}
	return chunks, nil
}

func (c *Collection[T]) Contains(item T) bool {
	return slices.Contains(c.All(), item)
}

func (c *Collection[T]) Count() int {
	return len(c.All())
}

func (c *Collection[T]) CountBy(fn func(T) bool) int {
	count := 0
	for _, v := range c.All() {
		if fn(v) {
			count++
		}
	}
	return count
}

func (c *Collection[T]) Counts() map[T]int {
	counts := make(map[T]int)
	for _, v := range c.All() {
		counts[v]++
	}
	return counts
}

// O(n**2) time complexity
func (c *Collection[T]) Diff(other Collection[T]) []T {
	diff := make([]T, 0)
	items := c.All()
	for i := range items {
		if !slices.Contains(other.All(), items[i]) {
			diff = append(diff, items[i])
		}
	}
	return diff
}

func (c *Collection[T]) DoesntContain(item T) bool {
	return !slices.Contains(c.All(), item)
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
	for _, v := range c.All() {
		if !fn(v) {
			return false
		}
	}
	return true
}

func (c *Collection[T]) Filter(fn func(T) bool) []T {
	items := c.All()
	filtered := make([]T, 0)
	for i := range items {
		if fn(items[i]) {
			filtered = append(filtered, items[i])
		}
	}
	return filtered
}

func (c *Collection[T]) First() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[0], nil
}

func (c *Collection[T]) Has(item T) bool {
	return c.Contains(item)
}

func (c *Collection[T]) HasAny(items ...T) bool {
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

func (c *Collection[T]) HasAll(items ...T) bool {
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
func (c *Collection[T]) Intersect(other Collection[T]) Collection[T] {
	items := c.All()
	intersect := make([]T, 0)
	for i := range items {
		if slices.Contains(other.All(), items[i]) {
			intersect = append(intersect, items[i])
		}
	}
	return Collection[T](intersect)
}

func (c *Collection[T]) Last() (*T, error) {
	items := c.All()
	if len(items) == 0 {
		return nil, errors.New("empty collection")
	}
	return &items[len(items)-1], nil
}

func (c *Collection[T]) Length() int {
	return c.Count()
}

func (c *Collection[T]) Map(fn func(T) T) *Collection[T] {
	items := c.All()
	mapped := make([]T, 0)
	for i := range items {
		mapped = append(mapped, fn(items[i]))
	}

	*c = mapped
	return c
}

func (c *Collection[T]) Mode() *T {
	items := c.Counts()

	var mode *T
	highest := -1

	for k, v := range items {
		if v > highest {
			*mode = k
		}
	}
	return mode
}

func (c *Collection[T]) Multiply(multiplier int) *Collection[T] {
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

func (c *Collection[T]) Nth(n int) (*T, error) {
	items := c.All()
	if n < 0 {
		return nil, errors.New("n must be positive")
	}

	if n >= len(items) {
		return nil, errors.New("n is out of range")
	}
	return &items[n], nil
}

func (c *Collection[T]) PadLeft(length int, value T) *Collection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append(newSlice, c.All()...)

	*c = appended

	return c
}

func (c *Collection[T]) PadRight(length int, value T) *Collection[T] {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append([]T{}, c.All()...)
	appended = append(appended, newSlice...)

	*c = appended

	return c
}

func (c *Collection[T]) Prepend(items ...T) *Collection[T] {
	newSlice := []T{}
	newSlice = append(newSlice, items...)
	newSlice = append(newSlice, c.All()...)
	*c = newSlice
	return c
}

func (c *Collection[T]) Pop() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	ret := c.At(c.Length() - 1)

	*c = c.All()[:c.Length()-1]
	return ret, nil
}

func (c *Collection[T]) Push(item ...T) *Collection[T] {
	*c = append(c.All(), item...)
	return c
}

func (c *Collection[T]) Random() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}
	return c.At(rand.Intn(c.Length())), nil
}

func (c *Collection[T]) Range(start int, end int) *Collection[T] {
	collection := InitFromSlice(c.All()[start:end])

	return &collection
}

func (c *Collection[T]) Reject(fn func(T) bool) *Collection[T] {
	rejected := make([]T, 0)
	for _, v := range c.All() {
		if !fn(v) {
			rejected = append(rejected, v)
		}
	}

	collection := InitFromSlice(rejected)

	return &collection
}

func (c *Collection[T]) Reverse() *Collection[T] {
	reversed := make([]T, 0)
	for i := c.Length() - 1; i >= 0; i-- {
		reversed = append(reversed, *c.At(i))
	}

	collection := InitFromSlice(reversed)

	return &collection
}

func (c *Collection[T]) Search(item T) int {
	return slices.Index(c.All(), item)
}

func (c *Collection[T]) Shift() (*T, error) {
	if c.Length() == 0 {
		return nil, errors.New("empty collection")
	}

	ret := c.At(0)
	*c = c.All()[1:]

	return ret, nil
}

func (c *Collection[T]) Shuffle() *Collection[T] {
	shuffled := make([]T, 0)
	for i := range c.All() {
		shuffled = append(shuffled, *c.At(i))
	}

	collection := InitFromSlice(shuffled)

	return &collection
}

func (c *Collection[T]) Skip(n int) *Collection[T] {
	collection := InitFromSlice(c.All()[n:])
	return &collection
}

func (c *Collection[T]) SkipUntil(fn func(T) bool) *Collection[T] {
	for i := range c.All() {
		if fn(*c.At(i)) {
			collection := InitFromSlice(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *Collection[T]) SkipWhile(fn func(T) bool) *Collection[T] {
	for i := range c.All() {
		if !fn(*c.At(i)) {
			collection := InitFromSlice(c.All()[i:])
			return &collection
		}
	}
	return c
}

func (c *Collection[T]) Slice(start int, end int) *Collection[T] {
	collection := InitFromSlice(c.All()[start:end])
	return &collection
}

func (c *Collection[T]) Splice(start int, take int, items ...T) *Collection[T] {
	collection := InitFromSlice(append(c.All()[:start], append(items, c.All()[start+take:]...)...))
	return &collection
}

func (c *Collection[T]) Split(n int) []Collection[T] {

	chunks := make([]Collection[T], 0)

	for i := 0; i < c.Length()/n; i += n {
		chunks = append(chunks, InitFromSlice(c.All()[i:i+n]))
	}
	return chunks
}

func (c *Collection[T]) SplitInto(n int) []Collection[T] {
	return c.Split(c.Length() / n)
}

func (c *Collection[T]) Take(n int) *Collection[T] {
	collection := InitFromSlice(c.All()[:n])
	return &collection
}

func (c *Collection[T]) TakeUntil(fn func(T) bool) *Collection[T] {
	for _, v := range c.All() {
		if fn(v) {
			collection := Init(v)
			return &collection
		}
	}
	return c
}

func (c *Collection[T]) TakeWhile(fn func(T) bool) *Collection[T] {
	for _, v := range c.All() {
		if !fn(v) {
			collection := Init(v)
			return &collection
		}
	}
	return c
}

func (c *Collection[T]) Union(other Collection[T]) *Collection[T] {
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

	collection := InitFromSlice(union)

	return &collection
}

func (c *Collection[T]) Unique() *Collection[T] {
	counts := c.Counts()
	unique := make([]T, 0)
	for k, v := range counts {
		if v == 1 {
			unique = append(unique, k)
		}
	}

	collection := InitFromSlice(unique)
	return &collection
}
