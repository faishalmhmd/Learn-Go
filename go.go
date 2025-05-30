package main

import "fmt"

func do(callback func() int) int {
	result := callback()
	return result + 10
}

func add5() int {
	return 5
}

func main() {
	fmt.Println(do(add5))
}
