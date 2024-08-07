package main

import "fmt"

// ! membuat type alias unntuk function
type Filter func(string) string

func filter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// ! Menggunakan type declaration sebagai alias unntuk filter
func sayHelloFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func main() {
	sayHello := sayHelloFilter

	sayHello("Anjing", filter)
	sayHello("Sasuke", filter)
}
