package main

import "fmt"

func main() {

	// function: this function to return hello + string
	// return: hello + string
	hello := func(name string) string {
		return "hello" + name
	}
	fmt.Println(hello("Azizi"))
}