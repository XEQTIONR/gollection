package gollection

import (
	"cmp"
	"errors"
	"math/rand"
	"slices"
	"sort"
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

func Average[T Numeric](c []T) T {
	if len(c) == 0 {
		return 0
	}

	sum := T(0)
	for _, item := range c {
		sum += item
	}

	return sum / T(len(c))
}

func Avg[T Numeric](c []T) T {
	return Average(c)
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

func Collapse[T any](c [][]T) []T {
	collapsed := make([]T, 0)
	for _, item := range c {
		collapsed = append(collapsed, item...)
	}
	return collapsed
}

func CollapseMap[T comparable, R any](c []map[T]R) map[T]R {
	collapsed := make(map[T]R)
	for _, item := range c {
		for k, v := range item {
			collapsed[k] = v
		}
	}
	return collapsed
}

func Combine[T comparable, R any](c []T, other []R) map[T]R {
	combined := make(map[T]R)
	for i := range len(c) {
		combined[c[i]] = other[i]
	}
	return combined
}

func CombineSlice[T comparable, R any](k []T, v []R) map[T]R {
	combined := make(map[T]R)
	for i := range len(k) {
		combined[k[i]] = v[i]
	}
	return combined
}

func CombineMap[T comparable, R any](c map[T]R, other map[T]R) map[T]R {
	combined := make(map[T]R)
	for k, v := range c {
		combined[k] = v
	}
	for k, v := range other {
		combined[k] = v
	}
	return combined
}

func Contains[T comparable](c []T, item T) bool {
	return slices.Contains(c, item)
}

func ContainsKey[T comparable, R any](c map[T]R, key T) bool {
	_, ok := c[key]
	return ok
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

func CrossJoin[T any, R any](c []T, other []R) [][]any {
	crossJoined := make([][]any, 0)
	for _, v := range c {
		for _, v2 := range other {
			crossJoined = append(crossJoined, []any{v, v2})
		}
	}
	return crossJoined
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

// TODO: Implement this
func DiffAssoc[T comparable, R any](c []T, other []R) []map[T]R {
	diffAssoc := make([]map[T]R, 0)
	for _, v := range c {
		for _, v2 := range other {
			diffAssoc = append(diffAssoc, map[T]R{v: v2})
		}
	}
	return diffAssoc
}

func DiffKeys[T comparable, R any](c map[T]R, other map[T]R) []T {
	diffKeys := make([]T, 0)
	for k := range c {
		if _, ok := other[k]; !ok {
			diffKeys = append(diffKeys, k)
		}
	}
	return diffKeys
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

func Except[T any](c []T, fn func(T) bool) []T {
	filtered := make([]T, 0)

	for _, item := range c {
		if !fn(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
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

func First[T any](c []T, fn func(T, int) bool) *T {
	for i, v := range c {
		if fn(v, i) {
			return &v
		}
	}

	return nil
}

func FirstOr[T any](c []T, fn func(T, int) bool, defaultValue T) *T {
	for i, v := range c {
		if fn(v, i) {
			return &v
		}
	}

	return &defaultValue
}

func FirstOrFail[T any](c []T, fn func(T, int) bool) *T {
	for i, v := range c {
		if fn(v, i) {
			return &v
		}
	}

	panic("no item found")
}

func FirstWhere[T comparable](c []map[T]any, key T, value any) map[T]any {
	for _, v := range c {
		if v[key] == value {
			return v
		}
	}

	return nil
}

func Flatten[T any](c [][]T) []T {
	return Collapse(c)
}

func FlattenMap[T comparable, R any](c []map[T]R) map[T]R {
	return CollapseMap(c)
}

func Flip[T comparable, R comparable](c map[T]R) map[R]T {
	flipped := make(map[R]T)
	for k, v := range c {
		flipped[v] = k
	}
	return flipped
}

// mutates the original map
func Forget[T comparable, R any](c map[T]R, keys ...T) map[T]R {
	for _, key := range keys {
		delete(c, key)
	}
	return c
}

func GroupBy[T any, R comparable](c []T, fn func(T) R) map[R][]T {
	grouped := make(map[R][]T)
	for _, item := range c {
		grouped[fn(item)] = append(grouped[fn(item)], item)
	}
	return grouped
}

func GroupByKeyValue[T comparable, R comparable](c []map[T]R, key T) map[R][]map[T]R {
	grouped := make(map[R][]map[T]R)
	for _, obj := range c {
		if value, ok := obj[key]; ok {
			grouped[value] = append(grouped[value], obj)
		}
	}
	return grouped
}

func HasAnyKeys[T comparable, R any](c map[T]R, keys ...T) bool {
	for _, key := range keys {
		if _, ok := c[key]; ok {
			return true
		}
	}
	return false
}

// implment from json

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

// func HasOne [T comparable, R comparable](c []T) bool {

// }

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

func IntersectAssoc[T comparable, R comparable](c map[T]R, other map[T]R) map[T]R {
	intersectAssoc := make(map[T]R, 0)

	for k, v1 := range c {
		if v2, ok := other[k]; ok {
			if v1 == v2 {
				intersectAssoc[k] = v1
			}
		}
	}

	return intersectAssoc
}

func IntersectByKeys[T comparable, R any](c map[T]R, other map[T]R) map[T]R {
	intersectByKeys := make(map[T]R, 0)

	for k, v := range c {
		if _, ok := other[k]; ok {
			intersectByKeys[k] = v
		}
	}

	return intersectByKeys
}

func KeyBy[T comparable, R comparable](c []map[T]R, key T) map[R][]map[T]R {
	keyBy := make(map[R][]map[T]R, 0)
	for _, obj := range c {
		if value, ok := obj[key]; ok {
			keyBy[value] = append(keyBy[value], obj)
		}
	}
	return keyBy
}

func Last[T any](c []T, fn func(T, int) bool) *T {
	for i := len(c) - 1; i >= 0; i-- {
		if fn(c[i], i) {
			return &c[i]
		}
	}

	return nil
}

func LastOr[T any](c []T, fn func(T, int) bool, defaultValue T) *T {
	for i := len(c) - 1; i >= 0; i-- {
		if fn(c[i], i) {
			return &c[i]
		}
	}

	return &defaultValue
}

func LastOrFail[T any](c []T, fn func(T, int) bool) *T {
	for i := len(c) - 1; i >= 0; i-- {
		if fn(c[i], i) {
			return &c[i]
		}
	}

	panic("no item found")
}

func LastWhere[T comparable](c []map[T]any, key T, value any) map[T]any {
	for i := len(c) - 1; i >= 0; i-- {
		if c[i][key] == value {
			return c[i]
		}
	}

	return nil
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

func Only[T comparable, R any](c []map[T]R, keys ...T) []map[T]R {
	only := make([]map[T]R, 0)

	for _, obj := range c {
		newObj := make(map[T]R)
		for _, key := range keys {
			if value, ok := obj[key]; ok {
				newObj[key] = value
			}
		}
		only = append(only, newObj)
	}

	return only
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

func Partition[T any](c []T, fn func(T) bool) ([]T, []T) {
	partitioned := make([]T, 0)
	other := make([]T, 0)
	for _, v := range c {
		if fn(v) {
			partitioned = append(partitioned, v)
		} else {
			other = append(other, v)
		}
	}

	return partitioned, other
}

func Percentage[T any](c []T, fn func(T) bool) float64 {
	percentage := []bool{}
	for _, v := range c {
		if fn(v) {
			percentage = append(percentage, true)
		}
	}

	return (float64(len(percentage)) / float64(len(c))) * 100
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

func Reduce[T any, R any](c []T, fn func(R, T) R, initial R) R {
	reduced := initial
	for _, v := range c {
		reduced = fn(reduced, v)
	}
	return reduced
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

func Replace[T any](c []T, replacements map[int]T) []T {
	for i, _ := range c {
		if replacement, ok := replacements[i]; ok {
			c[i] = replacement
		}
	}

	return c
}

func Reverse[T any](c []T) []T {
	reversed := make([]T, 0)

	for i := len(c) - 1; i >= 0; i-- {
		reversed = append(reversed, c[i])
	}

	return reversed
}

func ReverseMap[T comparable, R comparable](c map[T]R) map[R]T {
	reversed := make(map[R]T)

	for k, v := range c {
		reversed[v] = k
	}

	return reversed
}

func Select[T comparable, R any](c []map[T]R, keys ...T) []map[T]R {
	selected := make([]map[T]R, 0)

	for _, obj := range c {
		newObj := make(map[T]R)

		for _, key := range keys {
			if value, ok := obj[key]; ok {
				newObj[key] = value
			}
		}

		selected = append(selected, newObj)
	}

	return selected
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

func Sort[T cmp.Ordered](c []T, direction SortDirection) []T {
	if direction == SortDirectionAsc {
		return SortAsc(c)
	}

	return SortDesc(c)
}

func SortAsc[T cmp.Ordered](c []T) []T {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	return c
}

func SortBy[T comparable, R cmp.Ordered](c []map[T]R, field T, direction SortDirection) []map[T]R {
	if direction == SortDirectionAsc {
		return SortByAsc(c, field)
	}

	return SortByDesc(c, field)
}

func SortByAsc[T comparable, R cmp.Ordered](c []map[T]R, field T) []map[T]R {
	sort.Slice(c, func(i, j int) bool {
		return c[i][field] < c[j][field]
	})
	return c
}

func SortByDesc[T comparable, R cmp.Ordered](c []map[T]R, field T) []map[T]R {
	sort.Slice(c, func(i, j int) bool {
		return c[i][field] > c[j][field]
	})
	return c
}

func SortDesc[T cmp.Ordered](c []T) []T {
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	return c
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

func Where[T comparable, R comparable](c []map[T]R, key T, value R) []map[T]R {
	where := make([]map[T]R, 0)

	for _, obj := range c {
		if obj[key] == value {
			where = append(where, obj)
		}
	}
	return where
}

func WhereBetween[T comparable, R cmp.Ordered](c []map[T]R, key T, min R, max R) []map[T]R {
	whereBetween := make([]map[T]R, 0)
	for _, obj := range c {
		if obj[key] >= min && obj[key] <= max {
			whereBetween = append(whereBetween, obj)
		}
	}
	return whereBetween
}

func WhereIn[T comparable, R comparable](c []map[T]R, key T, values []R) []map[T]R {
	whereIn := make([]map[T]R, 0)
	valMap := make(map[R]bool)

	for _, value := range values {
		valMap[value] = true
	}

	for _, obj := range c {
		if valMap[obj[key]] {
			whereIn = append(whereIn, obj)
		}
	}
	return whereIn
}

func WhereNot[T comparable, R comparable](c []map[T]R, key T, value R) []map[T]R {
	whereNot := make([]map[T]R, 0)
	for _, obj := range c {
		if obj[key] != value {
			whereNot = append(whereNot, obj)
		}
	}
	return whereNot
}

func WhereNotBetween[T comparable, R cmp.Ordered](c []map[T]R, key T, min R, max R) []map[T]R {
	whereBetween := make([]map[T]R, 0)
	for _, obj := range c {
		if obj[key] < min && obj[key] > max {
			whereBetween = append(whereBetween, obj)
		}
	}
	return whereBetween
}

func WhereNotIn[T comparable, R comparable](c []map[T]R, key T, values []R) []map[T]R {
	whereNotIn := make([]map[T]R, 0)
	valMap := make(map[R]bool)

	for _, value := range values {
		valMap[value] = true
	}

	for _, obj := range c {
		if !valMap[obj[key]] {
			whereNotIn = append(whereNotIn, obj)
		}
	}
	return whereNotIn
}

func Zip[T any, R any](c []T, other []R) [][]any {
	zip := make([][]any, 0)
	for i := range c {
		zip = append(zip, []any{c[i], other[i]})
	}
	return zip
}
