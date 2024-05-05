package array

/*
54. Spiral Matrix

Given an m x n matrix, return all elements of the matrix in spiral order.



Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:


Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]



*/

func printOrder(matrix [][]int) []int {
	var ans []int
	rows := len(matrix)
	cols := len(matrix[0])
	dirc := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //  direction metrix right,down, left,up
	// it store number of step to be covered along row or col
	// after each iteration
	steps := []int{cols, rows - 1}

	iDir := 0
	r := 0
	c := -1 // so c does not out of bound
	for steps[iDir%2] > 0 {
		for i := 0; i < steps[iDir%2]; i++ {
			r += dirc[iDir][0] // update row index
			c += dirc[iDir][1] // update col index
			ans = append(ans, matrix[r][c])
		}
		steps[iDir%2]--
		iDir = (iDir + 1) % 4
	}

	return ans
}

func SpiralOrder(matrix [][]int) []int {
	return printOrder(matrix)
}
