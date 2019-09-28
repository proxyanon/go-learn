package main

import "fmt"

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func multipleReturnVals() (string, string) {
	return "return value 1", "return value 2"
}

// Functions : Examples of functions
func Functions() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(multipleReturnVals())

	func() {
		fmt.Println("Anonymous Function")
	}()

	g := greeter{
		greeting: "hello",
		name:     "aswin",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

// Methods
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
