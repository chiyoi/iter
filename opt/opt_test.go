package opt

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOption(t *testing.T) {
	for i, tc := range []struct {
		in        int
		transform func(int) int
		out       int
	}{
		{
			1,
			func(i int) int {
				a := Then(i, func(i int) string {
					return fmt.Sprint(i + 1)
				})
				b := Then(a, func(s string) int {
					i, err := strconv.Atoi(s)
					if err != nil {
						return 0
					}
					return i + 1
				})
				return b
			},
			3,
		},
	} {
		out := tc.transform(tc.in)
		if out != tc.out {
			t.Errorf("Testcase %d: out %v, expect %v.", i, out, tc.out)
		}
	}
}
