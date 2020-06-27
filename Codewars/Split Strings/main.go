package main

import "fmt"

// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go

/*
"""
Complete the solution so that it splits the string into pairs of two characters.
If the string contains an odd number of characters then it should replace the missing
second character of the final pair with an underscore ('_').

Examples:

solution('abc') # should return ['ab', 'c_']
solution('abcdef') # should return ['ab', 'cd', 'ef']
"""
*/

func solution(str string) []string {
	// Create an empty list of strings
	ans := []string{}
	f := len(str)

	// Loop through the input string and check the second values
	for i, v := range str {
		if (i+1)%2 == 0 {
			ans = append(ans, (string(str[i-1]) + string(v)))
		}
	}

	// Check if it needs to have an underscore appended to it
	if f > 0 && ((f % 2) != 0) {
		ans = append(ans, (string(str[f-1]) + "_"))
	}

	return ans
}

func main() {
	fmt.Println(solution("abc"))
}

// func Solution(str string) []string {
// 	var res []string
// 	if len(str) % 2 != 0 {
// 	  str += "_"
// 	}
// 	for i := 0; i < len(str); i+=2 {
// 	  res = append(res, str[i:i+2])
// 	}
// 	return res
//   }
