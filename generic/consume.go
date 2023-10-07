package iter

func Last[T any](it Iterator[T]) (t T, ok bool) {
	t, ok = it.Next()
	if !ok {
		return
	}
	for {
		t1, ok := it.Next()
		if !ok {
			break
		}
		t = t1
	}
	return
}

func At[T any](it Iterator[T], i int) (t T, ok bool) {
	return Skip(it, i).Next()
}

func Reduce[T any](it Iterator[T], f func(T, T) T) (ans T) {
	ans, ok := it.Next()
	if !ok {
		return
	}
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		ans = f(ans, t)
	}
	return
}

func Fold[T, St any](it Iterator[T], st St, f func(St, T) St) St {
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		st = f(st, t)
	}
	return st
}

func Collect[T any](it Iterator[T]) (ans []T) {
	for {
		t, ok := it.Next()
		if !ok {
			break
		}
		ans = append(ans, t)
	}
	return
}
