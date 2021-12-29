package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount    int
}

func main() {
	file, err := os.Open("./day2/input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []command

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(line[1])

		input = append(input, command{
			direction: line[0],
			amount:    num,
		})
	}

	fmt.Println("first", first(input))
	fmt.Println("second", second(input))
}

func first(commands []command) int {
	directionMap := map[string]*command{}

	for _, c := range commands {
		if val, ok := directionMap[c.direction]; ok {
			(*val).amount += c.amount
		} else {
			val := command{direction: c.direction, amount: c.amount}
			directionMap[val.direction] = &val
		}
	}

	return directionMap["forward"].amount * (directionMap["down"].amount - directionMap["up"].amount)
}

func second(commands []command) int {
	directionMap := map[string]*command{}

	directionMap["aim"] = &command{
		direction: "aim",
		amount:    0,
	}

	directionMap["depth"] = &command{
		direction: "depth",
		amount:    0,
	}

	for _, c := range commands {
		if val, ok := directionMap[c.direction]; ok {
			(*val).amount += c.amount
		} else {
			val := command{direction: c.direction, amount: c.amount}
			directionMap[val.direction] = &val
		}

		switch c.direction {
		case "forward":
			{
				directionMap["depth"].amount += c.amount * directionMap["aim"].amount
			}
		case "down":
			{
				directionMap["aim"].amount += c.amount
				break
			}
		case "up":
			{
				directionMap["aim"].amount -= c.amount
				break
			}

		}
	}

	return directionMap["forward"].amount * directionMap["depth"].amount
}
