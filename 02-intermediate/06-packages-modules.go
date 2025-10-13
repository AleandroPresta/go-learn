/*
LESSON 12: Packages and Modules

Packages group related code; modules are higher-level units that manage
dependencies across packages. You use packages to structure code and control
visibility of names.

Package basics:
- package main defines an executable; other package names produce libraries.
- Files in the same folder should generally use the same package name.

Visibility rules:
- Capitalized names are exported (public) and visible to other packages.
- Lowercase names are unexported (private) to the package.

Imports and modules:
- Standard import: import "fmt"
- Aliased import: import io "io/ioutil"
- Blank import: import _ "some/package" (for side effects like init())
- Modules: go.mod at the repo root identifies the module path and pinned versions.

Practice tip: create a small package in a subdirectory and import it from a main
program to see how exported/unexported names behave.
*/

package main

import "fmt"

// This lesson focuses on understanding packages and modules
// Since we can't create multiple packages in a single file,
// the exercises will be more conceptual and demonstrate
// package-related concepts

func main() {
	// Exercise 1: Understanding package structure
	// TODO: Research and document the structure of a Go package
	// TODO: Explain what makes a package vs a module
	
	fmt.Println("=== EXERCISE 1: Package Structure ===")
	// Write comments explaining:
	// - What is a package?
	// - What files belong to the same package?
	// - How does the main package differ from other packages?
	
	// Exercise 2: Import paths and naming
	// TODO: Demonstrate different import styles
	fmt.Println("\n=== EXERCISE 2: Import Styles ===")
	// Show examples of:
	// - Standard imports
	// - Aliased imports
	// - Dot imports (and why they're discouraged)
	// - Blank imports
	
	// Exercise 3: Exported vs Unexported identifiers
	// TODO: Create examples showing visibility rules
	fmt.Println("\n=== EXERCISE 3: Visibility Rules ===")
	// Demonstrate:
	exportedFunction()    // This would be accessible from other packages
	unexportedFunction()  // This would NOT be accessible from other packages
	
	// Exercise 4: Package documentation
	// TODO: Show how to document packages and exported functions
	fmt.Println("\n=== EXERCISE 4: Documentation ===")
	// Write examples of:
	// - Package-level documentation
	// - Function documentation
	// - Type documentation
	
	// Exercise 5: Standard library exploration
	// TODO: Explore and use various standard library packages
	fmt.Println("\n=== EXERCISE 5: Standard Library ===")
	exploreStandardLibrary()
	
	// Exercise 6: Module concepts (theoretical)
	// TODO: Explain go.mod file and dependency management
	fmt.Println("\n=== EXERCISE 6: Modules ===")
	explainModules()
	
	// Exercise 7: Package initialization
	// TODO: Understand init() functions and package initialization order
	fmt.Println("\n=== EXERCISE 7: Package Initialization ===")
	explainInitialization()
	
	// Exercise 8: Creating a simple library structure
	// TODO: Design a package structure for a math library
	fmt.Println("\n=== EXERCISE 8: Library Design ===")
	designMathLibrary()
	
	// Exercise 9: Dependency management concepts
	// TODO: Understand semantic versioning and go.mod
	fmt.Println("\n=== EXERCISE 9: Dependencies ===")
	explainDependencyManagement()
	
	// Exercise 10: Best practices
	// TODO: List best practices for package design
	fmt.Println("\n=== EXERCISE 10: Best Practices ===")
	packageBestPractices()
}

// ExportedFunction demonstrates an exported function
func ExportedFunction() {
	fmt.Println("This function is exported (starts with capital letter)")
}

// exportedFunction demonstrates another exported function
func exportedFunction() {
	fmt.Println("This function is actually unexported (starts with lowercase)")
}

// unexportedFunction demonstrates an unexported function
func unexportedFunction() {
	fmt.Println("This function is unexported (starts with lowercase)")
}

func exploreStandardLibrary() {
	// TODO: Import and use functions from:
	// - strings package
	// - strconv package
	// - time package
	// - math package
	// - os package
	fmt.Println("TODO: Explore standard library packages")
}

func explainModules() {
	fmt.Println("TODO: Explain go.mod structure and commands:")
	fmt.Println("- go mod init")
	fmt.Println("- go mod tidy") 
	fmt.Println("- go mod download")
	fmt.Println("- Semantic versioning")
}

func explainInitialization() {
	fmt.Println("TODO: Explain:")
	fmt.Println("- init() functions")
	fmt.Println("- Package initialization order")
	fmt.Println("- Import side effects")
}

func designMathLibrary() {
	fmt.Println("TODO: Design package structure for:")
	fmt.Println("- Basic arithmetic operations")
	fmt.Println("- Geometric calculations") 
	fmt.Println("- Statistical functions")
	fmt.Println("- Proper exported/unexported split")
}

func explainDependencyManagement() {
	fmt.Println("TODO: Explain:")
	fmt.Println("- go.mod file format")
	fmt.Println("- Semantic versioning (major.minor.patch)")
	fmt.Println("- Direct vs indirect dependencies")
	fmt.Println("- Minimum version selection")
}

func packageBestPractices() {
	fmt.Println("TODO: List best practices:")
	fmt.Println("- Package naming conventions")
	fmt.Println("- Single responsibility principle")
	fmt.Println("- Minimize dependencies")
	fmt.Println("- Good API design")
	fmt.Println("- Documentation standards")
}