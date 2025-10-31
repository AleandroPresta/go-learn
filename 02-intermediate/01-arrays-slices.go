/*
LESSON 7: Arrays and Slices

Arrays are fixed-size; slices are dynamic, flexible views over arrays. Slices are the
primary collection type you will use in Go.

Arrays:
- Declaration: var a [3]int
- Size is part of the type: [3]int != [4]int

Slices:
- Declaration: var s []int or s := []int{1,2,3}
- Make with make([]T, len, cap) to control length and capacity
- Append with append(s, v) which may allocate a new underlying array

Important behavior:
- Slicing shares the underlying array; modifying a subslice may modify the original
- Use copy to create an independent copy of slice content
- Capacity grows typically by doubling when append exceeds current capacity

Practice tip: experiment with len and cap while appending to understand growth behavior.
*/

package main

import "fmt"

func findMax(a []int) int {
	iMax := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[iMax] {
			iMax = i
		}
	}
	return a[iMax]
}

func main() {
	fmt.Println("============================================================")
	// Exercise 1: Array basics
	// TODO: Declare an array of 5 integers and initialize it with values 1,2,3,4,5
	// var a [5]int = [5]int{1,2,3,4,5}
	// TODO: Print the array and its length
	// fmt.Printf("a: %v\nlen(a): %d\n ", a, len(a))
	// TODO: Access and modify individual elements
	// a[2] = 10
	// fmt.Printf("a: %v\nlen(a): %d\n ", a, len(a))
	// Exercise 2: Array iteration
	// TODO: Iterate through the array using traditional for loop
	// for i:= 0; i<len(a); i++ {
	// fmt.Println(a[i])
	// }

	// Exercise 3: Slice creation and basic operations
	// TODO: Create a slice using make() with length 3 and capacity 5
	// s := make([]int, 3, 5)
	// fmt.Println(s)
	// TODO: Create a slice using slice literal: []int{1, 2, 3, 4, 5}
	// TODO: Print length and capacity of both slices
	// a1 := []int{1,2,3,4,5}
	// fmt.Printf("a1: %v\nlen(a1): %d\ncap(a1): %d\n", a1, len(a1), cap(a1))
	// Exercise 4: Slice appending
	// TODO: Start with an empty slice and append elements one by one
	// a2 := []int{}
	// a2 = append(a2, 1)
	// a2 = append(a2, 4)
	// fmt.Println(a2)
	// fmt.Printf("cap(a2): %d\n", cap(a2))
	// TODO: Append multiple elements at once
	// a2 = append(a2, 5, 6)
	// fmt.Println(a2)
	// TODO: Observe how capacity changes as you append
	// fmt.Printf("cap(a2): %d\n", cap(a2))

	// Exercise 5: Slice slicing TODO: Create a slice with values [0,1,2,3,4,5,6,7,8,9]
	// a3 := [...]int{0,1,2,3,4,5,6,7,8,9}
	// fmt.Printf("a3: %v\n", a3)
	// TODO: Create sub-slices: first 3 elements, last 3 elements, middle elements
	// var slice1 []int = a3[0:3]
	// fmt.Printf("slice1: %v\n", slice1)
	// var slice2 []int = a3[7:10]
	// fmt.Printf("slice2: %v\n", slice2)
	// TODO: Modify the original slice and see if the copy is affected
	// var a4 = [...]int{1,2,3,4}
	// fmt.Printf("array before modifying slice: %v\n", a4)
	// var slice3 []int = a4[2:]
	// slice3[0] = 999
	// fmt.Printf("array after modifying slice: %v\n", a4)

	// Exercise 7: Multi-dimensional arrays/slices
	// TODO: Create a 2D array (matrix) and initialize it with values
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("m1: %v\n", m1)

	fmt.Println("============================================================")
	// TODO: Create a slice of slices and populate it
	sm1 := m1[1:2]
	fmt.Printf("sm1: %v\n", sm1)
	// Exercise 8: Slice gotchas
	// TODO: Create a slice, take a sub-slice, modify the sub-slice
	// TODO: Check if the original slice is affected (understand shared underlying array)
	sm1[0][1] = 1
	fmt.Printf("m1: %v\n", m1)

	fmt.Println("============================================================")

	// Exercise 9: Practical exercise
	// TODO: Create a function that finds the maximum value in a slice

	fmt.Println(findMax(m1[0]))
}
