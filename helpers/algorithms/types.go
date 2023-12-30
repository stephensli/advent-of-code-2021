package algorithms

type Direction int

var (
	UpDirection    Direction = 0
	LeftDirection  Direction = 1
	RightDirection Direction = 3
	DownDirection  Direction = 4

	UpAndLeftDireciton  Direction = 5
	UpAndRightDireciton Direction = 6

	DownAndLeftDireciton  Direction = 7
	DownAndRIghtDireciton Direction = 8
)

var AllDirections = []Direction{
	UpDirection,
	LeftDirection,
	RightDirection,
	DownDirection,
	UpAndLeftDireciton,
	UpAndRightDireciton,
	DownAndLeftDireciton,
	DownAndRIghtDireciton,
}

var NonDigagnonalDirections = []Direction{
	UpDirection,
	LeftDirection,
	RightDirection,
	DownDirection,
}

var DigagnonalDirections = []Direction{
	UpAndLeftDireciton,
	UpAndRightDireciton,
	DownAndLeftDireciton,
	DownAndRIghtDireciton,
}

type Coords struct {
	X, Y int
}
