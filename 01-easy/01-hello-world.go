/*
LESSON 1: Hello World and Basic Syntax

Welcome! In this first lesson you’ll write your very first Go program and learn the essential
building blocks every Go file uses. Assume no external docs—you’ll get everything you need here.

What a minimal Go program looks like:
1) Package declaration
	 - The first line declares the package. A standalone program uses package "main".
		 Any file with package main can produce an executable when built or run.

2) Imports
	 - The import block lists standard library or external packages you’ll use.
	 - For printing, we use the standard library package "fmt" (format).

3) The main function
	 - func main() { ... } is the entry point of a Go program. When you run your program,
		 the Go runtime starts here.

Printing in Go with fmt:
- fmt.Print(x): prints without adding a newline
- fmt.Println(x): prints and adds a newline at the end
- fmt.Printf(format, values...): prints using a format string with placeholders, e.g.:
	- %s for string
	- %d for integer
	- %f for float (use "%.2f" to control decimal places)
	- %t for boolean

Comment syntax:
- // starts a single-line comment

Newlines and escape sequences:
- "\n" inserts a newline; other common ones: "\t" (tab), "\"" (quote)

How to run (from your project root):
- You can run a single file directly. Example: go run 01-easy/01-hello-world.go
- Or build an executable: go build -o hello 01-easy/01-hello-world.go and then ./hello

Goal of this lesson:
- Practice printing values of different types and using formatted output.
*/

package main

import "fmt"

func main() {
	// Quick sanity print so the file runs without errors right away.
	fmt.Println("Ready to code!")

	// Exercises — do them in order to cement the basics.

	// Exercise 1: Hello, World
	// Task: Print the classic greeting.
	// Hint: Use fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	// Exercise 2: Greeting with your name
	// Task: Print "Hello, <Your Name>!".
	// Hint: Replace <Your Name> with your name as a string literal.
	fmt.Println("Hello, Aleandro!")

	// Exercise 3: Multiple lines
	// Task: Print your name on one line and your favorite programming language on the next.
	// Hint: Use two fmt.Println calls, or a single fmt.Printf with "\n".
	fmt.Println("Aleandro\nC#")

	// Exercise 4: Formatted output with Printf
	// Task: Print: "My name is %s and I am learning %s".
	// Hint: Use fmt.Printf with two %s placeholders.
	fmt.Printf("My name is %s and I am learning %s\n", "Aleandro", "Go")

	// Exercise 5: Numbers
	// Task: Print your age (integer) and height (float) on one line.
	// Hint: Use fmt.Printf with %d for age and %.2f for height to show two decimals.
	fmt.Printf("I am %d years old and %.2f meters high\n", "29", "1.8")

	// Exercise 6: Boolean and other verbs
	// Task: Define a boolean like isLearning := true and print it with %t.
	// Bonus: Try printing the type of a value using %T.
	isLearning := true
	fmt.Printf("Aleandro is learing: %t\n", isLearning)
}