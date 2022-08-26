package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom type which implements the Writer interface
// thanks for logWriter()
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Returns not exactly what we want!
	fmt.Println(resp)

	// SO....

	// Create a byteslice (with 99,999 capacity)
	// Read function won't expand the slice if it's full so make it huge
	bs := make([]byte, 99999)
	// When calling the Read function, it generally needs a byteslice
	// to return data into (essentially a pointer, but a slice is a reference type)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// Go helpers can help make this more concise though
	// io.Copy takes a type of Writer as first arg, and Reader as second
	io.Copy(os.Stdout, resp.Body)

	// Below is valid because logWriter implements the Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}