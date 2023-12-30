package maths

// GreatestCommonDivisor finds the greatest common divisor (GreatestCommonDivisor) via Euclidean algorithm
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LowestCommonMultiple finds the Least Common Multiple (LowestCommonMultiple) via GCD
func LowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LowestCommonMultiple(result, integers[i])
	}

	return result
}

func Minimum[T int | int64 | float32 | float64](a, b T) T {
	if a > b {
		return b
	}

	return a
}
