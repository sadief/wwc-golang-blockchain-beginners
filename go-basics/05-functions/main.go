package main

import "fmt"

func main() {
g := sayHello("Sadie", 2019)
fmt.Println(g)
}


func sayHello(name string, num int) string {
   greeting := fmt.Sprintf("Hello %v the year is %v", name, num)
   return greeting
}