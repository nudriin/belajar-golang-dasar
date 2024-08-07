package main

import "fmt"

func recoveries() {
	fmt.Println("End App")
	// ! RECOVER digunakan untuk menangkap panic, jadi program tetap akan berjalan
	// * dengan recover panic akan terhenti
	message := recover() // menangkap data panic bila terjadi

	fmt.Println("Terjadi error =", message)
}

func sendPanic(errors bool) {
	defer recoveries()
	if errors {
		// ! PANIC akan menghentikan program aplikasi
		// * akan tetapi defer tetap akan dipanggil
		panic("ERROR")
	}
}
func main() {
	sendPanic(true)
	fmt.Println("INI AKAN TETAP DI EKSEKUSI")
}
