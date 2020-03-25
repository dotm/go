package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Greet sends a personalised greeting to writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler says Hello, world over HTTP
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "netizen")
}

func main() {
	//writing (printing) to standard output
	Greet(os.Stdout, "user")

	//writing to a http response
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
