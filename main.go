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

	// slices (array without specify the length)
	greetings := []string{"Hello", "Hi", "ola", "sup!"}

	// for loop
	for i := 0; i < len(greetings); i++ {
		fmt.Println(greetings[i])
	}

	// add data into slices
	greetings = append(greetings, "OI")

	// for loop with range
	for _, greet := range greetings {
		fmt.Println(greet)
	}
}

// function declaration with return type
func hello() string {
	return "Hello go from function!"
}
