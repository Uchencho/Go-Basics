// https://www.codewars.com/kata/52774a314c2333f0a7000688/train/go

// STILL WORKING ON THIS

package main

import (
	"fmt"
	"strings"
)

func checkParenthesis(parens string) bool {
	switch {
	case len(parens)%2 != 0:
		return false
	case parens == "":
		return true
	default:
		l := strings.Split(parens, "")

		for i, v := range l {
			// If the first value starts with a closing parenthesis
			if i == 0 && v == ")" {
				return false
			}
			// If the next value is as it shud
			// if v == "("
		}
		return true
	}
}

func main() {
	// fmt.Println(checkParenthesis(")"))
	// fmt.Println(checkParenthesis("()()()())"))
	fmt.Println(checkParenthesis("(()()()())(())"))
}

// Get the middle position, for 14, that is 7
