package main

import (
	"fmt"

	"github.com/AleandroPresta/go-learn/course"
)

func main() {
	// Print a short header and the course report from the registry
	fmt.Println("Go-Learn Course Runner")
	fmt.Println("=====================")
	fmt.Print(course.Report())
}
