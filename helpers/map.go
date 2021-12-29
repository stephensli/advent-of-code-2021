package helpers

func SetIfMissing(val map[string]int64, index string, def int64) {
	if _, ok := val[index]; !ok {
		val[index] = def
	}
}
