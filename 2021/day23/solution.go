package main

import (
	"container/heap"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/numbers"
	"github.com/stephensli/aoc/helpers/queue"
	"math"
)

type RoomStateCacheKey struct {
	Hallway []Position
	Rooms   []Room
}

func solveForInput(startingGameState GameState) int {
	// dijkstra algorithm

	// 1. for every single amphipod that is within a room  that is not its target room
	// or is the target room but other invalid types are below it. Move it into every
	// single space available
	//
	// (each space creates a new PATH)
	//
	// ensure to increment the cost for that move action.
	//
	// 2. For every single one in the all way, if CAN move into target room
	// 	  which meets the room requirements, then move

	minPriorityQueue := queue.MinPriorityQueue{{
		Value:    startingGameState,
		Priority: startingGameState.CurrentCost,
		Index:    0,
	}}

	// cache to check if we have already seen this one before.
	roomStateCache := cache.New[string, bool]()
	heap.Init(&minPriorityQueue)

	// current winning score
	winningScore := math.MaxInt

	for len(minPriorityQueue) != 0 {
		queueItem := heap.Pop(&minPriorityQueue).(*queue.Item)
		gameState := queueItem.Value.(GameState)

		// check to see if we have yet to beat our score. Since there is no
		// reason to check the cache if the value is already equal or higher
		// than our current winner.
		if gameState.CurrentCost >= winningScore {
			continue
		}

		// if we are in a winning state. Then go and set our winner and
		// continue. No need to check the cache if we are a winner since
		// this has yet to be seen. (hallway must be empty)
		if gameState.CurrentCost < winningScore && RoomsInWinningState(gameState.Rooms, gameState.RoomCapacity) {
			winningScore = gameState.CurrentCost
		}

		// have we seen this state of game before? if we have seen it before,
		// and it's the best room, then the score is already set, so we can just
		// skip it.
		if roomStateCache.Has(gameState.Key()) {
			continue
		}

		// mark as seen before we go and generate the new game states. since
		// we don't need to go and check this current state again if it comes
		// to it.
		roomStateCache.Set(gameState.Key(), true)

		// first handle all values within the hallway.
		for hallwayPositionIndex, hallwayPosition := range gameState.Hallway {
			if hallwayPosition.State != Occupied {
				continue
			}

			targetRoomNumber := hallwayPosition.Amphipod.TargetRoomIndex()
			targetRoom := gameState.Rooms[targetRoomNumber]

			// if the room is not full and all values within the room are matching the
			// target amphipod, then we can use this room and fill it with our current.
			//
			// This does not mean its reachable, since another amphipod could be blocking the path
			if targetRoom.Empty() || targetRoom.ContainsOnlyTarget(hallwayPosition.Amphipod) {
				targetIndex := (targetRoomNumber+1)*2 + 1
				reachable := true

				// check the hallway to see if anything is blocking the path into our target room.
				//
				// it can be banned since we are checking for the extract room hallwayPosition.
				// and banned is okay since we will be pushing into a room and not this
				// expect pos
				if hallwayPositionIndex < targetIndex {
					// moving forward down the hallway into the hallwayPosition
					for hallWayPos := hallwayPositionIndex + 1; hallWayPos <= targetIndex; hallWayPos++ {
						if gameState.Hallway[hallWayPos].State == Occupied {
							reachable = false
						}
					}
				} else {
					// moving backwards down the hallway into the hallwayPosition
					for hallWayPos := hallwayPositionIndex - 1; hallWayPos >= targetIndex; hallWayPos-- {
						if gameState.Hallway[hallWayPos].State == Occupied {
							reachable = false
						}
					}

				}

				// if its reachable, then determine the total number of steps
				// required to reach said hallwayPosition and determine the score.
				// increase the score to the new game state score value.
				if reachable {
					totalSteps := numbers.Abs[int]((targetRoomNumber+1)*2+1-hallwayPositionIndex) +
						(gameState.RoomCapacity - len(targetRoom))

					additionalCost := totalSteps * int(hallwayPosition.Amphipod)

					// clone our game state and then update our value.
					// since we have determined a new hallwayPosition game state
					// this must be pushed into the heap for the next iteration.
					newGameState := gameState.Clone()

					// set the old hallwayPosition as empty and the new hallwayPosition as occupied.
					newGameState.CurrentCost += additionalCost
					newGameState.Hallway[hallwayPositionIndex] = Position{State: Empty}
					newGameState.Rooms[targetRoomNumber] = append(newGameState.Rooms[targetRoomNumber], Position{
						State:    hallwayPosition.State,
						Amphipod: hallwayPosition.Amphipod,
					})

					// newGameState.history = append(newGameState.history, &gameState)
					heap.Push(&minPriorityQueue, &queue.Item{Value: newGameState, Priority: additionalCost})
				}
			}
		}

		// now handle all values within rooms which are not in a winning state.
		for roomIndex, room := range gameState.Rooms {
			// if it's in a winning state already, don't go and move them since
			// it will only cause a larger score putting them back in again.
			roomTargetAmphipod := getTargetAmphipodForRoomIndex(roomIndex)
			roomTop := room.Top()

			if roomTop.State == Empty {
				continue
			}

			if room.InWinningState(roomTargetAmphipod, gameState.RoomCapacity) || room.ContainsOnlyTarget(roomTargetAmphipod) {
				continue
			}

			// determine the current nodes' hallway hallwayPosition, and then we can check all values
			// back and forward until we hit another occupied space. Accepting all empty ones.
			hallWayPosIndex := (roomIndex+1)*2 + 1
			var validPositions []int

			// forward
			for i := hallWayPosIndex; i < len(gameState.Hallway); i++ {
				hallWayPosition := gameState.Hallway[i]

				if hallWayPosition.State == Empty {
					validPositions = append(validPositions, i)
					// make sure to break if we hit a occupied spot, since we cannot pass this
				} else if hallWayPosition.State != Banned {
					break
				}

			}

			// backwards
			for i := hallWayPosIndex; i >= 0; i-- {
				hallWayPosition := gameState.Hallway[i]

				if hallWayPosition.State == Empty {
					validPositions = append(validPositions, i)
					// make sure to break if we hit a occupied spot, since we cannot pass this
				} else if hallWayPosition.State != Banned {
					break
				}
			}

			// now for all valid positions clone the current game state a and push our
			// amphipod to that hallwayPosition, updating cost and empty the previous spot.
			// finally, pushing on the heap for checking.
			for _, validPositionIndex := range validPositions {
				newGameState := gameState.Clone()

				newGameStateRoom := newGameState.Rooms[roomIndex]
				movingAmphipod := newGameStateRoom.Top()

				// remove it from the room
				newGameState.Rooms[roomIndex] = newGameStateRoom[:len(newGameStateRoom)-1]

				// determine the cost of the entire travel.
				costToLeaveRoom := newGameState.RoomCapacity - len(room) + 1
				hallWayMoveCost := numbers.Abs[int](hallWayPosIndex - validPositionIndex)

				moveCost := (costToLeaveRoom + hallWayMoveCost) * int(movingAmphipod.Amphipod)
				newGameState.CurrentCost += moveCost

				// move it to the new hallwayPosition.
				newGameState.Hallway[validPositionIndex] = Position{
					Amphipod: movingAmphipod.Amphipod,
					State:    Occupied,
				}

				// newGameState.history = append(newGameState.history, &gameState)

				heap.Push(&minPriorityQueue, &queue.Item{
					Value:    newGameState,
					Priority: moveCost,
				})
			}
		}
	}

	return winningScore
}

