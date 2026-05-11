package gollection

import (
	"cmp"
	"errors"
	"math/rand"
	"slices"
)

func Append[T any](c []T, items ...T) []T {
	return append(c, items...)
}

func At[T any](c []T, index int) (*T, error) {
	if index < 0 || index >= len(c) {
		return nil, errors.New("Index out of bounds")
	}

	return &c[index], nil
}

func Chunk[T any](c []T, size int) ([][]T, error) {
	chunks := make([][]T, 0)

	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	for i := 0; i < len(c); i += size {
		end := i + size
		if end > len(c) {
			end = len(c)
		}
		chunks = append(chunks, []T(c[i:end]))
	}

	return chunks, nil
}

func ChunkBy[T any](c []T, fn func(T) bool) ([][]T, error) {
	chunks := make([][]T, 0)

	var chunk []T

	for _, item := range c {
		if fn(item) {
			if len(chunk) > 0 {
				chunks = append(chunks, chunk)
			}
			chunk = []T{item}
		} else {
			chunk = append(chunk, item)
		}
	}

	return chunks, nil
}

func Contains[T comparable](c []T, item T) bool {
	return slices.Contains(c, item)
}

func Count[T any](c []T) int {
	return len(c)
}

func CountBy[T any](c []T, fn func(T) bool) int {
	count := 0

	for _, item := range c {
		if fn(item) {
			count++
		}
	}

	return count
}

func Counts[T comparable](c []T) map[T]int {
	counts := make(map[T]int)

	for _, v := range c {
		counts[v]++
	}

	return counts
}

func Diff[T comparable](c []T, other []T) []T {
	items := make(map[T]bool, 0)

	for _, v := range other {
		items[v] = true
	}

	diff := make([]T, 0)
	for _, v := range c {
		if _, ok := items[v]; !ok {
			diff = append(diff, v)
		}
	}

	return diff
}

func DoesntContain[T comparable](c []T, item T) bool {
	return !slices.Contains(c, item)
}

func Duplicates[T comparable](c []T) []T {
	counts := Counts(c)
	duplicates := make([]T, 0)

	for k, v := range counts {
		if v > 1 {
			duplicates = append(duplicates, k)
		}
	}

	return duplicates
}

func Each[T any](c []T, fn func(T)) {
	for _, item := range c {
		fn(item)
	}
}

func Every[T any](c []T, fn func(T) bool) bool {
	for _, item := range c {
		if !fn(item) {
			return false
		}
	}

	return true
}

