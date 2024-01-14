package main

import (
	"fmt"
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

func addValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)

}
