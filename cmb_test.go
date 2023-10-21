package iter

import (
	"testing"
)

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
