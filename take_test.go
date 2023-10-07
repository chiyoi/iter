package iter

import "testing"

func TestSkipTake(t *testing.T) {
	AssertEqual(t, []int{3, 4, 5}, Iter([]int{1, 2, 3, 4, 5}).Skip(2).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))

	AssertEqual(t, []int{3, 4}, Iter([]int{1, 2, 3, 4, 5}).Skip(2).Take(2).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}

func TestFilter(t *testing.T) {
	AssertEqual(t, []int{3, 4}, Iter([]int{1, 2, 3, 4, 5}).Filter(func(v any) bool {
		return v.(int) > 2 && v.(int) <= 4
	}).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}
