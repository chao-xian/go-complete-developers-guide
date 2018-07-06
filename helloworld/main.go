package main

// only the main package is executable. Packages named
// anything else are only dependency packages.

import "fmt"

// All package main files must have a main fucntion
func main() {
	phunk()
}

func phunk() {
	fmt.Println("Hi there!")
}
