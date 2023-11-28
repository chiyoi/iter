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

type Mapping[A, B any] func(A) (B, error)

type Folding[A, B any] func(B) func(A) (B, error)

type Iterator[A any] func() (A, error)

func Iter[A any](sli []A) Iterator[A] {
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

func Map[A, B any](it Iterator[A], f Mapping[A, B]) Iterator[B] {
	return func() (b B, err error) {
		a, err := it()
		return res.R(a, err, f)
	}
}

func Collect[A any](it Iterator[A]) (ans []A, err error) {
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
