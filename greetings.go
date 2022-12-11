package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type greetings []string

func (g greetings) generate(times int) greetings  {

    greetWord := []string{
        "Hello, whats up bro",
        "Hi, what a beautifull day",
        "It's a wonderful day",
        "The weather is good",
        "What a beautifull night",
        "Suki da kirenai",
        "Ohaiyo gozaimasu",
    }

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < times; i++ {
        g = append(g, fmt.Sprintf("%s +%d", greetWord[rand.Intn(len(greetWord))], i+1))
    }

    return g
}

func (g greetings) limit(total int)  greetings {
	return g[:total]
}

func (g greetings) from(index int) greetings {
	return g[index:]
}

func (g greetings) split(index int) ([]string, []string) {
	start := g[:index]
	end := g[index:]

	return start, end
}

func (g greetings) toString(glue string) string {
	return strings.Join(g, glue)
}

func (g greetings) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(g.toString("\n")), 0666)
}

func load(filename string) (greetings, error) {
	data, error := ioutil.ReadFile(filename)

	result := greetings(strings.SplitAfter(string(data), "\n"))

	return result, error
}

func (g greetings) cetak() {
    for _, greet := range g {
        fmt.Println(greet)
    }
}