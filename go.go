package main

import "fmt"

func main() {
	bintang := 10

	for i := 1; i<= bintang; i++ {
		for j :=1; j <= i; j++ {
			fmt.Print("X")
		}
		fmt.Println()
	}
}