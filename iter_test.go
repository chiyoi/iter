package iter

import (
	"strconv"
	"testing"

	"github.com/chiyoi/apricot/test"
)

func TestMap(t *testing.T) {
	for i, tc := range []struct {
		a   []int
		f   func(int) (string, error)
		out []string
		err error
	}{
		{
			[]int{1, 2, 3, 4, 5},
			func(a int) (string, error) {
				return strconv.Itoa(a + 1), nil
			},
			[]string{"2", "3", "4", "5", "6"},
			nil,
		},
		{nil, nil, nil, nil},
		{
			nil,
			func(a int) (string, error) {
				return "", nil
			},
			nil,
			nil,
		},
	} {
		it := Iter(tc.a)
		m := Map(it, tc.f)
		out, err := Collect(m)
		test.CheckEqual(t, i, "err", err, nil)
		test.CheckDeepEqual(t, i, "out", out, tc.out)
	}
}
