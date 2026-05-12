package gollection

import (
	"math"
	"reflect"
	"slices"
	"sort"
	"testing"
)

func sortedInts(s []int) []int {
	out := slices.Clone(s)
	sort.Ints(out)
	return out
}

func mapsEqual[K comparable, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func TestAppend(t *testing.T) {
	got := Append([]int{1}, 2, 3)
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestAt(t *testing.T) {
	v, err := At([]string{"a", "b"}, 1)
	if err != nil || *v != "b" {
		t.Fatalf("got %v %v", v, err)
	}
	_, err = At([]int{1}, -1)
	if err == nil {
		t.Fatal("expected error")
	}
	_, err = At([]int{1}, 1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAverage(t *testing.T) {
	if got := Average([]int{1, 2, 3, 4}); got != 2 {
		t.Fatalf("got %d", got)
	}
	if got := Average([]int{}); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestAvg(t *testing.T) {
	if got := Avg([]float64{2, 4}); got != 3 {
		t.Fatalf("got %v", got)
	}
}

func TestChunk(t *testing.T) {
	ch, err := Chunk([]int{1, 2, 3, 4}, 2)
	if err != nil || len(ch) != 2 || !slices.Equal(ch[0], []int{1, 2}) || !slices.Equal(ch[1], []int{3, 4}) {
		t.Fatalf("got %v %v", ch, err)
	}
	odd, err := Chunk([]int{1, 2, 3}, 2)
	if err != nil || len(odd) != 2 || !slices.Equal(odd[1], []int{3}) {
		t.Fatalf("got %v %v", odd, err)
	}
	_, err = Chunk([]int{1}, 0)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestChunkBy(t *testing.T) {
	ch, err := ChunkBy([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 1 })
	if err != nil {
		t.Fatal(err)
	}
	if len(ch) != 2 {
		t.Fatalf("got %v", ch)
	}
}

func TestCollapse(t *testing.T) {
	got := Collapse([][]int{{1, 2}, {3}})
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestCollapseMap(t *testing.T) {
	got := CollapseMap([]map[string]int{{"a": 1}, {"b": 2, "a": 3}})
	if got["a"] != 3 || got["b"] != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestCombine(t *testing.T) {
	got := Combine([]string{"a", "b"}, []int{1, 2})
	if !mapsEqual(got, map[string]int{"a": 1, "b": 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestCombineSlice(t *testing.T) {
	got := CombineSlice([]rune{'x', 'y'}, []int{9, 8})
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestCombineMap(t *testing.T) {
	got := CombineMap(map[string]int{"a": 1}, map[string]int{"b": 2, "a": 3})
	if got["a"] != 3 || got["b"] != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2}, 2) || Contains([]int{1}, 9) {
		t.Fatal()
	}
}

func TestContainsKey(t *testing.T) {
	if !ContainsKey(map[string]int{"k": 1}, "k") {
		t.Fatal("expected key present")
	}
	if ContainsKey(map[string]int{}, "k") {
		t.Fatal("expected missing key")
	}
}

func TestCount(t *testing.T) {
	if Count([]int{1, 2, 3}) != 3 {
		t.Fatal()
	}
}

func TestCountBy(t *testing.T) {
	n := CountBy([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
	if n != 2 {
		t.Fatalf("got %d", n)
	}
}

func TestCounts(t *testing.T) {
	got := Counts([]string{"a", "a", "b"})
	if got["a"] != 2 || got["b"] != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestCrossJoin(t *testing.T) {
	got := CrossJoin([]int{1}, []string{"x"})
	if len(got) != 1 || !reflect.DeepEqual(got[0], []any{1, "x"}) {
		t.Fatalf("got %v", got)
	}
}

func TestDiff(t *testing.T) {
	got := Diff([]int{1, 2, 3}, []int{2})
	if !slices.Equal(sortedInts(got), []int{1, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestDiffAssoc(t *testing.T) {
	got := DiffAssoc(map[string]int{"a": 1, "b": 3}, map[string]int{"a": 1, "b": 2})
	if len(got) != 1 || !mapsEqual(got[0], map[string]int{"b": 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestDiffKeys(t *testing.T) {
	got := DiffKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3})
	if len(got) != 1 || got[0] != "a" {
		t.Fatalf("got %v", got)
	}
}

func TestDoesntContain(t *testing.T) {
	if !DoesntContain([]int{1}, 2) || DoesntContain([]int{1}, 1) {
		t.Fatal()
	}
}

func TestDuplicates(t *testing.T) {
	got := Duplicates([]int{1, 1, 2, 3, 3})
	sort.Ints(got)
	if !slices.Equal(got, []int{1, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestEach(t *testing.T) {
	var sum int
	Each([]int{1, 2, 3}, func(n int) { sum += n })
	if sum != 6 {
		t.Fatalf("got %d", sum)
	}
}

func TestEvery(t *testing.T) {
	if !Every([]int{2, 4}, func(n int) bool { return n%2 == 0 }) {
		t.Fatal()
	}
	if Every([]int{2, 3}, func(n int) bool { return n%2 == 0 }) {
		t.Fatal()
	}
}

func TestExcept(t *testing.T) {
	got := Except([]int{1, 2, 3}, func(n int) bool { return n == 2 })
	if !slices.Equal(got, []int{1, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1, 2, 3}, func(n int) bool { return n > 1 })
	if !slices.Equal(got, []int{2, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestFirst(t *testing.T) {
	v := First([]int{1, 2, 3}, func(n, i int) bool { return n == 2 })
	if v == nil || *v != 2 {
		t.Fatalf("got %v", v)
	}
	if First([]int{1}, func(int, int) bool { return false }) != nil {
		t.Fatal()
	}
}

func TestFirstOr(t *testing.T) {
	d := 99
	v := FirstOr([]int{1, 2}, func(int, int) bool { return false }, d)
	if v == nil || *v != 99 {
		t.Fatalf("got %v", v)
	}
	v = FirstOr([]int{1, 2}, func(n, _ int) bool { return n == 2 }, d)
	if v == nil || *v != 2 {
		t.Fatalf("got %v", v)
	}
}

func TestFirstOrFail(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic")
			}
		}()
		FirstOrFail([]int{1}, func(int, int) bool { return false })
	})
	t.Run("found", func(t *testing.T) {
		v := FirstOrFail([]int{1, 2}, func(n, _ int) bool { return n == 2 })
		if v == nil || *v != 2 {
			t.Fatalf("got %v", v)
		}
	})
}

func TestFirstWhere(t *testing.T) {
	rows := []map[string]any{{"id": 1, "size": "S"}, {"id": 2, "size": "L"}}
	got := FirstWhere(rows, "id", 2)
	if got == nil || got["id"] != 2 || got["size"] != "L" {
		t.Fatalf("got %v", got)
	}
	if FirstWhere([]map[string]any{}, "id", 1) != nil {
		t.Fatal()
	}
}

func TestFlatten(t *testing.T) {
	if !slices.Equal(Flatten([][]int{{1}, {2, 3}}), []int{1, 2, 3}) {
		t.Fatal()
	}
}

func TestFlattenMap(t *testing.T) {
	got := FlattenMap([]map[string]int{{"a": 1}, {"b": 2, "a": 3}, {"a": 5, "c": 4}})
	if got["a"] != 5 || got["b"] != 2 || got["c"] != 4 {
		t.Fatal()
	}
}

func TestFlip(t *testing.T) {
	got := Flip(map[int]string{1: "a", 2: "b"})
	if got["a"] != 1 || got["b"] != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestForget(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	Forget(m, "a", "c")
	if _, ok := m["a"]; ok {
		t.Fatalf("got %v", m)
	}

	if _, ok := m["c"]; ok {
		t.Fatalf("got %v", m)
	}
	if m["b"] != 2 {
		t.Fatalf("got %v", m)
	}
}
func TestGroupBy(t *testing.T) {
	got := GroupBy([]map[string]string{{"t": "a"}, {"t": "b"}, {"t": "a"}}, "t")
	if len(got["a"]) != 2 || len(got["b"]) != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestHas(t *testing.T) {
	if !Has([]int{1}, 1) {
		t.Fatal()
	}
}

func TestHasAny(t *testing.T) {
	if !HasAny([]int{1, 9}, []int{9, 8}) {
		t.Fatal()
	}
	if HasAny([]int{1}, []int{2}) {
		t.Fatal()
	}
}

func TestHasAnyKeys(t *testing.T) {
	if !HasAnyKeys(map[string]int{"a": 1}, "a", "b") {
		t.Fatal("expected any key hit")
	}
	if HasAnyKeys(map[string]int{}, "x") {
		t.Fatal("expected no keys")
	}
}

func TestHasEvery(t *testing.T) {
	// Implementation: every element of c must be present in items.
	if !HasEvery([]int{1, 2}, []int{1, 2, 3}) {
		t.Fatal()
	}
	if HasEvery([]int{1, 4}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func TestHasOne(t *testing.T) {
	// Implementation: every element of c must be present in items.
	if !HasOne([]int{1, 4, 5}, []int{1, 2, 3}) {
		t.Fatal()
	}
	if HasOne([]int{1, 2, 5}, []int{1, 2, 3}) {
		t.Fatal()
	}
}

func TestIndexOf(t *testing.T) {
	if IndexOf([]string{"a", "b", "c"}, "b") != 1 {
		t.Fatal()
	}

	if IndexOf([]string{"a", "b", "c"}, "d") != -1 {
		t.Fatalf("got %d", IndexOf([]string{"a", "b", "c"}, "d"))
	}
}

func TestIsEmpty(t *testing.T) {
	if !IsEmpty([]int{}) || IsEmpty([]int{1}) {
		t.Fatal()
	}
}

func TestIntersect(t *testing.T) {
	// Intersection is built by walking `other` and consuming counts from `c`.
	got := Intersect([]int{1, 1, 2}, []int{1, 2, 2})
	sort.Ints(got)
	if !slices.Equal(got, []int{1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestIntersectAssoc(t *testing.T) {
	got := IntersectAssoc(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 9})
	if len(got) != 1 || got["a"] != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestIntersectByKeys(t *testing.T) {
	got := IntersectByKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 9})
	if got["b"] != 2 || len(got) != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestKeyBy(t *testing.T) {
	got := KeyBy([]map[string]int{{"k": 1}, {"k": 2, "a": 3}, {"k": 2, "b": 4}}, "k")
	if len(got[1]) != 1 || len(got[2]) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestLast(t *testing.T) {
	v := Last([]int{1, 2, 3, 4}, func(n, i int) bool { return n >= 2 })
	if v == nil || *v != 4 {
		t.Fatalf("got %v", v)
	}
	if Last([]int{1}, func(int, int) bool { return false }) != nil {
		t.Fatal("expected nil")
	}
}

func TestLastOr(t *testing.T) {
	v := LastOr([]int{1}, func(int, int) bool { return false }, 0)
	if v == nil || *v != 0 {
		t.Fatalf("got %v", v)
	}
	v = LastOr([]int{1, 2, 1}, func(n, _ int) bool { return n == 1 }, 0)
	if v == nil || *v != 1 {
		t.Fatalf("got %v", v)
	}
}

func TestLastOrFail(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic")
			}
		}()
		LastOrFail([]int{1}, func(int, int) bool { return false })
	})
	t.Run("found", func(t *testing.T) {
		v := LastOrFail([]int{1, 2}, func(n, _ int) bool { return n == 1 })
		if v == nil || *v != 1 {
			t.Fatalf("got %v", v)
		}
	})
}

func TestLastWhere(t *testing.T) {
	rows := []map[string]any{{"n": 1}, {"n": 2}, {"n": 2}}
	got := LastWhere(rows, "n", 2)
	if got["n"] != 2 {
		t.Fatalf("got %v", got)
	}
	if LastWhere([]map[string]any{{"n": 1}}, "n", 9) != nil {
		t.Fatal("expected nil")
	}
}

func TestLength(t *testing.T) {
	if Length([]int{1, 2}) != 2 {
		t.Fatal()
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2}, func(n int) string { return string(rune('0' + n)) })
	if !slices.Equal(got, []string{"1", "2"}) {
		t.Fatalf("got %v", got)
	}
}

func TestMax(t *testing.T) {
	if *Max([]int{1, 5, 2}) != 5 {
		t.Fatal()
	}
	if Max([]int{}) != nil {
		t.Fatal()
	}
}
func TestMode(t *testing.T) {
	if *Mode([]int{1, 2, 1, 2, 3, 2}) != 2 {
		t.Fatal()
	}
	if Mode([]int{}) != nil {
		t.Fatal()
	}
}

func TestMin(t *testing.T) {
	if *Min([]int{3, 1, 2}) != 1 {
		t.Fatal()
	}
	if Min([]int{}) != nil {
		t.Fatal()
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply([]int{1, 2}, 3)
	if len(got) != 6 {
		t.Fatalf("got %v len %d", got, len(got))
	}
	if slices.Equal(Multiply([]int{}, 3), []int{}) == false {
		t.Fatal()
	}
}

func TestNth(t *testing.T) {
	if *Nth([]string{"a", "b"}, 1) != "b" {
		t.Fatal()
	}
	if Nth([]int{1}, 3) != nil {
		t.Fatal()
	}
}

func TestNthFromLast(t *testing.T) {
	if *NthFromLast([]int{10, 20, 30}, 0) != 30 {
		t.Fatal()
	}
	if NthFromLast([]int{1}, 5) != nil {
		t.Fatal()
	}
}

func TestOnly(t *testing.T) {
	got := Only([]map[string]int{{"a": 1, "b": 2}, {"a": 3, "b": 4}}, "a")
	if len(got) != 2 || got[0]["a"] != 1 || got[1]["a"] != 3 || len(got[0]) != 1 || len(got[1]) != 1 {
		t.Fatalf("got %v", got)
	}
}

func TestPadLeft(t *testing.T) {
	got := PadLeft([]int{1}, 2, 0)
	if !slices.Equal(got, []int{0, 0, 1}) {
		t.Fatalf("got %v", got)
	}
	if !slices.Equal(PadLeft([]int{1}, 0, 0), []int{1}) {
		t.Fatal()
	}
}

func TestPadRight(t *testing.T) {
	got := PadRight([]int{1}, 1, 9)
	if !slices.Equal(got, []int{1, 9}) {
		t.Fatalf("got %v", got)
	}
	if !slices.Equal(PadRight([]int{1}, 0, 9), []int{1}) {
		t.Fatal()
	}
}

func TestPartition(t *testing.T) {
	a, b := Partition([]int{1, 2, 3, 4}, func(i, n int) bool { return n%2 == 0 })
	if !slices.Equal(sortedInts(a), []int{2, 4}) || !slices.Equal(sortedInts(b), []int{1, 3}) {
		t.Fatalf("got %v %v", a, b)
	}

	a, b = Partition([]int{1, 2, 3, 4}, func(i, n int) bool { return i < 2 })
	if !slices.Equal(sortedInts(a), []int{1, 2}) || !slices.Equal(sortedInts(b), []int{3, 4}) {
		t.Fatalf("got %v %v", a, b)
	}
}

func TestPercentage(t *testing.T) {
	p := Percentage([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	if p != 50 {
		t.Fatalf("got %v", p)
	}
	p = Percentage([]int{}, func(int) bool { return true })
	if !math.IsNaN(p) {
		t.Fatalf("got %v", p)
	}
}

func TestPrepend(t *testing.T) {
	got := Prepend([]int{3}, 1, 2)
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestPush(t *testing.T) {
	if !slices.Equal(Push([]int{1}, 2), []int{1, 2}) {
		t.Fatal()
	}
}

func TestPop(t *testing.T) {
	last, rest := Pop([]int{1, 2})
	if *last != 2 || !slices.Equal(rest, []int{1}) {
		t.Fatalf("got %v %v", last, rest)
	}
	l, r := Pop([]int{})
	if l != nil || len(r) != 0 {
		t.Fatal()
	}
}

func TestRandom(t *testing.T) {
	_, err := Random([]int{})
	if err == nil {
		t.Fatal("expected error")
	}
	v, err := Random([]int{7, 8, 9})
	if err != nil {
		t.Fatal(err)
	}
	if *v != 7 && *v != 8 && *v != 9 {
		t.Fatalf("got %v", *v)
	}
}

func TestRange(t *testing.T) {
	got, err := Range([]int{0, 1, 2, 3}, 1, 3)
	if err != nil || !slices.Equal(got, []int{1, 2}) {
		t.Fatalf("got %v %v", got, err)
	}
	_, err = Range([]int{1}, 2, 0)
	if err == nil {
		t.Fatal("expected error")
	}
	_, err = Range([]int{1, 2}, -1, 1)
	if err == nil {
		t.Fatal("expected error for negative start")
	}
	_, err = Range([]int{1, 2}, 0, 10)
	if err == nil {
		t.Fatal("expected error for end past len")
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3}, func(acc, n int) int { return acc + n }, 0)
	if sum != 6 {
		t.Fatal()
	}
}

func TestReduceWithIndex(t *testing.T) {
	sum := ReduceWithIndex([]int{1, 2, 3}, func(acc, n, i int) int { return acc + i }, 0)
	if sum != 3 {
		t.Fatal()
	}
}

func TestReject(t *testing.T) {
	got := Reject([]int{1, 2, 3}, func(n int) bool { return n == 2 })
	if !slices.Equal(got, []int{1, 3}) {
		t.Fatalf("got %v", got)
	}
}

func TestRejectWithIndex(t *testing.T) {
	got := RejectWithIndex([]int{1, 2, 3}, func(n, i int) bool { return i == 2 })
	if !slices.Equal(got, []int{1, 2}) {
		t.Fatalf("got %v", got)
	}
}
func TestReplace(t *testing.T) {
	c := []int{1, 2, 3, 4, 5}
	got := Replace(c, map[int]int{1: 9, 3: 10})
	if !slices.Equal(got, []int{1, 9, 3, 10, 5}) {
		t.Fatalf("got %v", got)
	}
}

func TestReverse(t *testing.T) {
	if !slices.Equal(Reverse([]int{1, 2, 3}), []int{3, 2, 1}) {
		t.Fatal()
	}
}

func TestReverseMap(t *testing.T) {
	got := ReverseMap(map[int]string{1: "x"})
	if got["x"] != 1 {
		t.Fatal()
	}
}

func TestSelect(t *testing.T) {
	got := Select([]map[string]any{{"a": 1, "b": "x", "c": 2.0}, {"a": 2, "b": "y", "c": 1.5}}, "a", "b")
	if len(got) != 2 || got[0]["a"] != 1 || got[0]["b"] != "x" || got[1]["a"] != 2 || got[1]["b"] != "y" {
		t.Fatalf("got %v", got)
	}
}

func TestShuffle(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	out := Shuffle(slices.Clone(in))
	if len(out) != len(in) {
		t.Fatal()
	}
	sort.Ints(in)
	sort.Ints(out)
	if !slices.Equal(in, out) {
		t.Fatalf("multiset mismatch")
	}
}

func TestSkip(t *testing.T) {
	if !slices.Equal(Skip([]int{1, 1, 1, 2, 3}, 2), []int{1, 2, 3}) {
		t.Fatal()
	}
}

func TestSkipUntil(t *testing.T) {
	if !slices.Equal(SkipUntil([]int{1, 2, 3}, func(n int) bool { return n == 2 }), []int{2, 3}) {
		t.Fatal()
	}
	if !slices.Equal(SkipUntil([]int{1, 2}, func(int) bool { return false }), []int{1, 2}) {
		t.Fatal()
	}
}

func TestSkipWhile(t *testing.T) {
	if !slices.Equal(SkipWhile([]int{2, 2, 3}, func(n int) bool { return n == 2 }), []int{3}) {
		t.Fatal()
	}
	if !slices.Equal(SkipWhile([]int{1, 2}, func(int) bool { return false }), []int{1, 2}) {
		t.Fatal()
	}
	if !slices.Equal(SkipWhile([]int{1, 2}, func(int) bool { return true }), []int{1, 2}) {
		t.Fatal("expected full slice when predicate never false")
	}
}

func TestSlice(t *testing.T) {
	if !slices.Equal(Slice([]int{0, 1, 2}, 0, 2), []int{0, 1}) {
		t.Fatal()
	}
}

func TestSort(t *testing.T) {
	a := []int{3, 1, 2}
	b := Sort(slices.Clone(a), SortDirectionAsc)
	sort.Ints(a)
	if !slices.Equal(b, a) {
		t.Fatalf("got %v", b)
	}
	c := Sort([]int{1, 3, 2}, SortDirectionDesc)
	if !slices.Equal(c, []int{3, 2, 1}) {
		t.Fatalf("got %v", c)
	}
}

func TestSortAsc(t *testing.T) {
	x := []int{2, 1}
	SortAsc(x)
	if !slices.Equal(x, []int{1, 2}) {
		t.Fatal()
	}
}

func TestSortDesc(t *testing.T) {
	x := []int{1, 3, 2}
	SortDesc(x)
	if !slices.Equal(x, []int{3, 2, 1}) {
		t.Fatal()
	}
}

func TestSortBy(t *testing.T) {
	rows := []map[string]int{{"v": 2}, {"v": 1}}
	SortBy(rows, "v", SortDirectionAsc)
	if rows[0]["v"] != 1 {
		t.Fatalf("got %v", rows)
	}
	desc := []map[string]int{{"v": 1}, {"v": 3}}
	SortBy(desc, "v", SortDirectionDesc)
	if desc[0]["v"] != 3 {
		t.Fatalf("got %v", desc)
	}
}

func TestSortByAsc(t *testing.T) {
	rows := []map[string]int{{"v": 3}, {"v": 1}}
	SortByAsc(rows, "v")
	if rows[0]["v"] != 1 {
		t.Fatal()
	}
}

func TestSortByDesc(t *testing.T) {
	rows := []map[string]int{{"v": 1}, {"v": 3}}
	SortByDesc(rows, "v")
	if rows[0]["v"] != 3 {
		t.Fatal()
	}
}

func TestSplice(t *testing.T) {
	got := Splice([]int{1, 2, 3, 4}, 1, 2, 9, 9)
	if !slices.Equal(got, []int{1, 9, 9, 4}) {
		t.Fatalf("got %v", got)
	}
}

func TestSplit(t *testing.T) {
	got := Split([]int{1, 2, 3, 4, 5}, 2)
	if len(got) != 3 || !slices.Equal(got[2], []int{5}) {
		t.Fatalf("got %v", got)
	}
	if Split([]int{1}, 0) != nil {
		t.Fatal()
	}
}

func TestSplitInto(t *testing.T) {
	// SplitInto(c, n) calls Split(c, len(c)/n): 6 elements, n=2 → chunk size 3
	got := SplitInto([]int{1, 2, 3, 4, 5, 6}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestTake(t *testing.T) {
	if !slices.Equal(Take([]int{1, 2, 3}, 2), []int{1, 2}) {
		t.Fatal()
	}
}

func TestTakeUntil(t *testing.T) {
	if !slices.Equal(TakeUntil([]int{1, 2, 3}, func(n int) bool { return n == 2 }), []int{1}) {
		t.Fatal()
	}
	if !slices.Equal(TakeUntil([]int{1, 2}, func(int) bool { return false }), []int{1, 2}) {
		t.Fatal()
	}
}

func TestTakeWhile(t *testing.T) {
	if !slices.Equal(TakeWhile([]int{1, 2, 3}, func(n int) bool { return n < 3 }), []int{1, 2}) {
		t.Fatal()
	}
	if !slices.Equal(TakeWhile([]int{1, 2}, func(int) bool { return true }), []int{1, 2}) {
		t.Fatal()
	}
}

func TestUnion(t *testing.T) {
	got := Union([]int{1, 2}, []int{2, 3})
	if len(got) != 3 {
		t.Fatalf("got %v", got)
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 1, 2, 3, 3})
	sort.Ints(got)
	if !slices.Equal(got, []int{2}) {
		t.Fatalf("got %v", got)
	}
}

func TestWhere(t *testing.T) {
	got := Where([]map[string]int{{"k": 1}, {"k": 2}}, "k", 2)
	if len(got) != 1 {
		t.Fatal()
	}
}

func TestWhereBetween(t *testing.T) {
	got := WhereBetween([]map[string]int{{"x": 2}, {"x": 5}, {"x": 10}}, "x", 2, 5)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestWhereIn(t *testing.T) {
	got := WhereIn([]map[string]int{{"t": 1}, {"t": 9}}, "t", []int{9})
	if len(got) != 1 {
		t.Fatal()
	}
}

func TestWhereNot(t *testing.T) {
	got := WhereNot([]map[string]int{{"t": 1}, {"t": 2}}, "t", 1)
	if len(got) != 1 || got[0]["t"] != 2 {
		t.Fatal()
	}
}

func TestWhereNotBetween(t *testing.T) {
	got := WhereNotBetween([]map[string]int{{"x": 1}, {"x": 4}, {"x": 7}}, "x", 3, 6)
	// outside [3,6] are 1 and 7
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestWhereNotIn(t *testing.T) {
	got := WhereNotIn([]map[string]int{{"t": 1}, {"t": 2}}, "t", []int{1})
	if len(got) != 1 || got[0]["t"] != 2 {
		t.Fatal()
	}
}

func TestZip(t *testing.T) {
	got := Zip([]int{1, 2}, []string{"a", "b"})
	if len(got) != 2 || !reflect.DeepEqual(got[0], []any{1, "a"}) {
		t.Fatalf("got %v", got)
	}
}

func TestChunk_error(t *testing.T) {
	_, err := Chunk([]int{1}, -1)
	if err == nil || err.Error() != "size must be greater than 0" {
		t.Fatalf("got %v", err)
	}
}
