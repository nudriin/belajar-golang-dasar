package main //! wajib di import

import "fmt" // ! import fmt untuk package println

// ! PADA SEBUAH PROJECT MAIN FUCNTION HANYA BOLEH ADA 1 FUNC MAIN
// ! APABILA PADA FILE LAIN DI PROJECT INI ADA FUNC MAIN LAGI MAKA AKAN TERJADI ERROR
// ! ERROR TERJADI KETIKA DI COMPILE / BUILD
// ! TAPI TIDAK BERLAKU JIKA DI RUN
func main() {

	fmt.Println("Hello world", "test") // ! print menggunakan fmt
}
