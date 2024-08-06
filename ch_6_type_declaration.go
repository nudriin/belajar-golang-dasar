package main

import "fmt"

func main() {
	//! membuat tipe data NoKTP yang bertipe string
	type NoHP string

	var phone NoHP = "081549193834"

	var phoneString string = "08213322"
	var phoneCast = NoHP(phoneString) //! casting strinng ke NoHp
	fmt.Println(phone)
	fmt.Println(phoneCast)
}
