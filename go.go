package main

import "fmt"

type Profile struct {
	name string
	age  int
}

func printProfile(p Profile) {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func main() {
	person := Profile{
		name: "Orang Dong",
		age:  10,
	}
	printProfile(person)
}
