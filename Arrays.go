package main

import "fmt"

// Arrays : Basics Of Arrays
func Arrays() {
	// Integer Array
	grades := [...]int{9, 12, 33}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Number of Elements : %v\n", len(grades))

	// String Array
	var students [3]string
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)

	// 2-D Matrix
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Printf("Identity Matrix : %v\n", identityMatrix)

	//Slice
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", len(a))
	fmt.Printf("%v\n", cap(a))
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[5] = 100
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Make function to create array
	f := make([]int, 0, 100)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	//Append element to an array
	f = append(f, 10)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	//Append elemets to an array
	f = append(f, 11, 12, 13, 14)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	// Concatenate Two Arrays
	g := []int{2, 3, 4}
	f = append(f, g...)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	// Stack Operations
	// Push
	g = append(g, 5)
	fmt.Println(g)
	// Pop
	h := g[1:]
	fmt.Println(h)

	// Remove element from an index
	i := append(g[:2], g[3:]...)
	fmt.Println(i)
}
