package main

import (
	"fmt"
)

func main() {
	fmt.Print("Input quantity of numbers: ")
	var n int
	if m, err := fmt.Scan(&n); m != 1 {
		panic(err)
	}
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		if _, err := fmt.Scan(&nums[i]); err != nil {
			panic(err)
		}
	}
	fmt.Println("Input numbers:")
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[j-1] {
				t := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = t
			}
		}
	}
	fmt.Println(nums)
}
