package main

import "fmt"

func divide(weight int) bool {
	switch {
	case weight == 2:
		return false
	case weight%2 == 0:
		return true
	default:
		return false
	}
}

// func Divide(weight int) bool {
// 	return weight > 2 && weight % 2 == 0;
//   }

func main() {
	fmt.Println(divide(1))
}
