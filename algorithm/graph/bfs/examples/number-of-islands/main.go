package main

import (
	"fmt"

	queue "github.com/sernopay/go-examples/data-structure/queue/queue_dynamic_singly_linked_list/queue"
)

func numIslands(grid [][]string) int {
	if len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	numOfIslands := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == "1" {
				numOfIslands++
				grid[r][c] = "0" // mark visited
				neighbors := queue.New[int]()
				neighbors.Add(r*nc + c)
				for !neighbors.IsEmpty() {
					id, _ := neighbors.Poll()
					row := id / nc
					col := id % nc

					// up
					if row-1 >= 0 && grid[row-1][col] == "1" {
						neighbors.Add((row-1)*nc + col)
						grid[row-1][col] = "0"
					}
					// down
					if row+1 < nr && grid[row+1][col] == "1" {
						neighbors.Add((row+1)*nc + col)
						grid[row+1][col] = "0"
					}
					// left
					if col-1 >= 0 && grid[row][col-1] == "1" {
						neighbors.Add(row*nc + col - 1)
						grid[row][col-1] = "0"
					}
					// right
					if col+1 < nc && grid[row][col+1] == "1" {
						neighbors.Add(row*nc + col + 1)
						grid[row][col+1] = "0"
					}
				}
			}
		}
	}

	return numOfIslands
}

func main() {
	grid := [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "1"},
	}
	totalOfIslands := numIslands(grid)
	fmt.Println(totalOfIslands)
}
