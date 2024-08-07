package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func panicApp(errors bool) {
	defer endApp()
	if errors {
		// ! PANIC akan menghentikan program aplikasi
		// * akan tetapi defer tetap akan dipanggil
		panic("ERROR")
	}
}
func main() {
	panicApp(true)
}
