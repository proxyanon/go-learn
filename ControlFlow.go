package main

import "fmt"

// ControlFlow : If and Switch
func ControlFlow() {
	if false {
		fmt.Println("if loop")
	}

	ageStats := map[string]int{
		"Adam":  12,
		"Joe":   10,
		"Joan":  22,
		"Virat": 28,
	}

	if pop, ok := ageStats["Adam1"]; ok {
		fmt.Println(pop)
	}

	number := 20
	guess := 1000

	if guess > 100 {
		fmt.Println("Guess must be less than 100")
	} else if guess < 0 {
		fmt.Println("Guess must be greater than 0")
	} else {
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

	switch test := 1 + 1; test {
	case 1, 2, 3:
		fmt.Println("Case 1 or 2 or 3")
		if true {
			break
		} else {
			fmt.Println("This will not be executed")
		}
	case 4, 5, 6:
		fmt.Println("Case 4 or 5 or 6")
	default:
		fmt.Println("Default Case")
	}

	i := 10
	switch {
	case i <= 50:
		fmt.Println("Less than 50")
		fallthrough
	case i <= 20:
		fmt.Println("Less than 20")
	default:
		fmt.Println("Between 20 and 50")
	}

	var j interface{} = "hello"

	switch j.(type) {
	case int:
		fmt.Println("type Int")
	case float64:
		fmt.Println("type Float64")
	case string:
		fmt.Println("type string")
	default:
		fmt.Println("type others")
	}
}
