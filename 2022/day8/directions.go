package main

func checkLeft(grid [][]int, value, i, j int, distance int) (bool, int) {
	if j-1 < 0 {
		return true, distance
	}

	// If the next tree is larger than the target tree, then we are not visible
	// on this direction and can just return false and set the cache.
	if grid[i][j-1] >= value {
		return false, distance + 1
	}

	// otherwise let's go and check again with a recursive loop.
	return checkLeft(grid, value, i, j-1, distance+1)

}

func checkRight(grid [][]int, value int, i, j int, distance int) (bool, int) {
	// if we have hit the end, e.g past the last tree, then we must be visible
	// to the outside world, no tree is larger.
	if j+1 >= len(grid) {
		return true, distance
	}

	// If the next tree is larger than the target tree, then we are not visible
	// on this direction and can just return false and set the cache.
	if grid[i][j+1] >= value {
		return false, distance + 1
	}

	// otherwise let's go and check again with a recursive loop.
	return checkRight(grid, value, i, j+1, distance+1)
}

func checkUp(grid [][]int, value int, i, j int, distance int) (bool, int) {
	// if we have hit the end, e.g past the last tree, then we must be visible
	// to the outside world, no tree is larger.
	if i-1 < 0 {
		return true, distance
	}

	// If the next tree is larger than the target tree, then we are not visible
	// on this direction and can just return false and set the cache.
	if grid[i-1][j] >= value {
		return false, distance + 1
	}

	// otherwise let's go and check again with a recursive loop.
	return checkUp(grid, value, i-1, j, distance+1)
}

func checkDown(grid [][]int, value int, i, j int, distance int) (bool, int) {
	// if we have hit the end, e.g past the last tree, then we must be visible
	// to the outside world, no tree is larger.
	if i+1 >= len(grid) {
		return true, distance
	}

	// If the next tree is larger than the target tree, then we are not visible
	// on this direction and can just return false and set the cache.
	if grid[i+1][j] >= value {
		return false, distance + 1
	}

	// otherwise let's go and check again with a recursive loop.
	return checkDown(grid, value, i+1, j, distance+1)
}
