package main

import "fmt"

func main() {
	var nama string
	var nim int

	nama = "Ali"
	fmt.Println("Nama =", nama)

	nim = 210111008
	fmt.Println("NIM =", nim)

	prodi := "Informatika"
	fmt.Println("Prodi =", prodi)

	var (
		firstName = "Ali Sya'bana"
		lastName  = "Syukurillah"
	)
	fmt.Println("Nama Lengkap =", firstName, lastName)
}
