package iter

import (
	"errors"

	"github.com/chiyoi/iter/res"
)

var (
	ErrStopIteration   = errors.New("stop iteration")
	IsErrStopIteration = func(err error) bool {
		return errors.Is(err, ErrStopIteration)
	}
)

func Iter[A any](sli []A) func() (A, error) {
	i := 0
	return func() (a A, err error) {
		if i < len(sli) {
			a = sli[i]
			i++
			return
		}
		err = ErrStopIteration
		return
	}
}

func Map[A, B any](it func() (A, error), f func(A) (B, error)) func() (B, error) {
	return func() (b B, err error) {
		a, err := it()
		return res.R(a, err, f)
	}
}

func Collect[A any](it func() (A, error)) (ans []A, err error) {
	for {
		var a A
		a, err = it()
		if err != nil {
			if IsErrStopIteration(err) {
				err = nil
			}
			return
		}
		ans = append(ans, a)
	}
}
