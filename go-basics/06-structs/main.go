package main

import "fmt"

// Define the struct
type Cat struct {
	Name string
	Age int
}

func main() {
	//Instantiate the struct and set fields
	var newCat = Cat{}
	fmt.Println(newCat)

	newCat.Name = "Dougal"
	fmt.Println(newCat)

	newerCat := Cat{
		Name: "Echo",
		Age: 19,
	}
	fmt.Println(newerCat)
}
