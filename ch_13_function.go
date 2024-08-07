package main

import "fmt"

func main() {
	fmt.Println(getName())
	sayHello("Nurdin")

	// * wajib tangkap return valuenya
	fmt.Println()
	fmt.Println("MILTIPLES RETURN VALUE")
	name, age := getNameAndAge()
	fmt.Println(name)
	fmt.Println(age)

	fmt.Println()
	// * tidak menangkap salah satu retun valunye
	_, age2 := getNameAndAge() // Mengambil agenya saja
	fmt.Println(age2)

	name2, _ := getNameAndAge() // mengambil namenya saja
	fmt.Println(name2)

	fmt.Println()
	fmt.Println("NAMED RETURN VALUE")
	name3, age3 := namedReturnValue()
	fmt.Println(name3)
	fmt.Println(age3)

	fmt.Println()
	fmt.Println("VARIADIC FUNCTION")
	sum := variadicFunction(1, 2, 3, 4, 5, 6) //! mengisi paramter argument
	fmt.Println(sum)

	numbers := []int{10, 10, 10, 10}
	sum = variadicFunction(numbers...) // ! cara menggunakan slice msebagai inputan untuk paramter arguments
	fmt.Println(sum)

}

func getName() string {
	name := "Nurdin"
	return name
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

// ! FUNCTION DENGAN MULTIPLE RETURN VALUES
func getNameAndAge() (string, int) {
	return "Nurdin", 20
}

// ! RETURN SEBUAH VARIABLE
//  func namedReturnValue() (name, address, city string) // ! Bisa juga gini kalau tipe datanya sama semua
func namedReturnValue() (name string, age int) {
	name = "Nurdin"
	age = 20

	return name, age
}

// ! parameter yang dapat menerima data berupa argument
// mirip seperti parameter arguments di nodejs
func variadicFunction(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
