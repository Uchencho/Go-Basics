package main

import (
	"fmt"
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

	/*
		I haven't solved this yet
	*/

	for i, v := range text {
		fmt.Println(i, v)
	}
	// byteArray := []byte(text)

	// fmt.Println(byteArray, string(byteArray))
	return string(text)
}

func main() {
	fmt.Println(encryptThis("hello world"))
}
