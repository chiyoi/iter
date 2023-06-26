package iter

import (
	"reflect"
)

const (
	Continue = true
	Break    = false
)

type Iterator struct {
	Next func() (v any)
}

func Iter(sli any) (it Iterator) {
	rs := reflect.ValueOf(sli)
	if rs.Kind() != reflect.Slice {
		panic("sli should be a slice or a string")
	}

	return Iterator{func() (v any) {
		if rs.Len() > 0 {
			v, rs = rs.Index(0).Interface(), rs.Slice(1, rs.Len())
		}
		return
	}}
}

func From(f func() (v any)) Iterator {
	return Iterator{f}
}

func Range(start, stop int) Iterator {
	st := start
	return From(func() (v any) {
		if st >= stop {
			return
		}
		v = st
		st++
		return
	})
}

func Empty() Iterator {
	return From(func() (v any) { return })
}

func Repeat(v any) Iterator {
	return From(func() any { return v })
}

func Zip(ita, itb Iterator) Iterator {
	return From(func() (v any) {
		a, b := ita.Next(), itb.Next()
		if a == nil || b == nil {
			return
		}
		return [...]any{a, b}
	})
}

func (it Iterator) Chain(it1 Iterator) Iterator {
	return From(func() (v any) {
		if v = it.Next(); v == nil {
			v = it1.Next()
		}
		return
	})
}

func (it Iterator) ForEach(f func(v any, i int) bool) {
	i := 0
	for {
		v := it.Next()
		if v == nil {
			return
		}
		if ok := f(v, i); !ok {
			return
		}
		i++
	}
}
