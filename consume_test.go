package iter

import (
	"testing"
)

func TestLastAt(t *testing.T) {
	AssertEqual(t, 5, Iter([]int{1, 2, 3, 4, 5}).Last())
	AssertEqual(t, 3, Iter([]int{1, 2, 3, 4, 5}).At(2))
}

func TestReduce(t *testing.T) {
	AssertEqual(t, 15, Iter([]int{1, 2, 3, 4, 5}).Reduce(func(a any, b any) any {
		return a.(int) + b.(int)
	}))
}

func TestFold(t *testing.T) {
	AssertEqual(t, 16, Iter([]int{1, 2, 3, 4, 5}).Fold(1, func(st, v any) any {
		return st.(int) + v.(int)
	}))

	AssertEqual(t, 3, Iter([]int{1, 2, 3, 4, 5}).Fold(0, func(st, v any) any {
		if v.(int) > 3 {
			return nil
		}
		return v
	}))
}
