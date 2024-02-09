package main

import "fmt"

func main() {
	type nim string
	type married bool

	var nimAli nim = "210111008"
	var status married = false

	fmt.Println(nimAli, "\n", status)
}
