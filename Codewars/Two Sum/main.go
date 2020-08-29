package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for id, val := range numbers {
			if id == i {
				continue
			}
			if val+numbers[i] == target {
				return [2]int{i, id}
			}
		}
	}
	return [2]int{0, 0}
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 5))
}
