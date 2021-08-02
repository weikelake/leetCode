package main

import "fmt"

func main() {
	fmt.Print(twoSum([]int{3,2,4}, 6))
}


func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] + nums[j] == target && i!=j{
				return []int{j,i}
			}
		}
	}
	return []int{}
}