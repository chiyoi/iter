package opt

import (
	"testing"

	"github.com/chiyoi/iter/cmb"
)

func TestOr(t *testing.T) {
	for i, tc := range []struct {
		in        int
		transform func(int) int
		out       int
	}{
		{
			0,
			func(i int) int {
				a := Or(i, cmb.Literal(1))
				b := Or(a, cmb.Literal(2))
				return b
			},
			1,
		},
		{
			0,
			func(i int) int {
				a := Or(i, cmb.Literal(0))
				b := Or(a, cmb.Literal(1))
				return b
			},
			1,
		},
		{
			1,
			func(i int) int {
				a := Or(i, cmb.Literal(2))
				b := Or(a, cmb.Literal(3))
				return b
			},
			1,
		},
	} {
		out := tc.transform(tc.in)
		if out != tc.out {
			t.Errorf("Testcase %d: out %v, expect %v.", i, out, tc.out)
		}
	}
}
