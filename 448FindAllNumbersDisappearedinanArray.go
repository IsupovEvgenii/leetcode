/*Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]*/
package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	var newResult map[int]int
	newResult = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		newResult[nums[i]] = 0
	}
	var answer []int
	for i := 1; i <= len(nums); i++ {
		if _, ok := newResult[i]; !ok {
			answer = append(answer, i)
		}
	}
	return answer

}

func main() {
	fmt.Println(findDisappearedNumbers([]int{1, 2, 1}))
}
