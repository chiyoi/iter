package iter

func Fold[Elem any, St any](sli []Elem, st St, f func(St, Elem) St) St {
	for _, elem := range sli {
		st = f(st, elem)
	}
	return st
}

func Scan[Elem any, St any](sli []Elem, st St, f func(St, Elem) St) (ans []St) {
	ans = make([]St, 0, len(sli))
	for _, elem := range sli {
		st = f(st, elem)
		ans = append(ans, st)
	}
	return
}

func Filter[Elem any](sli []Elem, pred func(Elem) bool) (ans []Elem) {
	for _, elem := range sli {
		if pred(elem) {
			ans = append(ans, elem)
		}
	}
	return
}

func Map[A, B any](sli []A, f func(A) B) (ans []B) {
	ans = make([]B, 0, len(sli))
	for _, elem := range sli {
		ans = append(ans, f(elem))
	}
	return
}

func Reduce[Elem any](sli []Elem, f func(Elem, Elem) Elem) (ans Elem) {
	if len(sli) == 0 {
		return
	}
	ans = sli[0]
	for _, elem := range sli[1:] {
		ans = f(ans, elem)
	}
	return
}

func Range(a, b int) (ans []int) {
	ans = make([]int, 0, b-a)
	for i := a; i < b; i++ {
		ans = append(ans, i)
	}
	return
}

func Zip[A, B any](a []A, b []B) (ans []ZipItem[A, B]) {
	ans = make([]ZipItem[A, B], 0, min(len(a), len(b)))
	for i := range a {
		if i >= len(b) {
			break
		}
		ans = append(ans, ZipItem[A, B]{a[i], b[i]})
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type ZipItem[A, B any] struct {
	A A
	B B
}
