package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama = "Ali"
	var e = nama[1]
	var estring = string(e)

	fmt.Println(nama)
	fmt.Println(e)
	fmt.Println(estring)

}
