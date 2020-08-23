// https://www.codewars.com/kata/5946a0a64a2c5b596500019a/train/go

package main

import "fmt"

func help(l []int) []int {
	var result []int
	if len(l)%2 != 0 {
		result = append(result, l[int(len(l)/2)])
		for i := 0; i < int(len(l)/2); i++ {
			v := l[i] + l[len(l)/2+i+1]
			result = append(result, v)
		}
	} else {
		for i := 0; i < int(len(l)/2); i++ {
			v := l[i] + l[len(l)/2+i]
			result = append(result, v)
		}
	}

	return result
}

var (
	numbers = []int{4, 2, 5, 3, 2, 5, 7}
	n       = 2
)

func main() {
	ans := numbers
	for i := 0; i < n; i++ {
		ans = help(ans)
	}
	fmt.Println("Result is ", ans)
}
