package main

import "fmt"

func main() {

	// ! TIPE DATA DI GO
	fmt.Println("Int = ", 1)
	fmt.Println("Float = ", 1.5)
	fmt.Println("Boolean = ", true)
	fmt.Println("Boolean = ", false)
	fmt.Println("String = ", "Nurdin")

	fmt.Println(len("Nurdin")) //! Menghitung banyak karakter dari string
	fmt.Println("Nurdin"[2])   // ! Mengakses index ke 2 atau karakter ke 3. akan diambil dalam tipe byte
}
