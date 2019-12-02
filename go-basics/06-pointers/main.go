package main

import "fmt"

func main() {
	cats := 42

	c := &cats // Point to cats
	fmt.Println(c)
	fmt.Println(*c) // Reads cats through the pointer

	*c = 21 // Sets cats through the pointer
	fmt.Println(c)
	fmt.Println(cats) // See the new value of cats
}
