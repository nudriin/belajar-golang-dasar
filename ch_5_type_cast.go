package main

import "fmt"

func main() {
	var value int32 = 32768

	// ! casting dari int32 ke int64
	var value64 = int64(value)

	// ! casting dari int64 ke int16
	var value16 = int16(value) //! akan minus karena melebihisi kapasitas maksimal int16 (number overflow)

	fmt.Println(value)
	fmt.Println(value64)
	fmt.Println(value16)

	// ! casting dari byte ke string
	var name string = "Nurdin"
	var n uint8 = name[0]
	var nString = string(n) // * Casting dari byte ke string

	fmt.Println("Name =", name)
	fmt.Println("Byte =", n)
	fmt.Println("Sgtring =", nString)

}
