package main

type Scanner struct {
	id       int
	position ThreeDimensionPosition
	beacons  []ThreeDimensionPosition

	// if any applied rotation is made to all the beacons. since we don't know the
	// position in which the scanner is facing when to perform is made, we must
	// check all possible rotations to see if they have over lapping shared beacons.
	appliedRotation int
}

// GetBeaconsInWorld will adjust the beacons in which they are from
// the current scanners' location. e.g. if the current location is -1,1,1 and
// the beacon is 1,1,2 then the beacon will become (0,2,3)
func (s *Scanner) GetBeaconsInWorld() (beacons []ThreeDimensionPosition) {
	beacons = []ThreeDimensionPosition{}

	for _, beacon := range s.beacons {
		rotated := beacon.rotate(s.appliedRotation)
		beacons = append(beacons, ThreeDimensionPosition{
			x: s.position.x + rotated.x,
			y: s.position.y + rotated.y,
			z: s.position.z + rotated.z,
		})
	}

	return beacons
}

// rotate will create a new scanner and on creation the scanner
// will rotate all the beacon values by said amount.
func (s *Scanner) rotate() *Scanner {
	return newScanner(s.id, s.position, s.appliedRotation+1, s.beacons)
}

// setToLocation will create a new scanner with the location set
// to the provided, this will change the beacons' location when
// calling GetBeaconsInWorld.
func (s *Scanner) setToLocation(location ThreeDimensionPosition) *Scanner {
	return newScanner(s.id, location, s.appliedRotation, s.beacons)
}

// newScanner will create a new scanner and apply the cubic rotation to shift the given scanners
// local beacons to one of the 24 possible positions. Rotation of 0 will keep the current positions.
func newScanner(id int, position ThreeDimensionPosition, rotation int, scannedBeacons []ThreeDimensionPosition) *Scanner {
	return &Scanner{
		id:              id,
		position:        position,
		beacons:         scannedBeacons,
		appliedRotation: rotation,
	}
}
