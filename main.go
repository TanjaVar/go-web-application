package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is Home page")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// Add two integers and returns the sum
func addValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Start application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
