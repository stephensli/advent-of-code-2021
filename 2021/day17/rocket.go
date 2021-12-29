package main

type Rocket struct {
	target           *Target
	x, y, velX, velY int
	deltaX, deltaY   int

	maxY int
	maxX int
}

func (r *Rocket) madeTarget() bool {
	for !r.missedTarget(r.target) {
		r.x += r.velX
		r.y += r.velY

		// if x has already hit zero, then it will only go down.
		if r.velX == 0 {
			r.deltaX = 0
		}

		r.velX += r.deltaX
		r.velY += r.deltaY

		if r.x > r.maxX {
			r.maxX = r.x
		}

		if r.y > r.maxY {
			r.maxY = r.y
		}

		if r.target.contains(r) {
			return true
		}
	}

	return false
}

func (t *Target) contains(r *Rocket) bool {
	return r.x <= t.xMax && r.x >= t.xMin &&
		r.y <= t.yMax && r.y >= t.yMin
}

func (r *Rocket) missedTarget(t *Target) bool {
	return r.x > t.xMax || r.y < t.yMin
}
