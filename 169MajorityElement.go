/*Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.
*/
func majorityElement(nums []int) int {
	halfN := len(nums) / 2
	curArray := make(map[int]int)
	for _, v := range nums {
		_, ok := curArray[v]
		if !ok {
			curArray[v] = 0
		}
		curArray[v]++
		if curArray[v] > halfN {
			return v
		}

	}
	return nums[0]

}

