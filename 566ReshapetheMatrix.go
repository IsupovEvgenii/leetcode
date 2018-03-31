/*In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.*/
func matrixReshape(nums [][]int, r int, c int) [][]int {
	var curArray []int
	for i := 0; i < len(nums); i++ {
		curArray = append(curArray, nums[i]...)
	}
	var newMatrix [][]int
	if len(curArray) == r*c {

		if r == 1 {
			newMatrix = append(newMatrix, curArray)
			return newMatrix
		}
		if c == 1 {
			for i := 0; i < len(curArray); i++ {
				newMatrix = append(newMatrix, []int{curArray[i]})

			}
			return newMatrix
		}

		for i := 0; i < r; i++ {
			var row []int
			for j := i * c; j < (i+1)*c; j++ {
				row = append(row, curArray[j])
			}
			newMatrix = append(newMatrix, row)
			row = []int{}

		}
		return newMatrix
	} else {
		return nums
	}
}
