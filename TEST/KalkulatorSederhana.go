package main

import (
	"fmt"
)

func tambah(x int, y int) int {
	return x + y
}

func kurang(x int, y int) int {
	return x - y
}

func kali(x int, y int) int {
	return x * y
}

func bagi(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot bagi by zero")
	}
	return x / y, nil
}

func main() {
	fmt.Println("Select operation:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")

	var choice int
	fmt.Scan(&choice)

	var x, y int
	fmt.Print("Masukkan nomor pertama: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nomor kedua: ")
	fmt.Scan(&y)

	var result int
	var err error
	switch choice {
	case 1:
		result = tambah(x, y)
	case 2:
		result = kurang(x, y)
	case 3:
		result = kali(x, y)
	case 4:
		result, err = bagi(x, y)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("Hasil: %d\n", result)
}
