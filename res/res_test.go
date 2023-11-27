package res

import (
	"strconv"
	"testing"

	"github.com/chiyoi/apricot/test"
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
			func(x int) (int, error) {
				var err error
				y, err := R(x, err, runnerItoa)
				y += "1"
				x, err = R(y, err, strconv.Atoi)
				return R(x, err, runnerInc)
			},
			12,
			nil,
		},
		{
			1,
			func(x int) (int, error) {
				var err error
				y, err := R(x, err, runnerItoa)
				y += "nyan"
				x, err = R(y, err, strconv.Atoi)
				return R(x, err, runnerInc)
			},
			0,
			strconv.ErrSyntax,
		},
	} {
		out, err := tc.transform(tc.in)
		test.CheckEqual(t, i, "out", out, tc.out)
		test.CheckErrorIs(t, i, "err", err, tc.err)
	}
}

func runnerInc(x int) (int, error) {
	return x + 1, nil
}

func runnerItoa(x int) (string, error) {
	return strconv.Itoa(x), nil
}
