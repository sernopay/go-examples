package main

import "fmt"

func dfs(grid [][]byte, r int, c int) {
	nr := len(grid)
	nc := len(grid[0])

	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == 0 {
		return
	}

	grid[r][c] = 0
	dfs(grid, r-1, c) // up
	dfs(grid, r+1, c) // down
	dfs(grid, r, c-1) // left
	dfs(grid, r, c+1) // right
}

func numIslands(grid [][]byte) int {

	totalOfIslands := 0

	if len(grid) == 0 {
		return totalOfIslands
	}

	nr := len(grid)
	nc := len(grid[0])

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == 1 {
				totalOfIslands++
				dfs(grid, r, c)
			}
		}
	}

	return totalOfIslands
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	totalOfIslands := numIslands(grid)
	fmt.Println(totalOfIslands)
}
