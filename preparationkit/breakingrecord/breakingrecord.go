package main

import (
	"fmt"
)

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var max, min, max_count, min_count int32 = scores[0], scores[0], 0, 0
	var result []int32 = []int32{0, 0}
	for i := 0; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			max_count++
		} else if scores[i] < min {
			min = scores[i]
			min_count++
		}
	}
	result[1] = min_count
	result[0] = max_count
	return result
}

func main() {
	//scores := []int32{1,2,3,4,5,6}
	fmt.Println(breakingRecords([]int32{1, 2, 3, 4, 5, 6}))
}
