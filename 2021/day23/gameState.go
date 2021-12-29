package main

import (
	"fmt"
	"strings"
)

type Amphipod uint

const (
	AmberEnergy   Amphipod = 1
	BronzeEnergy  Amphipod = 10
	CopperEnergy  Amphipod = 100
	DessertEnergy Amphipod = 1000
)

// TargetRoomIndex returns the target room index for the given Amphipod
func (a Amphipod) TargetRoomIndex() int {
	switch a {
	case AmberEnergy:
		return 0
	case BronzeEnergy:
		return 1
	case CopperEnergy:
		return 2
	case DessertEnergy:
		return 3
	}

	return -1
}

func getTargetAmphipodForRoomIndex(roomIndex int) Amphipod {
	switch roomIndex {
	case 0:
		return AmberEnergy
	case 1:
		return BronzeEnergy
	case 2:
		return CopperEnergy
	case 3:
		return DessertEnergy
	}

	return 0
}

type PositionState uint

const (
	// Empty This position does not contain anything.
	Empty PositionState = iota
	// Occupied this position currently contains a body
	Occupied
	// Banned this position is banned, e.g. doorway.
	Banned
)

// Position A single position held within the game plane.
type Position struct {
	id       uint
	State    PositionState
	Amphipod Amphipod
}

// Room is a place which can hold Amphipods
type Room []Position

// Full returns true if all places within the room is full
// and not in any of the empty states.
func (r Room) Full() bool {
	for i := 0; i < len(r); i++ {
		if r[i].State == Empty {
			return false
		}
	}
	return true
}

// Empty returns true if all values are not occupied.
func (r Room) Empty() bool {
	for i := 0; i < len(r); i++ {
		if r[i].State != Empty {
			return false
		}
	}
	return true
}

func (r Room) Top() Position {
	if len(r) <= 0 {
		return Position{State: Empty}
	}

	return r[len(r)-1]
}

func (r Room) ContainsOnlyTarget(targetAmphipod Amphipod) bool {
	for _, position := range r {
		if position.Amphipod != targetAmphipod {
			return false
		}
	}

	return true
}

// InWinningState will return true if all values within the rooms are occupied,
// and they all meet the provided target targetAmphibious
func (r Room) InWinningState(targetAmphipod Amphipod, targetDepth int) bool {
	targetCount := 0

	for _, position := range r {
		if position.State == Occupied && position.Amphipod == targetAmphipod {
			targetCount += 1
		}
	}

	return targetCount == targetDepth
}

// RoomsInWinningState returns true if all rooms are in the winning state.
func RoomsInWinningState(rooms []Room, targetDepth int) bool {
	return rooms[0].InWinningState(AmberEnergy, targetDepth) &&
		rooms[1].InWinningState(BronzeEnergy, targetDepth) &&
		rooms[2].InWinningState(CopperEnergy, targetDepth) &&
		rooms[3].InWinningState(DessertEnergy, targetDepth)
}

// GameState The State of the game at any given point, this will be used by the Dijkstra
// algorithm to determine if we should be stopping at any point. The two points
// are:
//
// startState -> complete State
//
// Using Dijkstra to locate that cost
type GameState struct {
	CurrentCost  int
	RoomCapacity int
	Hallway      []Position
	Rooms        []Room
	history      []*GameState
}

// HallWayEmpty returns true if and only if the hallway is empty.
func (g GameState) HallWayEmpty() bool {
	for _, position := range g.Hallway {
		if position.State == Occupied {
			return false
		}
	}

	return true
}

// Key generate a unique key for the given game state for caching
func (g GameState) Key() string {
	var key []string

	for _, position := range g.Hallway {
		key = append(key, fmt.Sprintf("%d-%d", position.State, position.Amphipod))
	}

	for _, room := range g.Rooms {
		for _, position := range room {
			key = append(key, fmt.Sprintf("%d-%d", position.State, position.Amphipod))
		}
	}

	key = append(key, fmt.Sprintf("%d", g.CurrentCost))

	return strings.Join(key, "-")
}

func (g GameState) print(ignoreChildren bool) {
	if !ignoreChildren {
		for _, state := range g.history {
			state.print(false)
		}
	}

	fmt.Println("-----------------")

	for i, position := range g.Hallway {
		if i == 0 || i == len(g.Hallway)-1 {
			fmt.Print("#")
			continue
		}

		if position.State != Occupied {
			fmt.Print(".")
			continue
		}

		switch position.Amphipod {
		case AmberEnergy:
			fmt.Print("A")
			break
		case BronzeEnergy:
			fmt.Print("B")
			break
		case CopperEnergy:
			fmt.Print("C")
			break
		case DessertEnergy:
			fmt.Print("D")
			break
		}
	}

	fmt.Println()

	for i := g.RoomCapacity - 1; i >= 0; i-- {
		fmt.Print("###")

		for j := 0; j < len(g.Rooms); j++ {
			if len(g.Rooms[j]) <= i {
				fmt.Print(".#")
				continue
			}

			switch g.Rooms[j][i].Amphipod {
			case AmberEnergy:
				fmt.Print("A#")
				break
			case BronzeEnergy:
				fmt.Print("B#")
				break
			case CopperEnergy:
				fmt.Print("C#")
				break
			case DessertEnergy:
				fmt.Print("D#")
				break
			}
		}

		fmt.Print("#")
		fmt.Println()
	}

	fmt.Println("cost:", g.CurrentCost)
}

func (g GameState) Clone() GameState {
	gameState := GameState{
		CurrentCost:  g.CurrentCost,
		RoomCapacity: g.RoomCapacity,
		Hallway:      []Position{},
		Rooms:        []Room{},
		history:      g.history,
	}

	for _, position := range g.Hallway {
		gameState.Hallway = append(gameState.Hallway, Position{
			State:    position.State,
			Amphipod: position.Amphipod,
			id:       position.id,
		})
	}

	for _, room := range g.Rooms {
		newRoom := Room{}

		for _, position := range room {
			newRoom = append(newRoom, Position{
				State:    position.State,
				Amphipod: position.Amphipod,
				id:       position.id,
			})
		}

		gameState.Rooms = append(gameState.Rooms, newRoom)
	}

	return gameState
}
