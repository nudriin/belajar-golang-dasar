package main

import "fmt"

type Blacklist func(string) bool

func filterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklist")
	} else {
		fmt.Println("Hello", name)
	}
}

func main() {
	// ! Membuat anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	filterUser("Nurdin", blacklist)

	// ! Menginisiasikan langsung anonymous function ke dalam function filterUser
	filterUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
