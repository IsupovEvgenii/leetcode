/*Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.*/
func moveZeroes(nums []int) {
	i := 0
	numberNonZero := 0
	for i < len(nums) {
		if nums[i] != 0 {
			cur := nums[i]
			nums[i] = nums[numberNonZero]
			nums[numberNonZero] = cur
			numberNonZero++
		}
		i++

	}
}

