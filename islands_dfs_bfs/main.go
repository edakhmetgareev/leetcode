package main

import "fmt"

func main() {
	case1 := [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
	}
	case2 := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	}

	fmt.Println("case1: ", numIslands(case1))
	fmt.Println("case2: ", numIslands(case2))
}

func numIslands(grid [][]string) int {
	var res int
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			if grid[i][j] == "1" {
				res++
				removeIsland(grid, i, j)
			}
		}
	}

	return res
}

func removeIsland(grid [][]string, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != "1" {
		return
	}

	grid[i][j] = "0"

	removeIsland(grid, i-1, j)
	removeIsland(grid, i+1, j)
	removeIsland(grid, i, j+1)
	removeIsland(grid, i, j-1)
}
