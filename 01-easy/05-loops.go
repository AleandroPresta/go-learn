/*
LESSON 5: Loops - For Statement

Go uses a single looping construct: for. It supports all common loop styles and
is simple to read once you know the patterns.

Loop forms:
- Classic: for i := 0; i < n; i++ { ... }   // init; condition; post
- While-like: for condition { ... }          // when only condition needed
- Infinite: for { ... }                     // runs until a break
- Range: for idx, val := range collection     // iterate over arrays, slices, strings, maps, channels

Control flow:
- break exits the loop immediately
- continue skips to the next iteration
- Labels let you break or continue outer loops from inner loops

String iteration nuance:
- Range over a string yields rune values (Unicode code points) and byte indices.

Practice tip: Try translating a while loop you know into Go's for form and then
rewrite it with range whenever possible for readability.
*/

package main

import "fmt"

func main() {
	// Exercise 1: Traditional for loop
	// TODO: Print numbers from 1 to 10 using a traditional for loop
	
	// Exercise 2: While-like loop
	// TODO: Print numbers from 1 to 5 using for with only a condition
	
	// Exercise 3: Infinite loop with break
	// TODO: Create an infinite loop that prints numbers and breaks when it reaches 3
	
	// Exercise 4: Loop with continue
	// TODO: Print numbers 1-10 but skip even numbers using continue
	
	// Exercise 5: Nested loops
	// TODO: Create a multiplication table (1-5) using nested for loops
	
	// Exercise 6: Range over string
	// TODO: Iterate over a string and print each character with its index
	
	// Exercise 7: Range over slice/array
	// TODO: Create a slice of your favorite colors and iterate over it
	//       Print both index and value
	
	// Exercise 8: Loop patterns
	// TODO: Calculate factorial of a number using a for loop
	// TODO: Calculate the sum of all numbers from 1 to N
	
	// Exercise 9: Advanced loop control
	// TODO: Create nested loops with labels and use labeled break/continue
	//       Example: Print a pattern but skip certain combinations
	
	// Exercise 10: FizzBuzz challenge
	// TODO: Print numbers 1-50, but:
	//       - Print "Fizz" for multiples of 3
	//       - Print "Buzz" for multiples of 5
	//       - Print "FizzBuzz" for multiples of both 3 and 5
}