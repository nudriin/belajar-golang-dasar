package main

import "fmt"

func main() {

	// ! membuat array, indexnya di tentukan di awal
	var arr [3]string

	arr[0] = "Nurdin"
	arr[1] = "Hishasy"
	arr[2] = "Sunny"
	// arr[3] = "Sunny" //! error out of bound

	fmt.Println(arr)

	// ! Membuat array dengan deklarasi
	var ages = [3]int{
		20,
		19,
		17,
	}

	fmt.Println(ages)

	// ! array memiliki nilai default
	// ! karena datanya tidaklah 5  maka defaultnya adalah 0 pada index 3 dan 4
	var values = [5]int{
		10,
		89,
		67,
	}

	fmt.Println(values)

	// ! mengambil jumlah data pada array menggunakan function len
	fmt.Println(len(arr))
	fmt.Println(len(ages))
	fmt.Println(len(values))

	// ! array tanpa index terbatas
	// * Hanya berlaku untuk array yang langsung di deklarasi
	var newArr = [...]string{
		"Nurdin",
		"Hishasy",
		"Sunny",
		"Summer",
	}
	// Error
	// var newArr [...]string

	fmt.Println(newArr)
	fmt.Println(len(newArr))

}
