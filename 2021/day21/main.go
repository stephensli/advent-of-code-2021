package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers/aoc"
	"github.com/stephensli/advent-of-code-2021/helpers/cache"
	"github.com/stephensli/advent-of-code-2021/helpers/file"
	"math"
)

func valueOfDeterministicDie(current, min, max int) (value int, newDie int) {
	value = 0

	for i := 0; i < 3; i++ {

		value += current
		current += 1

		if current > max {
			current = min
		}
	}

	return value, current
}

func getPlayerNextPosition(current, dice, boardSize int) int {
	next := (current + dice) % boardSize

	if next == 0 {
		return boardSize
	}

	return next
}

type gameState struct {
	playerOnePos, playerTwoPos     int
	playerOneScore, playerTwoScore int64
}

func quantumDice(state gameState, diceValues []int, cache cache.SimpleCache[gameState, []int64]) []int64 {
	if val, ok := cache.Get(state); ok {
		return val
	}

	if state.playerOneScore >= 21 {
		value := []int64{1, 0}
		cache.Set(state, value)
		return value
	}

	if state.playerTwoScore >= 21 {
		value := []int64{0, 1}
		cache.Set(state, value)
		return value
	}

	score := []int64{0, 0}

	for _, dice := range diceValues {
		newPos := getPlayerNextPosition(state.playerOnePos, dice, 10)
		newScore := state.playerOneScore + int64(newPos)

		// flip the game role
		winners := quantumDice(gameState{
			playerOnePos:   state.playerTwoPos,
			playerOneScore: state.playerTwoScore,
			playerTwoPos:   newPos,
			playerTwoScore: newScore,
		}, diceValues, cache)

		score[0] += winners[1]
		score[1] += winners[0]
	}

	cache.Set(state, score)
	return score
}

func partTwo(p1, p2 int) {
	diceValues := []int{}

	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				diceValues = append(diceValues, i+j+k)
			}
		}
	}

	//because of a memory limit, instead lets do 1x at a time.
	//p1, p2, turn
	c := cache.New[gameState, []int64]()

	result := quantumDice(gameState{
		playerOnePos:   p1,
		playerTwoPos:   p2,
		playerOneScore: 0,
		playerTwoScore: 0,
	}, diceValues, c)

	if result[0] > result[1] {
		aoc.PrintAnswer(2, result[0])
	} else {
		aoc.PrintAnswer(2, result[1])
	}
}

func partOne(playerOnePosition, playerTwoPosition int) {
	turnCount := 0
	scores := []int{0, 0}
	die := 1

	for scores[0] < 1000 && scores[1] < 1000 {
		// start by incrementing the player turn. we dont' want to
		// increment after the turn is over since this will result
		// in an off by one error.
		turnCount += 1

		// start on the first player as the turn. true being player 1, false
		// being player two.
		playerTurn := turnCount%2 == 1

		diceValue, newDice := valueOfDeterministicDie(die, 1, 100)
		die = newDice

		if playerTurn {
			playerOnePosition = getPlayerNextPosition(playerOnePosition, diceValue, 10)
			scores[0] += playerOnePosition
		} else {
			playerTwoPosition = getPlayerNextPosition(playerTwoPosition, diceValue, 10)
			scores[1] += playerTwoPosition
		}

	}

	aoc.PrintAnswer(1, turnCount*3*int(math.Min(float64(scores[0]), float64(scores[1]))))
}

func main() {
	defer aoc.Setup(2021, 21)()

	lines := file.ToTextLines("./input-example.txt")

	var playerOnePosition, playerTwoPosition int

	_, _ = fmt.Sscanf(lines[0], "Player 1 starting position: %d", &playerOnePosition)
	_, _ = fmt.Sscanf(lines[1], "Player 2 starting position: %d", &playerTwoPosition)

	partOne(playerOnePosition, playerTwoPosition)
	partTwo(playerOnePosition, playerTwoPosition)
}
