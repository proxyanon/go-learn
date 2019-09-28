package main

import "fmt"

// Loops : Loops in go
func Loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i)
	}

	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	for {
		if k == 10 {
			break
		}
		fmt.Println(k)
		k++
	}

	s := [3]int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	ageStats := map[string]int{
		"aswin": 10,
		"bivin": 20,
	}

	for k, v := range ageStats {
		fmt.Println(k, v)
	}

	text := "hello world"

	for k, v := range text {
		fmt.Println(k, string(v))
	}

Loop:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(i, j)
			if i == 0 {
				break Loop
			}
		}
	}

}
