package main

import "fmt"

// MajorityElement
// решени с помощью map оптимально, когда задача связана с уникальными элементами
func main() {
	var arr = []int{6, 7, 6, 7, 7}
	fmt.Println(Majority(arr))
}
func Majority(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	maxCount := 0
	result := 0
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}

	return result
}
