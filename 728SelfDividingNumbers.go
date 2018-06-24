/*A self-dividing number is a number that is divisible by every digit it contains.

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

Also, a self-dividing number is not allowed to contain the digit zero.

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

Example 1:
Input:
left = 1, right = 22
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]*/
import (
	"fmt"
	"strconv"
)

func dividers(x int) []int {
	stringX := strconv.Itoa(x)
	curX := x
	var result []int

	for i := 0; i < len(stringX); i++ {
		if curX%10 == 0 {
			return []int{}
		}
		result = append(result, curX%10)
		curX = curX / 10

	}

	return result
}

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		curDividers := dividers(i)
		cur1 := 0
		fmt.Println(curDividers)
		if len(curDividers) != 0 {
			for _, j := range curDividers {
				if i%j != 0 {
					break
				}
				cur1++
			}
			if cur1 == len(curDividers) {
				result = append(result, i)
			}
		}
		cur1 = 0
	}
	return result
}
