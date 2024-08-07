package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil logging")
}

func runApp() {
	// ! DEFER akkan selalu di eksekusi ketika sebuah function selesai dijalankan
	// * Logging ini akan di ekseskusi ketika function runApp di eksekusi
	// * meskipun runApp error defer tetap akan memanggil logging
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApp()
}
