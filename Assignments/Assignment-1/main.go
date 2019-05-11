package main

import "fmt"

func main() {
	// initializing the slice
	nums := []int{}

	// generating the slice
	for i := 0; i < 11; i++ {
		nums = append(nums, i)
	}

	// printing the desired output based
	// on the slice value
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
