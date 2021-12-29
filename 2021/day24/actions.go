package main

func add(a, b *int) (ok bool) {
	*a = (*b) + (*a)
	return true
}

func mul(a, b *int) (ok bool) {
	*a = (*a) * (*b)
	return true
}

func div(a, b *int) (ok bool) {
	if *b == 0 {
		return false
	}

	*a = (*a) / (*b)
	return true
}

func mod(a, b *int) (ok bool) {
	if *a < 0 || *b <= 0 {
		return false
	}

	*a = (*a) % (*b)
	return true
}

func eql(a, b *int) (ok bool) {
	if (*a) == (*b) {
		*a = 1
	} else {
		*a = 0
	}
	return true
}
