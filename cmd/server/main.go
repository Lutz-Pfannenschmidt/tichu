package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go WebAssembly Example")
	err := http.ListenAndServe(":1337", http.FileServer(http.Dir("../../assets")))

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}

}
