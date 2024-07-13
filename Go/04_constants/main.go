package main

import "fmt"

func main() {

	const name string = "Rohit"

	// not allowed
	// name = "Rohit"

	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(port)
	fmt.Println(host)

}
