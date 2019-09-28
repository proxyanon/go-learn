package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Defer : Defer explaination
func Defer() {
	// Defer works on LIFO basis
	defer fmt.Println("Start")
	fmt.Println("Middle")
	defer fmt.Println("End")

	// Example
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// closing of resource happens at the end
	defer res.Body.Close()
	// resource is accessible since it will be closed at the end due to defer
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// even though a is reassigned and defer is executed at the end, the value of a at that point of time without defer will be taken
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
