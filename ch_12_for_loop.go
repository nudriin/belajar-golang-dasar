package main

import "fmt"

func main() {
	forLoop()
	forRange()
}

func forLoop() {
	i := 1

	fmt.Println("FOR LOOP")
	for i <= 10 {
		fmt.Println("Lopping -", i)
		i++
	}

	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Println("Lopping -", i)
	}
}

func forRange() {

	fmt.Println()
	fmt.Println("WITHOUT FOR RANGE")
	names := []string{
		"Nurdin",
		"Hishasy",
		"Sunny",
		"Summer",
		"Naruto",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println("Name =", names[i])
	}

	// * For range disini mirip seperti foreach
	fmt.Println()
	fmt.Println("WITH FOR RANGE")
	// key, value, variables
	for index, name := range names {
		println(index, "=", name)
	}

	// * Bisa seperti ini jiak tidak ingin mengambil indexnya
	fmt.Println()
	for _, name := range names {
		fmt.Println("Name :", name)
	}
}
