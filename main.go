package main

import (
	"fmt"
	"slices"
)

func main() {

	num1 := []int{1, 2}
	num2 := []int{3, 4}
	getMedian(num1, num2)
}

func getMedian(num1 []int, num2 []int) {
	num1 = append(num1, num2...)
	slices.Sort(num1)
	if len(num1)%2 == 0 {
		fmt.Println((num1[(len(num1)/2)-1] + num1[len(num1)/2]) / 2)
	} else {
		fmt.Println(num1[(len(num1)-1)/2])
	}
}
