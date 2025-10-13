/*
LESSON 3: Basic Data Types and Operations

This lesson reviews the common built-in types and the operations you can use
with them. You should be able to perform arithmetic, compare values, and
manipulate strings after this lesson.

Type highlights:
- Integers: signed and unsigned variants. "int" size depends on the platform (32 or 64-bit).
- Floating point: use float64 for most cases for precision.
- Complex numbers: available as complex64/complex128 for specialized math.
- Strings: immutable UTF-8 sequences. Use + to concatenate. Use backticks for raw strings to avoid escape processing.
- Booleans: results of comparisons and logical operators.

Operation notes & gotchas:
- Integer division truncates toward zero (5 / 2 == 2). Use floats to get fractional results.
- Mixing types requires conversion: float64(i) or int(f).
- Use fmt.Sprintf to convert values to strings in formatted ways.

Practice tip: print each result with its type using fmt.Printf("%v (%T)\n", value, value)
so you can see how expressions change types.
*/

package main

import "fmt"

func main() {
	// Exercise 1: Integer operations
	// TODO: Declare two integers and perform all arithmetic operations (+, -, *, /, %)
	// TODO: Print the results of each operation
	var a, b int = 10, 15
	fmt.Printf("sum: %d\nsub: %d\nmul: %d\ndiv: %d\nmod: %d\n", a+b, a-b, a*b, a/b, a%b)

	// Exercise 2: Floating-point operations
	// TODO: Declare two float64 numbers and perform arithmetic operations
	// TODO: Notice the difference in division between integers and floats
	var c, d float64 = 3.2, 4.6
	fmt.Printf("sum: %f\nsub: %f\nmul: %f\ndiv: %f\n", c+d, c-d, c*d, c/d)
	// Exercise 3: String operations
	// TODO: Declare two strings and concatenate them using +
	var s1, s2 string = "Hello ", "World!"
	var s3 = s1 + s2
	fmt.Printf("%s %T\n", s3, s3)
	// TODO: Create a multi-line string using backticks
	var s4 string = `
		this is a line
		this is another line
	`
	fmt.Printf("%s\n", s4)
	// TODO: Print the length of a string using len() function
	var s5 string = "Otnematrappa"
	fmt.Printf("s: %s\nlen(%s): %d\n", s5, s5, len(s5))

	// Exercise 4: Boolean operations
	// TODO: Declare two boolean variables
	var b1, b2 bool = true, false
	// TODO: Test all logical operations (&&, ||, !) and print results
	fmt.Printf("%t && %t = %t\n", b1, b2, b1 && b2)
	fmt.Printf("%t || %t = %t\n", b1, b2, b1 || b2)
	fmt.Printf("!%t: %t\n", b1, !b1)

	// Exercise 6: Type conversion
	// TODO: Convert a float to integer (notice what happens to the decimal part)
	var f float64 = 10.6245
	var i int = int(f)
	fmt.Printf("f: %f %T\ni: %d %T\n", f, f, i, i)
	// TODO: Convert an integer to float64
	var f2 float64 = float64(i)
	fmt.Printf("f2: %f %T\n", f2, f2)
	// TODO: Convert a number to string using fmt.Sprintf
	var s6 string = fmt.Sprintf("%d", i)
	fmt.Println(s6)
}
