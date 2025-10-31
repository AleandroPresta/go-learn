package course

import (
	"fmt"
	"strings"
)

// --- Lesson 01: Hello World ---
func Hello_Ex1() string {
	return "Hello, World!"
}

func Hello_Ex2() string {
	return "Hello, Aleandro!"
}

func Hello_Ex3() string {
	// name on one line, language on next
	return "Aleandro\nGo"
}

func Hello_Ex4() string {
	return fmt.Sprintf("My name is %s and I am learning %s", "Aleandro", "Go")
}

func Hello_Ex5() string {
	// age integer and height float demonstration
	age := 29
	height := 1.80
	return fmt.Sprintf("I am %d years old and %.2f meters high", age, height)
}

func Hello_Ex6() string {
	isLearning := true
	return fmt.Sprintf("Aleandro is learning: %t", isLearning)
}

// --- Lesson 02: Variables and Constants ---
func Vars_Ex1() string {
	name := "Aleandro"
	age := 29
	isStudent := false
	return fmt.Sprintf("name: %s age: %d isStudent: %t", name, age, isStudent)
}

func Vars_Ex2() string {
	const pi = 3.14159
	return fmt.Sprintf("pi: %.5f", pi)
}

func Vars_Ex3() string {
	firstName, lastName := "Aleandro", "Presta"
	fullName := firstName + " " + lastName
	return fullName
}

func Vars_Ex4() string {
	a := 1.8
	b := false
	c := "Hello"
	return fmt.Sprintf("types: %T %T %T", a, b, c)
}

func Vars_Ex5() string {
	var i int
	var s string
	var k bool
	return fmt.Sprintf("zero values: %v %q %t", i, s, k)
}

func Vars_Ex6() string {
	d := 5
	before := fmt.Sprintf("d: %d", d)
	d = 6
	after := fmt.Sprintf("d: %d", d)
	return before + " -> " + after
}

// --- Lesson 03: Data types and operations ---
func Data_Ex1() string {
	a, b := 10, 15
	return fmt.Sprintf("sum=%d sub=%d mul=%d div=%d mod=%d", a+b, a-b, a*b, a/b, a%b)
}

func Data_Ex2() string {
	c, d := 3.2, 4.6
	return fmt.Sprintf("sum=%f sub=%f mul=%f div=%f", c+d, c-d, c*d, c/d)
}

func Data_Ex3() string {
	s1, s2 := "Hello ", "World!"
	return s1 + s2
}

func Data_Ex4() string {
	s := "Otnematrappa"
	return fmt.Sprintf("%s len=%d", s, len(s))
}

// --- Lesson 04: Control Flow ---
func Flow_Ex1() string {
	i := 10
	if i > 0 {
		return fmt.Sprintf("%d is positive", i)
	} else if i < 0 {
		return fmt.Sprintf("%d is negative", i)
	}
	return fmt.Sprintf("%d is zero", i)
}

func Flow_Ex2() string {
	x := 11
	if x%2 == 0 {
		return "x is even"
	}
	return "x is odd"
}

func Flow_Ex3() string {
	x := 29
	switch {
	case x < 13:
		return "child"
	case x >= 13 && x <= 19:
		return "teenager"
	case x >= 20 && x <= 64:
		return "adult"
	default:
		return "senior"
	}
}

func Flow_Ex4() string {
	day := 4
	names := map[int]string{1: "monday", 2: "tuesday", 3: "wednesday", 4: "thursday", 5: "friday", 6: "saturday", 7: "sunday"}
	return names[day]
}

// --- Lesson 05: Loops ---
func Loops_Ex1() string {
	sb := strings.Builder{}
	for i := 1; i <= 10; i++ {
		sb.WriteString(fmt.Sprintf("%d ", i))
	}
	return sb.String()
}

func Loops_Ex2() string {
	sb := strings.Builder{}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("%d ", i))
	}
	return sb.String()
}

// --- Lesson 06: Functions ---
func Funcs_Ex1(a, b int) string {
	return fmt.Sprintf("%d", a+b)
}

func Funcs_Ex2(name string, age int) string {
	return fmt.Sprintf("Hello %s, you are %d", name, age)
}

