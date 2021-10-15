package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting...")

	http.Handle("/", http.FileServer(http.Dir("../../web")))
	http.ListenAndServe(":8080", nil)

}
