package main

import "fmt"

// ControlFlow : If For etc..
func ControlFlow() {
	if true {
		fmt.Println("if loop")
	}

	ageStats := map[string]int{
		"Adam":  12,
		"Joe":   10,
		"Joan":  22,
		"Virat": 28,
	}

	if pop, ok := ageStats["Adam"]; ok {
		fmt.Println(pop)
	}

	number := 20
	guess := 30

	if number > guess {
		fmt.Println("Too High")
	}

	if number < guess {
		fmt.Println("Too Low")
	}

	if number == guess {
		fmt.Println("Correct Guess")
	}
}
