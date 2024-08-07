package main

import "fmt"

// ! Tipe interface dapat menerimaa dan mereturn data apapun
func anyReturn() interface{} {
	// return 1 // ! Bisa
	// return true // ! Bisa
	// return "string" // ! Bisa
	return "Hello" // ! Bisa
}

func main() {
	anyData := anyReturn()

	// ! ini implementasi dari interface kosong
	var data any = 1
	var data2 any = "Anything"
	var data3 any = true

	fmt.Println(anyData)
	fmt.Println(data)
	fmt.Println(data2)
	fmt.Println(data3)
}
