package iter

import (
	"reflect"
	"testing"
)

func TestFold(t *testing.T) {
	for i, tc := range []struct {
		sli []int
		st  int
		f   func(int, int) int
		out int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			0,
			func(st int, e int) int { return st + e },
			15,
		},
	} {
		out := Fold(tc.sli, tc.st, tc.f)
		if out != tc.out {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}

func TestScan(t *testing.T) {
	for i, tc := range []struct {
		sli []int
		st  int
		f   func(int, int) int
		out []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			0,
			func(st int, e int) int { return st + e },
			[]int{1, 3, 6, 10, 15},
		},
	} {
		out := Scan(tc.sli, tc.st, tc.f)
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}
