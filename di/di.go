package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet someone.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// GreetHandler respond to HTTP request.
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Amazing")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
