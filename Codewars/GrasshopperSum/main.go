package main

import "fmt"

func summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum
}

func main() {
	summation(6)
	// fmt.Println("Hello, playground")
}
