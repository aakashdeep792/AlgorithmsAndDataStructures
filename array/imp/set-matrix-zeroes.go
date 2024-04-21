package array

/*
73. Set Matrix Zeroes

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.



Example 1:


Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
Example 2:


Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

*/

func settingZeros1(matrix [][]int) {
	rlen := len(matrix)
	clen := len(matrix[0])
	topRow := rlen
	topCol := clen

	for row := 0; row < rlen; row++ {
		for col := 0; col < clen; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}

			if matrix[row][col] == 0 && (row == 0 || col == 0) {
				topRow = min(topRow, row)
				topCol = min(topCol, col)
			}

		}
	}
	// populate all the  rows and columns except the top row and leftmost column
	for row := 1; row < rlen; row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < clen; col++ {
				matrix[row][col] = 0
			}
		}
	}

	for col := 1; col < clen; col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < rlen; row++ {
				matrix[row][col] = 0
			}
		}
	}

	// populate the top row at the last
	if topRow == 0 {
		for col := 0; col < clen; col++ {
			matrix[0][col] = 0
		}
	}

	// populate the left most column at the last
	if topCol == 0 {
		for row := 0; row < rlen; row++ {
			matrix[row][0] = 0
		}
	}

}

func setZeroes(matrix [][]int) {
	settingZeros1(matrix)
}
