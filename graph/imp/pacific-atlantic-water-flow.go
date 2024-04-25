package graph

/*
417. Pacific Atlantic Water Flow
Solved
Medium
Topics
Companies
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.



Example 1:


Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
Explanation: The following cells can flow to the Pacific and Atlantic oceans, as shown below:
[0,4]: [0,4] -> Pacific Ocean
       [0,4] -> Atlantic Ocean
[1,3]: [1,3] -> [0,3] -> Pacific Ocean
       [1,3] -> [1,4] -> Atlantic Ocean
[1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
       [1,4] -> Atlantic Ocean
[2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
       [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
[3,0]: [3,0] -> Pacific Ocean
       [3,0] -> [4,0] -> Atlantic Ocean
[3,1]: [3,1] -> [3,0] -> Pacific Ocean
       [3,1] -> [4,1] -> Atlantic Ocean
[4,0]: [4,0] -> Pacific Ocean
       [4,0] -> Atlantic Ocean
Note that there are other possible paths for these cells to flow to the Pacific and Atlantic oceans.
Example 2:

Input: heights = [[1]]
Output: [[0,0]]
Explanation: The water can flow from the only cell to the Pacific and Atlantic oceans.

*/

func dfs(height [][]int, visited [][]bool, i, j int) {
	if visited[i][j] {
		return
	}

	visited[i][j] = true
	// Up
	if i-1 >= 0 && height[i-1][j] >= height[i][j] {
		dfs(height, visited, i-1, j)
	}
	// down
	if i+1 < len(height) && height[i+1][j] >= height[i][j] {
		dfs(height, visited, i+1, j)
	}
	//left
	if j-1 >= 0 && height[i][j-1] >= height[i][j] {
		dfs(height, visited, i, j-1)
	}
	// right
	if j+1 < len(height[0]) && height[i][j+1] >= height[i][j] {
		dfs(height, visited, i, j+1)
	}
}

func PacificAtlantic(heights [][]int) [][]int {
	// create the visited array
	rows, cols := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	// row
	for i := 0; i < rows; i++ {
		dfs(heights, pacific, i, 0)
		dfs(heights, atlantic, i, cols-1)
	}

	// cols
	for j := 0; j < cols; j++ {
		dfs(heights, pacific, 0, j)
		dfs(heights, atlantic, rows-1, j)
	}

	var ans [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}
