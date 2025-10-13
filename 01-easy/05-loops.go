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
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// Exercise 4: Loop with continue
	// TODO: Print numbers 1-10 but skip even numbers using continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", i)
	}
}
