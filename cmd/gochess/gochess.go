package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("[GLOBAL] Starting Gochess...")

	fmt.Println("[BACKEND] Starting backend API...")

	fmt.Println("[BACKEND] Started backend API...")

	fmt.Println("[FRONTEND] Starting frontend server...")

	http.Handle("/", http.FileServer(http.Dir("../../web")))
	http.ListenAndServe(":80", nil)

	fmt.Println("[FRONTEND] Started frontend server...")

	fmt.Println("[GLOBAL] Started Gochess...")
}
