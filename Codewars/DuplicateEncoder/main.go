package main

import "fmt"

// func stringInSlice(s string, l []string) bool {
// 	for _, value := range l {
// 		if value == s {
// 			return true
// 		}
// 	}
// 	return false
// }

func duplicateEncode(word string) string {
	// f := []string{}
	for _, val := range word {
		// res := stringInSlice(val, f)
		// if res == false {
		// 	f = append(f, val)
		// }
		// Issue is Val is a rune and if i use &val, it becomes
		// a memory location and not the exact value
		fmt.Println(val)

	}
	// fmt.Println(f)
	return word
}

func main() {
	duplicateEncode("Hi there")
}
