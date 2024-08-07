package main

import "fmt"

func funcAsValue(name string) string {
	return "Hello " + name
}

func asParameter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// ! Memiliki funnction sebagain parameter
func useParameter(name string, filter func(string) string) string {
	return "Hello " + filter(name)
}

func main() {
	// Menginisiasikan function ke variabel
	fmt.Println()
	fmt.Println("FUNCTION AS VALUE")
	sayHello := funcAsValue
	fmt.Println(sayHello("Nurdin"))

	fmt.Println()
	fmt.Println("FUNCTION AS PARAMETER")
	fmt.Println(useParameter("Anjing", asParameter)) // ! Memanggil function paramater

	fmt.Println()
	fmt.Println("FUNCTION AS VALUE AND PARAMETER")
	filter := asParameter
	fmt.Println(useParameter("Anjing", filter)) // ! Memanggil function paramater
}
