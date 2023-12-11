package main

import "fmt"

// leetcode 88
// не учитывается случай когда n<m, а также когда m = 0
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5}
	m := 3
	n := 2
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = append(nums1[:0], nums2...)
	} else {
		nums1 = append(nums1[:len(nums1)-n], nums2...)
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
	fmt.Println(nums1)
}
