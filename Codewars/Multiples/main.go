// https://www.codewars.com/kata/514b92a657cdc65150000006/train/go

package main

import "fmt"

func sumList(x int) int {
	var solution int
	for i := 0; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			solution += i
		}
	}
	return solution
}

func main() {
	fmt.Println(sumList(10))
}
