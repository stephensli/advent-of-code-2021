package main

type assigment struct {
	rangeLeft  int
	rangeRight int
}

// Contains returns true if and only if the provided pair can be directly included inside the
// current assignment range.
func (a assigment) Contains(pair assigment) bool {
	return pair.rangeLeft >= a.rangeLeft && pair.rangeRight <= a.rangeRight
}

// Overlap returns true if and only if the provided pair can be directly included inside the
// current assignment range on either end.
func (a assigment) Overlap(pair assigment) bool {
	if pair.rangeLeft > a.rangeRight {
		return false
	}

	if pair.rangeRight < a.rangeLeft {
		return false
	}

	return pair.rangeLeft >= a.rangeLeft || pair.rangeRight <= a.rangeRight
}
