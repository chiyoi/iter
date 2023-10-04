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
