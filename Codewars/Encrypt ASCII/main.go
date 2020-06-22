package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encryptThis(text string) string {
	// Description:
	// Encrypt this!

	// You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

	// Your message is a string containing space separated words.
	// You need to encrypt each word in the message using the following rules:
	// The first letter needs to be converted to its ASCII code.
	// The second letter needs to be switched with the last letter
	// Keepin' it simple: There are no special characters in input.
	// Examples:
	// EncryptThis("Hello") == "72olle"
	// EncryptThis("good") == "103doo"
	// EncryptThis("hello world") == "104olle 119drlo"

	t := strings.Split(text, " ")

	f := ""
	for _, item := range t {
		for i, v := range item {
			if i == 0 {
				// Get the ascii character, append it to an empty string
				f += strconv.Itoa(int(v))
				continue
			} else if i == 1 {
				f += string(item[len(item)-1])
				continue
			} else if i == len(item)-1 {
				f += string(item[1])
				continue
			}
			f += string(v)
		}
		f += " "
	}
	return strings.TrimSpace(f)
}

func main() {
	fmt.Println(encryptThis("hello world"))
}
