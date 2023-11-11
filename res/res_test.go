package res

import (
	"fmt"
	"strconv"
	"testing"
)

func TestResult(t *testing.T) {
	for i, tc := range []struct {
		in        int
		transform func(int) (int, error)
		out       int
		err       error
	}{
		{
			1,
			func(i int) (int, error) {
				b, err := Then(i, nil, func(a int) (b string, err error) {
					return fmt.Sprint(a + 1), nil
				})
				c, err := Then(b, err, strconv.Atoi)
				return Then(c, err, func(c int) (d int, err error) {
					return c + 1, nil
				})
			},
			3,
			nil,
		},
		{
			1,
			func(i int) (int, error) {
				a, err := And[int](nil, func() (int, error) {
					return 2, nil
				})
				b, err := And[int](err, func() (int, error) {
					return 3, nil
				})
				return And[int](err, func() (int, error) {
					return i + a + b, nil
				})
			},
			6,
			nil,
		},
	} {
		out, err := tc.transform(tc.in)
		if out != tc.out || err != tc.err {
			t.Errorf("Testcase %d: out (%v, %v), expect (%v, %v).", i, out, err, tc.out, tc.err)
		}
	}
}
