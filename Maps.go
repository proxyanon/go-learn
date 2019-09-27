package main

import "fmt"

// Maps : Basics of Maps
func Maps() {
	ageStats := make(map[string]int)
	ageStats = map[string]int{
		"Adam":  12,
		"Joe":   10,
		"Joan":  22,
		"Virat": 28,
	}
	fmt.Println(ageStats)
	fmt.Println(len(ageStats))
	ageStats["Adam"] = 99
	fmt.Println(ageStats)
	delete(ageStats, "Adam")
	fmt.Println(ageStats)

	// ok => true : key exist, false : key does not exist
	age, ok := ageStats["Adam"]
	fmt.Println(age, ok)
}