func Filter[T any](c []T, fn func(T) bool) []T {
	filtered := make([]T, 0)

	for _, item := range c {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func First[T any](c []T) *T {
	if len(c) == 0 {
		return nil
	}

	return &c[0]
}

func IndexOf[T comparable](c []T, item T) int {
	return slices.Index(c, item)
}

func IsEmpty[T any](c []T) bool {
	return len(c) == 0
}

func Has[T comparable](c []T, item T) bool {
	return Contains(c, item)
}

func HasAny[T comparable](c []T, items []T) bool {
	m := make(map[T]bool)

	for _, item := range items {
		m[item] = true
	}

	for _, item := range c {
		if m[item] {
			return true
		}
	}

	return false
}

func HasEvery[T comparable](c []T, items []T) bool {
	m := make(map[T]bool)

	for _, item := range items {
		m[item] = true
	}

	for _, item := range c {
		if !m[item] {
			return false
		}
	}

	return true
}

func Intersect[T comparable](c []T, other []T) []T {
	m := make(map[T]int)
	intersect := make([]T, 0)

	for _, item := range c {
		m[item]++
	}

	for _, item := range other {
		if m[item] > 0 {
			intersect = append(intersect, item)
			m[item]--
		}
	}

	return intersect
}

func Last[T any](c []T) *T {
	if len(c) == 0 {
		return nil
	}

	return &c[len(c)-1]
}

func Length[T any](c []T) int {
	return Count(c)
}

func Map[T any, R any](c []T, fn func(T) R) []R {
	mapped := make([]R, 0)
	for _, item := range c {
		mapped = append(mapped, fn(item))
	}

	return mapped
}

func Max[T cmp.Ordered](c []T) *T {
	if len(c) == 0 {
		return nil
	}
	max := c[0]
	for _, item := range c {
		if item > max {
			max = item
		}
	}

	return &max
}

func Min[T cmp.Ordered](c []T) *T {
	if len(c) == 0 {
		return nil
	}
	max := c[0]
	for _, item := range c {
		if item > max {
			max = item
		}
	}

	return &max
}

func Mode[T cmp.Ordered](c []T) *T {
	if len(c) == 0 {
		return nil
	}
	items := Counts(c)

	var mode *T
	highest := -1

	for k, v := range items {
		if v > highest {
			*mode = k
		}
	}

	return mode
}

func Multiply[T any](c []T, multiplier int) []T {
	if Length(c) == 0 {
		return c
	}

	for i := Length(c) - 1; i >= 0; i-- {
		repeat := slices.Repeat([]T{c[i]}, multiplier-1)
		c = slices.Insert(c, i, repeat...)
	}

	return c
}

func Nth[T any](c []T, index int) *T {
	if index < 0 || index >= len(c) {
		return nil
	}

	return &c[index]
}

func NthFromLast[T any](c []T, index int) *T {
	if index < 0 || index >= len(c) {
		return nil
	}

	return &c[len(c)-index-1]
}

func PadLeft[T any](c []T, length int, value T) []T {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append(newSlice, c...)
	c = appended

	return c
}

func PadRight[T any](c []T, length int, value T) []T {
	if length <= 0 {
		return c
	}

	newSlice := slices.Repeat([]T{value}, length)
	appended := append([]T{}, c...)
	appended = append(appended, newSlice...)
	c = appended

	return c
}

func Prepend[T any](c []T, items ...T) []T {
	newSlice := []T{}
	newSlice = append(newSlice, items...)
	newSlice = append(newSlice, c...)
	c = newSlice

	return c
}

func Push[T any](c []T, item T) []T {
	return append(c, item)
}

func Pop[T any](c []T) (*T, []T) {
	if len(c) == 0 {
		return nil, c
	}

	return &c[len(c)-1], c[:len(c)-1]
}

func Random[T any](c []T) (*T, error) {
	if Length(c) == 0 {
		return nil, errors.New("empty collection")
	}

	return At(c, rand.Intn(Length(c)))
}

func Range[T any](c []T, start int, end int) ([]T, error) {
	if start < 0 || end > Length(c) {
		return nil, errors.New("start and end must be within the range of the collection")
	}

	if start > end {
		return nil, errors.New("start must be less than end")
	}

	return c[start:end], nil
}

func Reject[T any](c []T, fn func(T) bool) []T {
	rejected := make([]T, 0)

	for _, v := range c {
		if !fn(v) {
			rejected = append(rejected, v)
		}
	}

	return rejected
}

func Reverse[T any](c []T) []T {
	reversed := make([]T, 0)

	for i := len(c) - 1; i >= 0; i-- {
		reversed = append(reversed, c[i])
	}

	return reversed
}

func Shuffle[T any](c []T) []T {
	shuffled := make([]T, 0)
	options := []int{}

	for i := range Length(c) {
		options = append(options, i)
	}

	for i := 0; i < Length(c); i++ {
		randomIndex := rand.Intn(len(options))
		shuffled = append(shuffled, c[options[randomIndex]])
		options = append(options[:randomIndex], options[randomIndex+1:]...)
	}

	return shuffled
}

func Skip[T any](c []T, n int) []T {
	return c[n:]
}

func SkipUntil[T any](c []T, fn func(T) bool) []T {
	for i := range c {
		if fn(c[i]) {
			return c[i:]
		}
	}

	return c
}

func SkipWhile[T any](c []T, fn func(T) bool) []T {
	for i := range c {
		if !fn(c[i]) {
			return c[i:]
		}
	}

	return c
}

func Slice[T any](c []T, start int, end int) []T {
	return c[start:end]
}

func Splice[T any](c []T, start int, take int, items ...T) []T {
	return append(c[:start], append(items, c[start+take:]...)...)
}

func Split[T any](c []T, n int) [][]T {
	chunks := make([][]T, 0)

	for i := 0; i < Length(c)/n; i += n {
		chunks = append(chunks, c[i:i+n])
	}

	return chunks
}

func SplitInto[T any](c []T, n int) [][]T {
	return Split(c, Length(c)/n)
}

func Take[T any](c []T, n int) []T {
	return c[:n]
}

func TakeUntil[T any](c []T, fn func(T) bool) []T {
	for i, v := range c {
		if fn(v) {
			return c[:i]
		}
	}

	return c
}

func TakeWhile[T any](c []T, fn func(T) bool) []T {
	for i, v := range c {
		if !fn(v) {
			return c[:i]
		}
	}

	return c
}

func Union[T comparable](c []T, other []T) []T {

	m := make(map[T]bool)

	for _, v := range c {
		m[v] = true
	}
	for _, v := range other {
		m[v] = true
	}

	union := make([]T, 0, len(m))
	for v := range m {
		union = append(union, v)
	}

	return union
}

func Unique[T comparable](c []T) []T {
	counts := Counts(c)
	uniques := make([]T, 0)
	for k, v := range counts {
		if v == 1 {
			uniques = append(uniques, k)
		}
	}

	return uniques
}
