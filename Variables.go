package main

import (
	"fmt"
	"strconv"
)

// Global Variable
var p int = 32

// Grouping
var (
	name  string = "Aswin"
	email string = "aswinvb.aswin6@gmail.com"
	age   int    = 20
)

// Variables : Basics of variables
func Variables() {
	fmt.Println(p)
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)

	// variable declarations
	// Type 1
	var i int
	i = 10
	fmt.Println(i)

	// Type 2
	var j int = 20
	fmt.Println(j)

	//Type 3
	k := 30
	fmt.Println(k)

	// Data Types
	var n bool = false
	fmt.Printf("%v,%T\n", n, n)
	m := 1 == 2
	fmt.Printf("%v,%T\n", m, m)
	r := 1 == 1
	fmt.Printf("%v,%T\n", r, r)
	u := 2.1E14
	fmt.Printf("%v,%T\n", u, u)
	u = 13.2e12
	fmt.Printf("%v,%T\n", u, u)
	var x complex64 = 2 + 3i
	fmt.Printf("%v,%T\n", x, x)
	x = complex(3, 5)
	fmt.Printf("%v,%T\n", real(x), real(x))
	fmt.Printf("%v,%T\n", imag(x), imag(x))
	s := "hello world"
	b := []byte(s)
	fmt.Printf("%v,%T\n", b, b)
	var t rune = 'a'
	fmt.Printf("%v,%T\n", t, t)

	// Type Conversion
	// int to float
	var q float32
	q = float32(k)
	fmt.Printf("%v,%T\n", k, k)
	fmt.Printf("%v,%T\n", q, q)

	// int to string
	var w string
	w = strconv.Itoa(k)
	fmt.Printf("%v,%T\n", k, k)
	fmt.Printf("%v,%T\n", w, w)

	// Operations
	fmt.Println(k, j)
	fmt.Println(k + j)
	fmt.Println(k - j)
	fmt.Println(k * j)
	fmt.Println(k / j)
	fmt.Println(k % j)

	// Constants
	const myConst int = 32
	fmt.Printf("%v,%T\n", myConst, myConst)
}
