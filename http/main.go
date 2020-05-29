package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://localhost:8000/") //https://google.com/
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// // Create a byte slice that store the body of the response
	// // Give it an arbitrary size of 99999
	// bs := make([]byte, 99999)

	// // Use the read function which accepts only a byte type
	// resp.Body.Read(bs)

	// // The byte type can be converted easily to string
	// fmt.Println(string(bs))

	// Create a log writer element (remember it is of type struct)
	lw := logWriter{}

	// Because lw, a struct which implements the writer Interface,
	// It can be used with the io.copy func
	io.Copy(lw, resp.Body)

	// io.Copy(os.Stdout, resp.Body)
	// io.copy takes in two arguements, first arg must implement the writer interface
	// Second arguement must implement the Reader interface
	// os.stdout has a value of type file that has a write function
	// Hence it implements the write interface
	// The resp.body implements the Reader interface

	// The io.copy function takes resp.body and inserts it to os.stdout
}

func (logWriter) Write(bs []byte) (int, error) {
	// Creating a function that has all the attributes of a Writter Interface
	// Takes in a byte slice and returns an integer and an error msg, if any
	// It is used as receiver function that can be appiled on a struct

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
