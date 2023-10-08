package iter

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIterCollect(t *testing.T) {
	for i, tc := range [][2][]int{
		{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}},
		{{0}, {0}},
		{{}, nil},
		{nil, nil},
	} {
		got := Collect(Iter(tc[0]))
		if !reflect.DeepEqual(got, tc[1]) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, got, tc[1])
		}
	}
}

func TestRange(t *testing.T) {
	for i, tc := range []struct {
		a, b int
		step int
		out  []int
	}{
		{0, 10, 1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{0, 10, 3, []int{0, 3, 6, 9}},
		{10, 0, -3, []int{10, 7, 4, 1}},
		{-3, 3, 1, []int{-3, -2, -1, 0, 1, 2}},
		{0, 0, 1, nil},
		{10, 0, 1, nil},
	} {
		out := Collect(Range(tc.a, tc.b, tc.step))
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}

func TestEmptyRepeatChainSkipTake(t *testing.T) {
	for i, tc := range []struct {
		n                    int
		countSkip, countTake int
		out                  []int
	}{
		{1, 3, 3, []int{1, 1, 1}},
		{0, 0, 3, []int{0, 0, 0}},
		{-1, -1, 3, []int{-1, -1, -1}},
		{1, 3, 0, nil},
		{1, 3, -1, nil},
	} {
		out := Collect(Take(Skip(Chain(Empty[int](), Repeat(tc.n)), tc.countSkip), tc.countTake))
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}

func TestZip(t *testing.T) {
	for i, tc := range []struct {
		a   []int
		b   []string
		out []ZipItem[int, string]
	}{
		{[]int{1, 2, 3}, []string{"a", "b", "c"}, []ZipItem[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}},
		{[]int{1, 2}, []string{"a", "b", "c"}, []ZipItem[int, string]{{1, "a"}, {2, "b"}}},
		{[]int{1, 2, 3}, []string{"a", "b"}, []ZipItem[int, string]{{1, "a"}, {2, "b"}}},
		{nil, []string{"a", "b", "c"}, nil},
		{[]int{1, 2, 3}, nil, nil},
	} {
		out := Collect(Zip(Iter(tc.a), Iter(tc.b)))
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}

func TestLast(t *testing.T) {
	for i, tc := range []struct {
		sli []int
		out int
		ok  bool
	}{
		{[]int{1, 2, 3}, 3, true},
		{[]int{0}, 0, true},
		{[]int{}, 0, false},
		{nil, 0, false},
	} {
		out, ok := Last(Iter(tc.sli))
		if out != tc.out || ok != tc.ok {
			t.Errorf("Testcase %d: got (%v, %v), expect (%v, %v).", i, out, ok, tc.out, tc.ok)
		}
	}
}

func TestAt(t *testing.T) {
	for i, tc := range []struct {
		sli []int
		i   int
		out int
		ok  bool
	}{
		{[]int{1, 2, 3}, 2, 3, true},
		{[]int{1, 2, 3}, 3, 0, false},
		{[]int{1, 2, 3}, 1, 2, true},
		{[]int{1, 2, 3}, -1, 0, false},
		{[]int{}, 0, 0, false},
		{nil, 0, 0, false},
	} {
		out, ok := At(Iter(tc.sli), tc.i)
		if out != tc.out || ok != tc.ok {
			t.Errorf("Testcase %d: got (%v, %v), expect (%v, %v).", i, out, ok, tc.out, tc.ok)
		}
	}
}

func TestMap(t *testing.T) {
	for i, tc := range []struct {
		a   []int
		f   func(int) string
		out []string
	}{
		{[]int{1, 2, 3, 4, 5}, func(a int) string { return strconv.Itoa(a + 1) }, []string{"2", "3", "4", "5", "6"}},
		{nil, nil, nil},
		{nil, func(a int) string { return "" }, nil},
	} {
		out := Collect(Map(Iter(tc.a), tc.f))
		if !reflect.DeepEqual(out, tc.out) {
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
		{[]int{1, 2, 3, 4, 5}, 0, func(st int, v int) int { return st + v }, []int{1, 3, 6, 10, 15}},
		{[]int{1, 2, 3, 4, 5}, 1, func(st int, v int) int { return st + v }, []int{2, 4, 7, 11, 16}},
		{nil, 0, nil, nil},
	} {
		out := Collect(Scan(Iter(tc.sli), tc.st, tc.f))
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Testcase %d: got %v, expect %v.", i, out, tc.out)
		}
	}
}
