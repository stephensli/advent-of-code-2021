package main

import "github.com/stephensli/aoc/helpers/file"

func parseInput(filePath string) GameState {
	lines := file.ToTextSplit(filePath, "")[1:]

	var hallway []Position
	rooms := make([]Room, 4)

	for i := 0; i < len(rooms); i++ {
		rooms[i] = []Position{}
	}

	// fill in the Hallway.
	for _, value := range lines[0] {
		switch value {
		case "#":
			hallway = append(hallway, Position{State: Banned})
			break
		case "A":
			hallway = append(hallway, Position{State: Occupied, Amphipod: AmberEnergy})
			break
		case "B":
			hallway = append(hallway, Position{State: Occupied, Amphipod: BronzeEnergy})
			break
		case "C":
			hallway = append(hallway, Position{State: Occupied, Amphipod: CopperEnergy})
			break
		case "D":
			hallway = append(hallway, Position{State: Occupied, Amphipod: DessertEnergy})
			break
		case ".":
			hallway = append(hallway, Position{State: Empty})
			break
		}

	}

	// these are doorways into a room,  which means they are banned from entry
	// and should be marked as so, allowing the algorithm to continue.
	for _, value := range []int{3, 5, 7, 9} {
		hallway[value].State = Banned
	}

	roomCapacity := 0

	// for the remainder of the lines, check the room index values to determine
	// which Amphipod is located within to build the rooms up.
	// last line is a base.
	for i := len(lines) - 2; i >= 1; i-- {
		line := lines[i]
		roomCapacity += 1

		for _, roomIdx := range []int{3, 5, 7, 9} {
			switch line[roomIdx] {
			case "A":
				rooms[(roomIdx-1)/2-1] = append(rooms[(roomIdx-1)/2-1],
					Position{State: Occupied, Amphipod: AmberEnergy})
				break
			case "B":
				rooms[(roomIdx-1)/2-1] = append(rooms[(roomIdx-1)/2-1],
					Position{State: Occupied, Amphipod: BronzeEnergy})
				break
			case "C":
				rooms[(roomIdx-1)/2-1] = append(rooms[(roomIdx-1)/2-1],
					Position{State: Occupied, Amphipod: CopperEnergy})
				break
			case "D":
				rooms[(roomIdx-1)/2-1] = append(rooms[(roomIdx-1)/2-1],
					Position{State: Occupied, Amphipod: DessertEnergy})
				break
			}
		}
	}

	return GameState{
		CurrentCost:  0,
		RoomCapacity: roomCapacity,
		Hallway:      hallway,
		Rooms:        rooms,
	}
}
