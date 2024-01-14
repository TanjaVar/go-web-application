package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
}

func divideValues(x, y int) (float64, error) {
	if y <= x {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := y / x
	return float64(result), nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Start application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
