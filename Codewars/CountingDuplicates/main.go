package main

import (
	"fmt"
	"strings"
)

func checkString(sl []string, val string) bool {
	for _, s := range sl {
		if s == val {
			return true
		}
	}
	return false
}

func main() {

	v := "abcdeabcdeaabcdeabllindivisibility"
	v = strings.ToLower(v)
	var count int
	setSlice := []string{}
	resultSlice := []string{}
	for _, val := range v {
		result := checkString(setSlice, string(val))
		chec := checkString(resultSlice, string(val))
		if !result {
			setSlice = append(setSlice, string(val))
		} else if result && !chec {
			count++ // append i to a result list
			resultSlice = append(resultSlice, string(val))
		}
	}
	fmt.Println(count, len(setSlice), setSlice, resultSlice)
}
