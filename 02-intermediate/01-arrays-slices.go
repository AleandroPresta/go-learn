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

func main() {
	// Exercise 1: Array basics
	// TODO: Declare an array of 5 integers and initialize it with values 1,2,3,4,5
	// TODO: Print the array and its length
	// TODO: Access and modify individual elements
	
	// Exercise 2: Array iteration
	// TODO: Iterate through the array using traditional for loop
	// TODO: Iterate using range and print index and value
	
	// Exercise 3: Slice creation and basic operations
	// TODO: Create a slice using make() with length 3 and capacity 5
	// TODO: Create a slice using slice literal: []int{1, 2, 3, 4, 5}
	// TODO: Print length and capacity of both slices
	
	// Exercise 4: Slice appending
	// TODO: Start with an empty slice and append elements one by one
	// TODO: Append multiple elements at once
	// TODO: Observe how capacity changes as you append
	
	// Exercise 5: Slice slicing
	// TODO: Create a slice with values [0,1,2,3,4,5,6,7,8,9]
	// TODO: Create sub-slices: first 3 elements, last 3 elements, middle elements
	// TODO: Create a slice that skips every other element
	
	// Exercise 6: Slice copying
	// TODO: Create two slices and copy elements from one to another
	// TODO: Modify the original slice and see if the copy is affected
	
	// Exercise 7: Multi-dimensional arrays/slices
	// TODO: Create a 2D array (matrix) and initialize it with values
	// TODO: Create a slice of slices and populate it
	
	// Exercise 8: Slice gotchas
	// TODO: Create a slice, take a sub-slice, modify the sub-slice
	// TODO: Check if the original slice is affected (understand shared underlying array)
	
	// Exercise 9: Dynamic slice operations
	// TODO: Implement a function that removes an element at a specific index
	// TODO: Implement a function that inserts an element at a specific index
	
	// Exercise 10: Practical exercise
	// TODO: Create a function that finds the maximum value in a slice
	// TODO: Create a function that reverses a slice in place
	// TODO: Create a function that removes duplicates from a slice
}