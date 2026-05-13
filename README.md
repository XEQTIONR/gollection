<img width="2000" height="800" alt="GollectionBanner" src="https://github.com/user-attachments/assets/8bf18ccf-11ec-420f-81ca-de5bb96aad07" />


# Gollection
<a href="https://github.com/XEQTIONR/gollection/actions/workflows/test.yaml"><img src="https://github.com/XEQTIONR/gollection/actions/workflows/test.yaml/badge.svg" alt="Go Tests" style="max-width: 100%;"></a>
<a href="https://github.com/XEQTIONR/gollection/?tab=MIT-1-ov-file#readme"><img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License" style="max-width: 100%;"></a>

### About Gollection

Gollection is a simple library that provides a convenient helpers for working with collections of data. The word "Gollection" is a [portmanteau](https://en.wikipedia.org/wiki/Portmanteau) of the words "go" and "collection".

#### What is a collection?
A collection is a container of related data such as a slice or a map.

#### What can gollection do?
It can manipulate, partition, transform, and get stats about collections.

#### But why?
In an attempt to keep our go files small and clean.



### Available Functions

* [Append](#append)
* [At](#at)
* [Average](#average)
* [Avg](#avg)
* [Chunk](#chunk)
* [ChunkBy](#chunkby)
* [Collapse](#collapse)
* [CollapseMap](#collapsemap)
* [Combine](#combine)
* [CombineSlice](#combineslice)
* [CombineMap](#combinemap)
* [Contains](#contains)
* [ContainsKey](#containskey)
* [Count](#count)
* [CountBy](#countby)
* [Counts](#counts)
* [CrossJoin](#crossjoin)
* [Diff](#diff)
* [DiffAssoc](#diffassoc)
* [DiffKeys](#diffkeys)
* [DoesntContain](#doesntcontain)
* [Duplicates](#duplicates)
* [Each](#each)
* [Every](#every)
* [Except](#except)
* [Filter](#filter)
* [First](#first)
* [FirstOr](#firstor)
* [FirstOrFail](#firstorfail)
* [FirstWhere](#firstwhere)
* [Flatten](#flatten)
* [FlattenMap](#flattenmap)
* [Flip](#flip)
* [Forget](#forget)
* [GroupBy](#groupby)
* [Has](#has)
* [HasAny](#hasany)
* [HasAnyKeys](#hasanykeys)
* [HasEvery](#hasevery)
* [HasOne](#hasone)
* [IndexOf](#indexof)
* [IsEmpty](#isempty)
* [Intersect](#intersect)
* [IntersectAssoc](#intersectassoc)
* [IntersectByKeys](#intersectbykeys)
* [KeyBy](#keyby)
* [Last](#last)
* [LastOr](#lastor)
* [LastOrFail](#lastorfail)
* [LastWhere](#lastwhere)
* [Length](#length)
* [Map](#map)
* [Max](#max)
* [Mode](#mode)
* [Min](#min)
* [Multiply](#multiply)
* [Nth](#nth)
* [NthFromLast](#nthfromlast)
* [Only](#only)
* [PadLeft](#padleft)
* [PadRight](#padright)
* [Partition](#partition)
* [Percentage](#percentage)
* [Prepend](#prepend)
* [Push](#push)
* [Pop](#pop)
* [Random](#random)
* [Range](#range)
* [Reduce](#reduce)
* [ReduceWithIndex](#reducewithindex)
* [Reject](#reject)
* [RejectWithIndex](#rejectwithindex)
* [Replace](#replace)
* [Reverse](#reverse)
* [ReverseMap](#reversemap)
* [Select](#select)
* [Shuffle](#shuffle)
* [Skip](#skip)
* [SkipUntil](#skipuntil)
* [SkipWhile](#skipwhile)
* [Slice](#slice)
* [Sort](#sort)
* [SortAsc](#sortasc)
* [SortBy](#sortby)
* [SortByAsc](#sortbyasc)
* [SortByDesc](#sortbydesc)
* [SortDesc](#sortdesc)
* [Splice](#splice)
* [Split](#split)
* [SplitInto](#splitinto)
* [Take](#take)
* [TakeUntil](#takeuntil)
* [TakeWhile](#takewhile)
* [Union](#union)
* [Unique](#unique)
* [Where](#where)
* [WhereBetween](#wherebetween)
* [WhereIn](#wherein)
* [WhereNot](#wherenot)
* [WhereNotBetween](#wherenotbetween)
* [WhereNotIn](#wherenotin)
* [Zip](#zip)

#### Append
```go
gollection.Append([]int{1, 2, 3}, 4, 5, 6)
// [1, 2, 3, 4, 5, 6]
```

#### At
```go
gollection.At([]int{1, 2, 3}, 1)
// 2, nil
```

#### Average
```go
gollection.Average([]int{1, 2, 3})
// 2
```

#### Avg
alias for `Average`

#### Chunk
```go
gollection.Chunk([]int{1, 2, 3, 4, 5, 6}, 3)
// [[1, 2, 3], [4, 5, 6]]
```

#### ChunkBy
```go
gollection.ChunkBy([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 1 })
// [[1, 2], [3, 4], [5, 6]]
```

#### Collapse
```go
gollection.Collapse([][]int{{1, 2}, {3, 4}, {5, 6}})
// [1, 2, 3, 4, 5, 6]
```

#### CollapseMap
```go
gollection.CollapseMap([]map[string]int{{"a": 1, "c": 4}, {"b": 2, "a": 3}})
// {"a": 3, "b": 2, "c": 4}
```

#### Combine
```go
gollection.Combine([]string{"a", "b"}, []int{1, 2})
// {"a": 1, "b": 2}
```

#### CombineSlice
alias for `Combine`

#### CombineMap
```go
gollection.CombineMap(map[string]int{"a": 1}, map[string]int{"b": 2, "a": 3})
// {"a": 3, "b": 2}
```

#### Contains
```go
gollection.Contains([]int{1, 2, 3, 4}, 4)
// true
```

#### ContainsKey
```go
gollection.ContainsKey(map[string]int{"k": 1}, "k")
// true
```

#### Count
alias for `len`

#### CountBy
```go
gollection.CountBy([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
// 2
```

#### Counts
```go
gollection.Counts([]string{"a", "a", "b"})
// {"a": 2, "b": 1}
```

#### CrossJoin
```go
gollection.CrossJoin([]int{1, 2, 3}, []string{"x", "y"})
// [[1, "x"], [1, "y"], [2, "x"], [2, "y"], [3, "x"], [3, "y"]]
```

#### Diff
```go
gollection.Diff([]int{1, 2, 3}, []int{2})
// [1,3]
```

#### DiffAssoc
```go
gollection.DiffAssoc(map[string]int{"a": 1, "b": 3}, map[string]int{"a": 1, "b": 2})
// {"b": 3}
```

#### DiffKeys
```go
gollection.DiffKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3})
// ["a"]
```

#### DoesntContain
```go
gollection.DoesntContain([]int{1, 3, 4, 5}, 2)
// true
```

#### Duplicates
```go
gollection.Duplicates([]int{1, 1, 2, 3, 3})
// [1,3]
```

#### Each
```go
sum:= 0
gollection.Each([]int{1, 2, 3}, func(n int) { sum += n })
// sum = 6
```

#### Every
```go
gollection.Every([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 })
// true
```

#### Except
```go
gollection.Except([]int{1, 2, 3}, func(n int) bool { return n == 2 })
// [1,3]
```

#### Filter
```go
gollection.Filter([]int{1, 2, 3}, func(n int) bool { return n == 2 })
// [2]
```

#### First
```go
gollection.First([]int{1, 2, 3}, func(n, i int) bool { return n > 2 })
// 3
```

#### FirstOr
```go
gollection.FirstOr([]int{1, 2, 3}, func(n, i int) bool { return n > 3 }, 4)
// 4
```

#### FirstOrFail
```go
gollection.FirstOrFail([]int{1, 2, 3}, func(n, i int) bool { return n > 3 })
// panic
```

#### FirstWhere
```go
gollection.FirstWhere([]map[string]any{{"id": 1, "size": "S"}, {"id": 2, "size": "L"}}, "size", "L")
// {"id": 2, "size": "L"}
```

#### Flatten
```go
gollection.Flatten([][]int{{1, 2, 3}, {4, 5}, {6}})
// [1,2,3,4,5,6]
```

#### FlattenMap
```go
gollection.FlattenMap([]map[string]int{{"a": 1}, {"b": 2, "a": 3}, {"a": 5, "c": 4}})
// {"a": 5, "b": 2, "c": 4}
```

#### Flip
```go
gollection.Flip(map[string]int{"b": 2, "a": 3})
// {2: "b", 3: "c"}
```

#### Forget
```go
gollection.Forget(map[string]int{"a": 2, "b": 3, "c": 4}, "a", "c")
// {"b": 3}
```

#### GroupBy
```go
gollection.GroupBy([]map[string]string{{"t": "a", "u": "1"}, {"t": "b", "u": "2"}, {"t": "a", "u": "3"}}, "t")
// {"a": [{"t": "a", "u": "1"}, {"t": "a", "u": "3"}], "b": {"t": "b", "u": "2"}}
```

#### Has
alias for `Contains`

### HasAny
```go
gollection.HasAny([]int{1, 2, 3}, []int{0, 3})
// true
```

#### HasAnyKeys
```go
gollection.HasAnyKeys(map[string]int{"a": 1}, "a", "b")
// true
```

#### HasEvery
```go
gollection.HasEvery([]int{1, 2, 3}, []int{1, 2})
// true
```

#### HasOne
```go
gollection.HasOne([]int{1, 4, 5}, []int{1, 2, 3})
// true
gollection.HasOne([]int{1, 2, 5}, []int{1, 2, 3})
// false
```

#### IndexOf
```go
gollection.IndexOf([]string{"a", "b", "c"}, "b")
// 1
gollection.IndexOf([]string{"a", "b", "c"}, "d")
// -1
```

#### IsEmpty
```go
gollection.IsEmpty([]string{"a", "b", "c"})
// false
```

#### Intersect
```go
gollection.Intersect([]int{1, 1, 2}, []int{1, 2, 2})
// [1, 2]
```

#### IntersectAssoc
```go
gollection.IntersectAssoc(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 9})
// {"a": 1}
```

#### IntersectByKeys
```go
gollection.IntersectByKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 9})
// {"b": 2}
```

#### KeyBy
```go
gollection.KeyBy([]map[string]int{{"k": 1}, {"k": 2, "a": 3}, {"k": 2, "b": 4}}, "k")
// {1: [{"k": 1}], 2: [{"k": 2, "a": 3}, {"k": 2, "b": 4}]}
```

#### Last
```go
gollection.Last([]int{1, 2, 3, 4}, func(n, i int) bool { return n >= 2 })
// 4
gollection.Last([]int{1, 2, 3, 4}, func(n, i int) bool { return n > 4 })
// nil
```

#### LastOr
```go
gollection.LastOr([]int{1, 2, 3, 4}, func(n, i int) bool { return n >= 4 }, 99)
// 4
gollection.LastOr([]int{1, 2, 3, 4}, func(n, i int) bool { return n > 4 }, 99)
// 99
```

#### LastOrFail
```go
gollection.LastOrFail([]int{1, 2, 3, 4}, func(n, i int) bool { return n >= 4 })
// 4
gollection.LastOrFail([]int{1, 2, 3, 4}, func(n, i int) bool { return n > 4 })
// panic
```

#### LastWhere
```go
gollection.LastWhere([]map[string]any{{"n": 1}, {"n": 2}, {"n": 2}}, "n", 2)
// {"n": 2}
gollection.LastWhere([]map[string]any{{"n": 1}, {"n": 2}, {"n": 2}}, "n", 3)
// nil
```

#### Length
alias for `Count`

### Map
```go
gollection.Map([]int{1, 2}, func(n int) string { return string(rune('0' + n)) })
// ["1", "2"]
```


#### Max
```go
gollection.Max([]int{1, 3, 2, 5, 4})
// 5
```

#### Mode
```go
gollection.Mode([]int{1, 2, 1, 2, 3, 2})
// ["1", "2"]
```

#### Min
```go
gollection.Min([]int{3, 2, 5, 4})
// 2
```

#### Multiply
```go
gollection.Multiply([]int{1, 2}, 3)
// [1, 1, 1, 2, 2, 2]
```

#### Nth
```go
gollection.Nth([]string{"a", "b"}, 1)
// "b"
```

#### NthFromLast
```go
gollection.NthFromLast([]string{"a", "b"}, 1)
// "a"
```

#### Only
```go
gollection.Only([]map[string]int{{"a": 1, "b": 2}, {"a": 3, "b": 4, "c": 5}}, "a", "b")
// [{"a": 1, "b": 2}, {"a": 3, "b": 4}]
```

#### PadLeft
```go
gollection.PadLeft([]int{1}, 2, 0)
// [0, 0, 1]
```

#### PadRight
```go
gollection.PadRight([]int{1}, 4, 3)
// [1, 3, 3, 3, 3]
```

#### Partition
```go
gollection.Partition([]int{1, 2, 3, 4}, func(i, n int) bool { return n%2 == 0 })
// [[2,4], [1,3]]
gollection.Partition([]int{1, 2, 3, 4}, func(i, n int) bool { return i < 2 })
// [[1,2], [3,4]]
```


#### Percentage
```go
gollection.Percentage([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
// 50
```

#### Prepend
```go
gollection.Prepend([]int{3, 4}, 1, 2)
// [1, 2, 3, 4]
```

#### Push
```go
gollection.Push([]int{3, 4}, 1)
// [3, 4, 1]
```

#### Pop
```go
gollection.Pop([]int{3, 4, 1})
// 1, [3, 4]
```

#### Random
```go
gollection.Random([]int{1, 2, 3, 4, 5})
// 4, nil
gollection.Random([]int{})
// nil, Error
```

#### Range
```go
gollection.Range([]int{0, 1, 2, 3, 4, 5}, 2, 4)
// [2, 3], nil
gollection.Range([]int{0, 1, 2, 3, 4, 5}, 4, 2)
// nil, Error
```

#### Reduce
```go
gollection.Reduce([]int{1, 2, 3}, func(acc, n int) int { return acc + n }, 0)
// 6
```

#### ReduceWithIndex
```go
gollection.ReduceWithIndex([]int{1, 2, 3}, func(acc, n, i int) int { return acc + i }, 0)
// 3
```

#### Reject
```go
gollection.Reject([]int{1, 2, 3}, func(n int) bool { return n == 2 })
// [1, 3]
```

#### RejectWithIndex
```go
gollection.RejectWithIndex([]int{1, 2, 3}, func(n, i int) bool { return i == 2 })
// [1, 2]
```

#### Replace
```go
gollection.Replace([]int{1, 2, 3, 4, 5}, map[int]int{1: 9, 3: 10})
// [1, 9, 3, 10, 5]
```

#### Reverse
```go
gollection.Reverse([]int{1, 2, 3, 4, 5})
// [5, 4, 3, 2, 1]
```

#### ReverseMap
```go
gollection.ReverseMap(map[int]string{1: "x", 3: "y"})
// {"x": 1, "y": 3}
```

#### Select
```go
gollection.Select([]map[string]any{{"a": 1, "b": "x", "c": 2.0}, {"a": 2, "b": "y", "c": 1.5}}, "a", "b")
// [{"a": 1, "b": "x"}, {"a": 2, "b": "y"}]
```


#### Shuffle
```go
gollection.Shuffle([]int{1, 2, 3, 4, 5})
// [4, 1, 3, 5, 2]
```

#### Skip
```go
gollection.Skip([]int{1, 1, 1, 2, 3}, 2)
// [1, 2, 3]
```

#### SkipUntil
```go
gollection.SkipUntil([]int{1, 2, 3}, func(n int) bool { return n == 2 })
// [2, 3]
```

#### SkipWhile
```go
gollection.SkipWhile([]int{1, 2, 1}, func(n int) bool { return n < 2 })
// [2, 1]
```

#### Slice
```go
gollection.Slice([]int{1, 2, 3, 4, 5}, 1, 3)
// [2, 3]
```

#### Sort
```go
gollection.Sort([]int{3, 1, 2}, gollection.SortDirectionAsc)
// [1, 2, 3]
gollection.Sort([]int{1, 3, 2}, gollection.SortDirectionDesc)
// [3, 2, 1]
```

#### SortAsc
```go
gollection.SortAsc([]int{2, 1})
// [1, 2]
```

#### SortBy
```go
gollection.SortBy([]map[string]int{{"v": 2}, {"v": 1}}, "v", gollection.SortDirectionAsc)
// [{"v": 1}, {"v": 2}]
```

#### SortByAsc
```go
gollection.SortByAsc([]map[string]int{{"v": 3}, {"v": 1}}, "v")
// [{"v": 1}, {"v": 3}]
```

#### SortByDesc
```go
gollection.SortByDesc([]map[string]int{{"v": 1}, {"v": 3}}, "v")
// [{"v": 3}, {"v": 1}]
```

#### SortDesc
```go
gollection.SortDesc([]int{1, 3, 2})
// [3, 2, 1]
```

#### Splice
```go
gollection.Splice([]int{1, 2, 3, 4}, 1, 2, 9, 9)
// [1, 9, 9, 4]
```

#### Split
```go
gollection.Split([]int{1, 2, 3, 4, 5}, 2)
// [[1, 2], [3, 4], [5]]
```

#### SplitInto
```go
gollection.SplitInto([]int{1, 2, 3, 4, 5, 6}, 2)
// [[1, 2, 3], [4, 5, 6]]
```

#### Take
```go
gollection.Take([]int{1, 2, 3}, 2)
// [1, 2]
```

#### TakeUntil
```go
gollection.TakeUntil([]int{1, 2, 3}, func(n int) bool { return n == 2 })
// [1]
```

#### TakeWhile
```go
gollection.TakeWhile([]int{1, 2, 3}, func(n int) bool { return n < 3 })
// [1, 2]
```

#### Union
```go
gollection.Union([]int{1, 2}, []int{2, 3})
// [1, 2, 3]
```

#### Unique
```go
gollection.Unique([]int{1, 1, 2, 3, 3})
// [2]
```

#### Where
```go
gollection.Where([]map[string]int{{"k": 1}, {"k": 2}}, "k", 2)
// [{"k": 2}]
```

#### WhereBetween
```go
gollection.WhereBetween([]map[string]int{{"x": 2}, {"x": 5}, {"x": 10}}, "x", 2, 5)
// [{"x": 2}, {"x": 5}]
```

#### WhereIn
```go
gollection.WhereIn([]map[string]int{{"t": 1}, {"t": 9}}, "t", []int{9})
// [{"t": 9}]
```

#### WhereNot
```go
gollection.WhereNot([]map[string]int{{"t": 1}, {"t": 2}}, "t", 1)
// [{"t": 2}]
```

#### WhereNotBetween
```go
gollection.WhereNotBetween([]map[string]int{{"x": 1}, {"x": 4}, {"x": 7}}, "x", 3, 6)
// [{"x": 1}, {"x": 7}]
```

#### WhereNotIn
```go
gollection.WhereNotIn([]map[string]int{{"t": 1}, {"t": 2}}, "t", []int{1})
// [{"t": 2}]
```

#### Zip
```go
gollection.Zip([]int{1, 2}, []string{"a", "b"})
// [[1, "a"], [2, "b"]]
```
