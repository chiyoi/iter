package iter

import (
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	AssertEqual(t, []int{1, 1, 1}, From(func() (v any) { return 1 }).Take(3).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}

func TestForEach(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5}
	Iter(sli).ForEach(func(v any, i int) bool {
		AssertEqual(t, sli[i], v)
		return Continue
	})

	Iter(sli).ForEach(func(v any, i int) bool {
		if i > 3 {
			t.Error("out of bound")
		}

		if i == 3 {
			return Break
		}
		return Continue
	})
}

func TestRange(t *testing.T) {
	AssertEqual(t, []int{0, 1, 2, 3, 4}, Range(0, 5).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}

func TestEmptyRepeat(t *testing.T) {
	AssertEqual(t, []int(nil), Empty().Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))

	AssertEqual(t, []int{2, 2, 2, 2, 2}, Repeat(2).Take(5).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}

func TestChain(t *testing.T) {
	AssertEqual(t, []int{3, 3, 3, 2, 2}, Repeat(3).Take(3).Chain(Repeat(2).Take(2)).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}

func TestZip(t *testing.T) {
	AssertEqual(t, [][2]any{{3, 2}, {3, 2}}, Zip(Repeat(3).Take(3), Repeat(2).Take(2)).Fold([][2]any(nil), func(st, v any) any {
		return append(st.([][2]any), v.([2]any))
	}))
}

func AssertEqual(t *testing.T, exp, rev any) {
	if !reflect.DeepEqual(exp, rev) {
		t.Errorf("expect: %v, got: %v", exp, rev)
	}
}
