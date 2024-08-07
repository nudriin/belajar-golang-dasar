package main

import "fmt"

func main() {

	slice()
	appendSlice()
	makeSlice()
	copySlice()
	arrayVsSlice()
}

func slice() {
	// * SLICE MIRIP ARRAY NAMUN INDEXNYA DAPAT BERUBAH
	// * 3 DATA PENTING PADA SLICE ADALAH POINTER, LENGTH, CAPACITY
	// * POINTER = PENTUNJUK DATA PERTAMA DI ARRAY PADA SLICE
	// * LENGTH = PANJANG SLICE
	// * CAPACITY = KAPASITAS DARI SLICENYA, DIMANA LENGTH TIDAK BOLEH LEBIH DARI CAPACITY

	names := [...]string{
		"Nurdin",
		"Hishasy",
		"Sunny",
		"Summer",
		"Sasuke",
		"Naruto",
	}

	slice1 := names[4:6] // * MEMBUAT SLICE DENGAN INDEX AWALNYA ADALAH SASUKE DAN INDEX AKHIRNYA ADALAH SEBELUM 6 = INDEX 5
	slice2 := names[:3]  // * MEMBUAT SLICE DG INDEX AWAL 0 - 2
	slice3 := names[2:]  // * MEMBUAT SLICE DG INDEX AWAL 2 - AKHIR
	slice4 := names[:]   // * MEMBUAT SLICE DG INDEX SEMUANYA

	fmt.Println("ARRAY =", names)
	fmt.Println()
	fmt.Println("SLICE 1 =", slice1)
	fmt.Println("LEN SLICE 1 =", len(slice1))
	fmt.Println("CAPACITY SLICE 1 =", cap(slice1))
	fmt.Println()

	fmt.Println("SLICE 2 =", slice2)
	fmt.Println("LEN SLICE 2 =", len(slice2))
	fmt.Println("CAPACITY SLICE 2 =", cap(slice2))
	fmt.Println()

	fmt.Println("SLICE 3 =", slice3)
	fmt.Println("LEN SLICE 3 =", len(slice3))
	fmt.Println("CAPACITY SLICE 3 =", cap(slice3))
	fmt.Println()

	fmt.Println("SLICE 4 =", slice4)
	fmt.Println("LEN SLICE 4 =", len(slice4))
	fmt.Println("CAPACITY SLICE 4 =", cap(slice4))
	fmt.Println()
}

func appendSlice() {
	fmt.Println("=============================================================")
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	fmt.Println("Days =", days)
	daySlice := days[5:]
	daySlice[0] = "Sabtu Baru"
	daySlice[1] = "Minggu Baru"
	fmt.Println("Days =", days) // * data days akan berubah, karena slice menggunakan pointer dan akan mengubah data array

	daySlice2 := append(daySlice, "Hari Libur") //* daySlice2 ini akan membuat array baru di dalam slice karena melebihi kapasitas slice pertama
	daySlice2[0] = "Hari Baru Lagi"
	fmt.Println("daySlice1 =", daySlice)
	fmt.Println("daySlice2 =", daySlice2)
	fmt.Println("Days =", days)
}

func makeSlice() {
	fmt.Println()
	fmt.Println("=============================================================")
	// * Membuat slice baru menggunakan make
	// (tipe, length, capacity)
	slice := make([]string, 2, 5)
	slice[0] = "Nurdin"
	slice[1] = "Hishasy"

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println()
	slice2 := append(slice, "Sunny")
	fmt.Println(slice)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	fmt.Println()
	slice2[0] = "Sasuke"
	fmt.Println(slice)
	fmt.Println(slice2)
}

func copySlice() {
	fmt.Println()
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("From Slice =", fromSlice)
	fmt.Println("To Slice =", toSlice)
}

func arrayVsSlice() {
	fmt.Println()
	// Ini adalah array
	arr := [...]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	// Ini adalah slice, mirip seperti array cara membuatnnnya
	slice := []int{1, 2, 3}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(slice)
}
