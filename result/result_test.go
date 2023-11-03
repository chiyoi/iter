package result

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
				a := Result(i, nil)
				b := Try(a, func(a int) (b string, err error) {
					return fmt.Sprintf("%d", a+1), nil
				})
				c := Try(b, strconv.Atoi)
				d := Try(c, func(c int) (d int, err error) {
					return c + 1, nil
				})
				return Values(d)
			},
			3, nil,
		},
	} {
		out, err := tc.transform(tc.in)
		if out != tc.out || err != tc.err {
			t.Errorf("Testcase %d: out (%v, %v), expect (%v, %v).", i, out, err, tc.out, tc.err)
		}
	}
}
