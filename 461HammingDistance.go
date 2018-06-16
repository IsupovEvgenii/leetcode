/*The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.*/
import (
	"fmt"
	"math"
	"strconv"
)

func convertToBit(x int, len int) string {
	curX := x
	result := ""
	pow2 := int(math.Pow(2, float64(len)))
	for i := len; i > 0; i-- {
		curRes := curX / pow2
		println(pow2)
		if curRes == x {
			result = "0" + result
		} else {
			result = strconv.Itoa(curRes) + result
		}

		curX = curX % pow2
		pow2 /= 2
	}
	result = result + strconv.Itoa(curX)
	return result
}

func diffString(x string, y string) int {
	result := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			result++
		}
	}
	return result
}

func hammingDistance(x int, y int) int {
	strX := convertToBit(x, 30)
	strY := convertToBit(y, 30)
	fmt.Println(strX)
	fmt.Println(strY)
	return diffString(strX, strY)
}