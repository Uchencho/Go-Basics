package main

import (
	"fmt"
	"strings"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func points(games []string) int {
	// your code here!
	f := []int{}

	for _, i := range games {
		y := strings.Split(i, ":")
		if y[0] > y[1] {
			f = append(f, 3)
		} else if y[0] < y[1] {
			f = append(f, 0)
		} else {
			f = append(f, 1)
		}
	}
	fmt.Println(sum(f))
	return sum(f)
}

func main() {
	s := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:7", "4:2", "4:4"}
	points(s)
}

// package kata

// import ( "strings" )

// func Points(games []string) int {
//   result := 0
//   for _, game := range games {
//     str := strings.Split(game, ":")
//     x, y := str[0], str[1]
//     switch {
//       case x > y:
//         result += 3
//       case x == y:
//         result += 1
//     }
//   }
//   return result
// }
