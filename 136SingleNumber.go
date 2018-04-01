/*Given an array of integers, every element appears twice except for one. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?*/
import (
	"sort"
)

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := 0
	for i < len(nums)-1 {
		if nums[i] != nums[i+1] {
			return nums[i]
		} else {
			i += 2
		}
	}
	return nums[len(nums)-1]
}
