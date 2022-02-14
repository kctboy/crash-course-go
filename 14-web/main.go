package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello world")

}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server strating ....")
	http.ListenAndServe(":3000", nil)
}
