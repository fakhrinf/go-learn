package main

import "fmt"

type greetings []string

func (g greetings) cetak() {
	for _, greet := range g {
		fmt.Println(greet)
	}
}