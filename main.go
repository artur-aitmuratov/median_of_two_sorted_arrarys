package main

import (
	"fmt"
	"slices"
)

func main() {

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	printMedian(nums1, nums2)
}

func printMedian(nums1 []int, nums2 []int) {
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)
	if len(nums1)%2 == 0 {
		fmt.Println((nums1[(len(nums1)/2)-1] + nums1[len(nums1)/2]) / 2)
	} else {
		fmt.Println(nums1[(len(nums1)-1)/2])
	}
}
