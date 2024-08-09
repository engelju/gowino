package main

import "fmt"

func main() {
	server := NewAPIServer(":8080")

	err := server.Start()
	if err != nil {
		println("Error starting server: ", err)
	}

	fmt.Println("Test")
}
