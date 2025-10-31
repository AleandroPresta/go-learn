package course

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// --- Lesson 02 (Intermediate): Arrays & Slices ---
func Arrays_Ex1() string {
	arr := [5]int{1, 2, 3, 4, 5}
	return fmt.Sprintf("array: %v len=%d", arr, len(arr))
}

func Arrays_Ex2() string {
	s := make([]int, 0, 2)
	caps := []int{cap(s)}
	for i := 1; i <= 8; i++ {
		s = append(s, i)
		caps = append(caps, cap(s))
	}
	return fmt.Sprintf("slice final: %v cap_growth=%v", s, caps)
}

func Arrays_Ex3() string {
	s := []int{3, 7, 2, 9, 4}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return fmt.Sprintf("slice: %v max=%d", s, max)
}

// --- Lesson 03 (Intermediate): Maps ---
func Maps_Ex1() string {
	words := []string{"go", "is", "fun", "go", "concurrency", "is", "fun"}
	freq := map[string]int{}
	for _, w := range words {
		freq[w]++
	}
	// produce stable-ish output
	var parts []string
	for k, v := range freq {
		parts = append(parts, fmt.Sprintf("%s:%d", k, v))
	}
	return fmt.Sprintf("words=%v counts=%s", words, strings.Join(parts, ", "))
}

func Maps_Ex2() string {
	m := map[string]int{"a": 1, "b": 2}
	if v, ok := m["c"]; ok {
		return fmt.Sprintf("found c=%d", v)
	}
	return fmt.Sprintf("c not found, default=%d", -1)
}

func Maps_Ex3() string {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 3, "z": 4}
	for k, v := range b {
		a[k] = v
	}
	return fmt.Sprintf("merged=%v", a)
}

// --- Lesson 04 (Intermediate): Structs & Methods ---
type Person struct {
	Name  string
	Age   int
	email string // unexported field
}

func (p Person) Greet() string { return fmt.Sprintf("Hi, I'm %s", p.Name) }
func (p *Person) Birthday()    { p.Age++ }

func Structs_Ex1() string {
	p := Person{Name: "Ada", Age: 30, email: "ada@example.com"}
	before := p.Greet()
	p.Birthday()
	after := fmt.Sprintf("%s is now %d", p.Name, p.Age)
	return fmt.Sprintf("%s -> %s", before, after)
}

// Rectangle example with methods and pointer vs value receivers
type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64      { return r.W * r.H }
func (r Rectangle) Perimeter() float64 { return 2 * (r.W + r.H) }
func (r *Rectangle) Scale(s float64)   { r.W *= s; r.H *= s }

func Structs_Ex2() string {
	rect := Rectangle{W: 3, H: 4}
	area := rect.Area()
	per := rect.Perimeter()
	rect.Scale(2)
	return fmt.Sprintf("area=%.2f per=%.2f scaled_area=%.2f", area, per, rect.Area())
}

// --- Lesson 05 (Intermediate): Interfaces ---
type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

func Interfaces_Ex1() string {
	var s Shape = Circle{R: 2}
	return fmt.Sprintf("shape area=%.2f", s.Area())
}

func Interfaces_Ex2() string {
	shapes := []Shape{Rectangle{W: 2, H: 3}, Circle{R: 1}}
	var sum float64
	for _, sh := range shapes {
		sum += sh.Area()
	}
	return fmt.Sprintf("shapes_total_area=%.2f", sum)
}

// --- Lesson 06 (Intermediate): Error Handling ---
var ErrDivideByZero = errors.New("divide by zero")

func SafeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func Error_Ex1() string {
	if v, err := SafeDiv(10, 0); err != nil {
		return fmt.Sprintf("error: %v", err)
	} else {
		return fmt.Sprintf("result=%.2f", v)
	}
}

type ValidationError struct{ Field, Problem string }

func (v ValidationError) Error() string { return fmt.Sprintf("%s: %s", v.Field, v.Problem) }

func Error_Ex2() string {
	err := ValidationError{Field: "age", Problem: "must be >= 0"}
	return fmt.Sprintf("validation error example: %v", err)
}

// --- Lesson 07 (Intermediate): Packages & Modules (conceptual with demo) ---
func Packages_Ex1() string {
	// Demonstrate exported vs unexported behavior via explained string
	return "Packages: exported names start with capital letters. Use small packages with clear APIs. See COPILOT_INSTRUCTIONS.md for workflow."
}

// Register intermediate exercises
func init() {
	Register(Exercise{"02-01-01", "Arrays: basic array", true, func() string { return Arrays_Ex1() }})
	Register(Exercise{"02-01-02", "Arrays: slice capacity growth", true, func() string { return Arrays_Ex2() }})
	Register(Exercise{"02-01-03", "Arrays: find max", true, func() string { return Arrays_Ex3() }})

	Register(Exercise{"02-02-01", "Maps: word frequency", true, func() string { return Maps_Ex1() }})
	Register(Exercise{"02-02-02", "Maps: safe access (comma ok)", true, func() string { return Maps_Ex2() }})
	Register(Exercise{"02-02-03", "Maps: merge maps", true, func() string { return Maps_Ex3() }})

	Register(Exercise{"02-03-01", "Structs: Person and methods", true, func() string { return Structs_Ex1() }})
	Register(Exercise{"02-03-02", "Structs: Rectangle methods and scale", true, func() string { return Structs_Ex2() }})

	Register(Exercise{"02-04-01", "Interfaces: Circle implements Shape", true, func() string { return Interfaces_Ex1() }})
	Register(Exercise{"02-04-02", "Interfaces: Sum areas", true, func() string { return Interfaces_Ex2() }})

	Register(Exercise{"02-05-01", "Errors: Safe division", true, func() string { return Error_Ex1() }})
	Register(Exercise{"02-05-02", "Errors: Validation example", true, func() string { return Error_Ex2() }})

	Register(Exercise{"02-06-01", "Packages: concepts and guidance", true, func() string { return Packages_Ex1() }})
}
