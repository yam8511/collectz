package collectz

import (
	"constraints"
)

type FilterFn[T any] func(int, T) bool
type EachFn[T, T2 any] func(int, T) T2

func IndexOf[T comparable](datas []T, target T) int {
	var index = -1
	for i := 0; i < len(datas); i++ {
		if datas[i] == target {
			index = i
			break
		}
	}

	return index
}

func LastIndexOf[T comparable](datas []T, target T) int {
	var index = -1
	for i := 0; i < len(datas); i++ {
		if datas[i] == target {
			index = i
		}
	}

	return index
}

func IndexOfAny[T2, T comparable](datas []T, target T2, fn func(int, T) T2) int {
	var index = -1
	if fn == nil {
		return index
	}
	for i := 0; i < len(datas); i++ {
		if fn(i, datas[i]) == target {
			index = i
			break
		}
	}

	return index
}

func Unique[T constraints.Ordered](datas []T) []T {
	var m = map[any]struct{}{}

	return Filter(datas, func(i int, t T) bool {
		_, ok := m[t]
		m[t] = struct{}{}
		return ok
	})
}

func UniqueAny[T any](datas []T, fn func(int, T) any) []T {
	var m = map[any]struct{}{}
	if fn == nil {
		fn = func(i int, t T) any {
			return t
		}
	}

	return Filter(datas, func(i int, t T) bool {
		key := fn(i, t)
		_, ok := m[key]
		m[key] = struct{}{}
		return ok
	})
}

func Map[T2, T any](datas []T, fn EachFn[T, T2]) []T2 {
	if len(datas) == 0 {
		return []T2{}
	}

	var datas2 = make([]T2, len(datas))
	if fn == nil {
		fn = func(int, T) T2 {
			var t2 T2
			return t2
		}
	}

	for i := 0; i < len(datas); i++ {
		datas2[i] = fn(i, datas[i])
	}

	return datas2
}

func Filter[T any](datas []T, filters ...FilterFn[T]) []T {
	if len(datas) == 0 {
		return []T{}
	}

	var datas2 = make([]T, 0)
	for i := 0; i < len(datas); i++ {
		var pass bool
		for _, filter := range filters {
			if filter != nil && filter(i, datas[i]) {
				pass = true
				break
			}
		}

		if !pass {
			datas2 = append(datas2, datas[i])
		}
	}

	return datas2
}

// The chunk method breaks the collection into multiple, smaller collections of a given size
func Chunk[T any](datas []T, size int) [][]T {
	var chunks = make([][]T, 0)
	var chunk = make([]T, 0, size)
	for i := 0; i < len(datas); i++ {
		chunk = append(chunk, datas[i])
		if i%size == size-1 { // 1, 2, 3, n
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, size)
		}
	}

	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks
}

// The last method returns the last element in the collection that passes a given truth test
func Last[T any](datas []T, filters ...FilterFn[T]) (T, bool) {
	var v T
	if len(datas) == 0 {
		return v, false
	}

	if len(filters) == 0 {
		return datas[len(datas)-1], true
	}

	var found bool
	for i, data := range datas {
		for _, filter := range filters {
			if filter != nil && filter(i, data) {
				v = data
				found = true
			}
		}
	}

	// TODO: filters given and not found last, then still return last ?
	// if !found {
	// 	return datas[len(datas)-1], true
	// }

	return v, found
}

// The first method returns the first element in the collection that passes a given truth test
func First[T any](datas []T, filters ...FilterFn[T]) (T, bool) {
	var v T
	if len(datas) == 0 {
		return v, false
	}

	if len(filters) == 0 {
		return datas[0], true
	}

	var found bool
	for i, data := range datas {
		for _, filter := range filters {
			if filter != nil && filter(i, data) {
				found = true
				v = data
				break
			}
		}

		if found {
			break
		}
	}

	// TODO: filters given and not found first, then still return first ?
	// if !found {
	// 	return datas[0], true
	// }

	return v, found
}
