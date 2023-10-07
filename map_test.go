package iter

import (
	"testing"
)

func TestMap(t *testing.T) {
	AssertEqual(t, []complex128{complex(1, 0), complex(2, 0), complex(3, 0)}, Iter([]int{1, 2, 3}).Map(func(a any) any {
		return complex(float64(a.(int)), 0)
	}).Fold([]complex128(nil), func(st, v any) any {
		return append(st.([]complex128), v.(complex128))
	}))
}

func TestScan(t *testing.T) {
	AssertEqual(t, []float64{2, 4, 7, 11, 16}, Iter([]int{1, 2, 3, 4, 5}).Scan(float64(1), func(st, v any) any {
		return float64(st.(float64) + float64(v.(int)))
	}).Fold([]float64(nil), func(st, v any) any {
		return append(st.([]float64), v.(float64))
	}))

	AssertEqual(t, []int{1, 3, 6}, Iter([]int{1, 2, 3, 4, 5}).Scan(0, func(st, v any) any {
		if v.(int) > 3 {
			return nil
		}
		return st.(int) + v.(int)
	}).Fold([]int(nil), func(st, v any) any {
		return append(st.([]int), v.(int))
	}))
}
