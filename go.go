package main

import "fmt"



func main() {

	// this function to call anonymous function
	// return: hello + name
	hello := func(name string)string {
		return "Hello " + name
	}
	
	fmt.Println(hello("Azizi"))
}