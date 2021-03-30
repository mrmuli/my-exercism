package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	var indy []int
//	indices := make([]int, len(nums))
//
//	fmt.Println(nums[0:])
//
//	for i,j := range indices {
//		fmt.Println(i, j)
//	}
//
//	return indy
//}



func main() {
	nums := []int{3,2,3}
	fmt.Println(twoSum(nums, 6))
}


func twoSum(nums []int, target int) []int {
	var indices []int
	for i, j := range nums {
		nextIndex := i + 1
		last := nums[len(nums)-1]

		if j + nums[i + 1] == target {
			indices = append(indices, i, nextIndex)
			break
		}
	}
	return indices
}

//func threeSum(nums []int, target int) []int {
//	var indices []int
//	for i := 0; i < len(nums); i++ {
//		curInd := i
//		curVal := nums[i]
//		nextIndex := i + 1
//		nextVal := nums[i + 1]
//		fmt.Println(curInd, curVal)
//		fmt.Println(nextIndex, nextVal)
//
//
//		if nums[i] + nextVal == target {
//			indices = append(indices, curInd, nextIndex)
//			continue
//		}
//	}
//	return indices
//}