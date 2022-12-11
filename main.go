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

	// print message using parameter
	sayWhat("say something using parameter")

	// print message into console
	fmt.Println(message)

	// slices (array without specify the length)
	_greetings := []string{"Hello", "Hi", "ola", "sup!"}

	// for loop
	for i := 0; i < len(_greetings); i++ {
		fmt.Println(_greetings[i])
	}

	// add data into slices
	_greetings = append(_greetings, "OI")

	// for loop with range
	for _, greet := range _greetings {
		fmt.Println(greet)
	}

	// create greetings from another file with custom type
	greetAgain := greetings{"Whaaaaats up brooo!"}

	// call function to print
	greetAgain.cetak()
}

// function declaration with return type
func hello() string {
	return "Hello go from function!"
}

// functio with parameter
func sayWhat(message string) {
	fmt.Println(message)
}
