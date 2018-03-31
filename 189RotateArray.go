/*Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].*/
func rotate(nums []int, k int)  {
	shift := k % len(nums)
	var newNums []int
	 
	newNums = append(newNums, nums[len(nums) - shift:]...)
	newNums = append(newNums, nums[0:len(nums) - shift]...)
	for i := 0; i < len(newNums); i++ {
		nums[i] = newNums[i]
	}
}
