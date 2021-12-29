package main

import (
	"fmt"
	"strings"
)

func parseInput(lines []string) []Action {
	actions := []Action{}

	for _, line := range lines {
		var x1, y1, x2, y2 int

		if strings.HasPrefix(line, "toggle") {
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			actions = append(actions, Action{
				From:   Position{x1, y1},
				To:     Position{x2, y2},
				Toggle: true,
				OnOff:  false,
			})
		} else {
			var onOff string
			fmt.Sscanf(line, "turn %s %d,%d through %d,%d", &onOff, &x1, &y1, &x2, &y2)
			actions = append(actions, Action{
				From:   Position{x1, y1},
				To:     Position{x2, y2},
				Toggle: false,
				OnOff:  onOff == "on",
			})
		}

	}

	return actions
}
