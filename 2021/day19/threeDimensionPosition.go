package main

type ThreeDimensionPosition struct {
	x, y, z int64
}

// equal returns true if the provided threeDimensionPosition values
// match the current one, (x,y,z,) all equal.
func (p ThreeDimensionPosition) equal(sp ThreeDimensionPosition) bool {
	return p.x == sp.x && p.y == sp.y && p.z == sp.z
}

// rotate around any one of the possible 24 positions
//
// rotation must be less than or equal to 24 otherwise its just going to
// continue rotating around those positions again.
func (p ThreeDimensionPosition) rotate(rotation int) ThreeDimensionPosition {
	var x, y, z = p.x, p.y, p.z

	// rotate the coordinates' system so that the x-axis points in the
	// possible 6 directions.
	switch rotation % 6 {
	case 0:
		x, y, z = x, y, z
		break
	case 1:
		x, y, z = -x, y, -z
		break
	case 2:
		x, y, z = y, -x, z
		break
	case 3:
		x, y, z = -y, x, z
		break
	case 4:
		x, y, z = z, y, -x
		break
	case 5:
		x, y, z = -z, y, x
		break
	}

	// rotate around the x-axis
	switch (rotation / 6) % 4 {
	case 0:
		x, y, z = x, y, z
		break
	case 1:
		x, y, z = x, -z, y
		break
	case 2:
		x, y, z = x, y, -z
		break
	case 3:
		x, y, z = x, z, -y
		break
	}

	return ThreeDimensionPosition{x, y, z}
}
