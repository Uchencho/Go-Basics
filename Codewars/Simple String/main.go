// https://www.codewars.com/kata/5a24254fe1ce0ec2eb000078/train/go

package main

import (
	"errors"
	"fmt"
	"strings"
)

func Solution(str string, idx uint) (uint, error) {

	s := strings.Split(str, "")

	var opencount int

	for i := idx; i < uint(len(s)); i++ {
		if s[i] != "(" {
			return 0, errors.New("Not an opening bracket")
		}
		for j := i + 1; j < uint(len(s)); j++ {
			if s[j] == ")" {
				if opencount > 0 {
					opencount--
					continue
				}
				return uint(j), nil
			} else if s[j] == "(" {
				opencount++
				continue
			}
		}
	}

	return 0, errors.New("No matching bracket")
}

func main() {
	fmt.Println(Solution("((1)23(45))(aB)", 6))
}
