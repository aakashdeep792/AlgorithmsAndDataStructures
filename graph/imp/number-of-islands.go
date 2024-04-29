package graph

/*

200. Number of Islands

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

*/

func dfs(grid [][]byte, r, c int) {
	rows, cols := len(grid), len(grid[0])
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '2'

	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
}

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				ans++
				dfs(grid, r, c)
			}
		}
	}

	return ans
}