func Funcs_Ex3(a, b int) string {
	sum := a + b
	prod := a * b
	return fmt.Sprintf("sum=%d prod=%d", sum, prod)
}

func Funcs_Ex4(width, height float64) string {
	area := width * height
	perimeter := 2 * (width + height)
	return fmt.Sprintf("area=%.2f perimeter=%.2f", area, perimeter)
}

func Funcs_Ex5(nums ...int) string {
	s := 0
	for _, v := range nums {
		s += v
	}
	return fmt.Sprintf("sum=%d", s)
}

// Wire up the registry for the easy lessons
func init() {
	// Lesson 1
	Register(Exercise{"01-01-01", "Hello: Hello World", true, func() string { return Hello_Ex1() }})
	Register(Exercise{"01-01-02", "Hello: Greeting with name", true, func() string { return Hello_Ex2() }})
	Register(Exercise{"01-01-03", "Hello: Multiple lines", true, func() string { return Hello_Ex3() }})
	Register(Exercise{"01-01-04", "Hello: Formatted output", true, func() string { return Hello_Ex4() }})
	Register(Exercise{"01-01-05", "Hello: Numbers", true, func() string { return Hello_Ex5() }})
	Register(Exercise{"01-01-06", "Hello: Boolean and types", true, func() string { return Hello_Ex6() }})

	// Lesson 2
	Register(Exercise{"01-02-01", "Vars: Declarations", true, func() string { return Vars_Ex1() }})
	Register(Exercise{"01-02-02", "Vars: Constants", true, func() string { return Vars_Ex2() }})
	Register(Exercise{"01-02-03", "Vars: Multiple declarations", true, func() string { return Vars_Ex3() }})
	Register(Exercise{"01-02-04", "Vars: Type inference", true, func() string { return Vars_Ex4() }})
	Register(Exercise{"01-02-05", "Vars: Zero values", true, func() string { return Vars_Ex5() }})
	Register(Exercise{"01-02-06", "Vars: Reassignment", true, func() string { return Vars_Ex6() }})

	// Lesson 3
	Register(Exercise{"01-03-01", "Data: Integer ops", true, func() string { return Data_Ex1() }})
	Register(Exercise{"01-03-02", "Data: Float ops", true, func() string { return Data_Ex2() }})
	Register(Exercise{"01-03-03", "Data: Strings", true, func() string { return Data_Ex3() }})
	Register(Exercise{"01-03-04", "Data: String length", true, func() string { return Data_Ex4() }})

	// Lesson 4
	Register(Exercise{"01-04-01", "Flow: Positive/Negative/Zero", true, func() string { return Flow_Ex1() }})
	Register(Exercise{"01-04-02", "Flow: Even/Odd with init", true, func() string { return Flow_Ex2() }})
	Register(Exercise{"01-04-03", "Flow: Age ranges", true, func() string { return Flow_Ex3() }})
	Register(Exercise{"01-04-04", "Flow: Day of week switch", true, func() string { return Flow_Ex4() }})

	// Lesson 5
	Register(Exercise{"01-05-01", "Loops: Print 1-10", true, func() string { return Loops_Ex1() }})
	Register(Exercise{"01-05-02", "Loops: Skip evens", true, func() string { return Loops_Ex2() }})

	// Lesson 6
	Register(Exercise{"01-06-01", "Functions: Sum two ints", true, func() string { return Funcs_Ex1(2, 3) }})
	Register(Exercise{"01-06-02", "Functions: Greeting", true, func() string { return Funcs_Ex2("Aleandro", 29) }})
	Register(Exercise{"01-06-03", "Functions: Sum and Product", true, func() string { return Funcs_Ex3(3, 4) }})
	Register(Exercise{"01-06-04", "Functions: Area/Perimeter", true, func() string { return Funcs_Ex4(3.0, 4.0) }})
	Register(Exercise{"01-06-05", "Functions: Variadic sum", true, func() string { return Funcs_Ex5(1, 2, 3, 4) }})

	// Intermediate and advanced lessons will be registered by their own files.
}
