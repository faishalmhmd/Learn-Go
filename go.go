package main

import "fmt"

// function: this function to return hello + string
// return: hello + string
func hello(name string) string {
		return "Hello " + name 
}

func main() {
	fmt.Println(hello("Azizi"))
}