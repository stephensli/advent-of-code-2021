package main

import (
	"strconv"
	"strings"
)

type Color string

var (
	red   Color = "red"
	green Color = "green"
	blue  Color = "blue"
)

type Session struct {
}

type Cube struct {
	count int64
	color Color
}

type Game struct {
	id    int64
	cubes []*Cube
}

func parseInput(lines []string) (games []*Game) {
	for _, line := range lines {
		lineSplit := strings.Split(line, ": ")

		id := strings.TrimPrefix(strings.ToLower(lineSplit[0]), "game ")
		gameID, _ := strconv.ParseInt(id, 10, 64)

		game := Game{
			id:    gameID,
			cubes: make([]*Cube, 0),
		}

		for _, games := range strings.Split(lineSplit[1], "; ") {
			for _, gameValues := range strings.Split(games, ", ") {
				cubes := strings.Split(gameValues, " ")
				cubeCount, _ := strconv.ParseInt(cubes[0], 10, 64)
				game.cubes = append(game.cubes, &Cube{
					count: cubeCount,
					color: Color(cubes[1]),
				})
			}
		}

		games = append(games, &game)
	}

	return games
}
