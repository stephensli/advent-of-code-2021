package helpers

func SetIfMissing[T, K comparable](val map[T]K, index T, def K) {
	if _, ok := val[index]; !ok {
		val[index] = def
	}
}
