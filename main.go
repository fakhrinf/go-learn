package main

import "fmt"

// variable declaration
// var message string

func main() {

	// normal variable declaration
	// var message string = "hello go in gitpod!"

	// abv variable declaration
	message := "hello go in gitpod"

	// replace variable value
	message = "second time hello go in gitpod"

	// replace variable value by function
	message = hello()

	fmt.Println(message)
}

// function declaration with return type
func hello() string {
	return "Hello go from function!"
}
