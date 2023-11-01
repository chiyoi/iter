package iter

import (
	"testing"
)

func TestCall(t *testing.T) {
	for i, tc := range []struct {
		f   func(int) int
		x   int
		out int
	}{
		{func(x int) int { return x + 1 }, 1, 2},
		{func(x int) int { return x + 1 }, 2, 3},
		{func(x int) int { return 0 }, 1, 0},
	} {
		c := Call(tc.f, tc.x)
		if out := c(); out != tc.out {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}

func TestY(t *testing.T) {
	fact := Y(func(fact func(int) int) func(int) int {
		return func(x int) int {
			if x <= 1 {
				return 1
			}
			return x * fact(x-1)
		}
	})
	for i, tc := range []struct {
		in, out int
	}{
		{5, 120},
	} {
		if out := fact(tc.in); out != tc.out {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}
