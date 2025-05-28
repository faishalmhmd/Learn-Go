package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka := "11"

	nilai, err := strconv.Atoi(angka)
	if err != nil {
		fmt.Println("Gagal konversi:", err)
		return
	}

	fmt.Println("Hasil konversi:", nilai)
}
