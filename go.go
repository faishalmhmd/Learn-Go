package main

import "fmt"


type Profile struct {
	name string
	age int
}

func PrintDataOrang(P Profile) {
	fmt.Println(P.name)
	fmt.Println(P.age)
}

func main() {
	profile_orang_1 := Profile{
		name: "Orang Dong",
		age: 10,
	}
	PrintDataOrang(profile_orang_1)

}
