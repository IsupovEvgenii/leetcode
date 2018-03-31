func rotate(nums []int, k int)  {
	shift := k % len(nums)
	var newNums []int
	 
	newNums = append(newNums, nums[len(nums) - shift:]...)
	newNums = append(newNums, nums[0:len(nums) - shift]...)
	for i := 0; i < len(newNums); i++ {
		nums[i] = newNums[i]
	}
}
