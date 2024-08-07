package main

import (
	"fmt"
)

func main() {
	maps()
	mapFunction()
}

func maps() {
	// * Membuat map dengan deklarasi
	person := map[string]string{
		"name":    "Nurdin",
		"age":     "20",
		"address": "Sukamara",
	}

	fmt.Println(person)
	fmt.Println("Name =", person["name"])
	fmt.Println("Age =", person["age"])
	fmt.Println("Address =", person["address"])

	fmt.Println()
	// * Membuat map kosong
	blankMap := map[string]string{}
	blankMap["id"] = "20"
	blankMap["name"] = "Pencil"
	blankMap["qty"] = "160"
	blankMap["price"] = "5000"
	fmt.Println(blankMap)
	fmt.Println("Id =", blankMap["name"])
	fmt.Println("Name =", blankMap["name"])
	fmt.Println("qty =", blankMap["qty"])
	fmt.Println("Price =", blankMap["price"])
}

func mapFunction() {
	// * Membuat map dengan function make
	fmt.Println()
	fmt.Println("Make Function")
	var maps = make(map[string]string)
	maps["name"] = "Nurdin"
	maps["age"] = "20"
	maps["gender"] = "Male"
	maps["balance"] = "9999999999999999999"
	fmt.Println(maps)

	fmt.Println()
	fmt.Println("Map Length")
	fmt.Println(len(maps))

	fmt.Println()
	fmt.Println("Delete Map")
	// * (map, key)
	delete(maps, "gender")
	fmt.Println(maps)

}