func partOne(startingGameState GameState) {
	winningScore := solveForInput(startingGameState)
	aoc.PrintAnswer(1, winningScore)
}

func partTwo(startingGameState GameState) {
	startingGameState.Rooms[0] = []Position{
		startingGameState.Rooms[0][0],
		{State: Occupied, Amphipod: DessertEnergy},
		{State: Occupied, Amphipod: DessertEnergy},
		startingGameState.Rooms[0][1],
	}

	startingGameState.Rooms[1] = []Position{
		startingGameState.Rooms[1][0],
		{State: Occupied, Amphipod: BronzeEnergy},
		{State: Occupied, Amphipod: CopperEnergy},
		startingGameState.Rooms[1][1],
	}

	startingGameState.Rooms[2] = []Position{
		startingGameState.Rooms[2][0],
		{State: Occupied, Amphipod: AmberEnergy},
		{State: Occupied, Amphipod: BronzeEnergy},
		startingGameState.Rooms[2][1],
	}

	startingGameState.Rooms[3] = []Position{
		startingGameState.Rooms[3][0],
		{State: Occupied, Amphipod: CopperEnergy},
		{State: Occupied, Amphipod: AmberEnergy},
		startingGameState.Rooms[3][1],
	}

	startingGameState.RoomCapacity = 4
	winningScore := solveForInput(startingGameState)
	aoc.PrintAnswer(2, winningScore)
}

func main() {
	defer aoc.Setup(2021, 23)()

	startGameState := parseInput("./input.txt")
	// printers.JsonPrint(startGameState, true)

	// Rules
	// 1. They will never stop on the space immediately outside any room.
	// 	  They can if they continue moving.
	//
	// 2. They will never move from the Hallway into a room unless that room is
	//	  their room *AND* and room contains no one or only of their kind.
	//
	// 3. If its starting room it's not its own room, it can stay in that room
	// 	  until it leaves the room and cannot enter again.
	//
	// 4. Once it stops moving in the Hallway, it will stay in that location
	//    until it can enter its room.

	partOne(startGameState)
	partTwo(startGameState)
}
