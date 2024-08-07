package main

import "fmt"

type Customer struct {
	Id    int
	Name  string
	Email string
}

// ! Membuat method struct
// func (struct)namaMethod(parama){}
func (customer Customer) methodStruct(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	// ! cara menggunakannya, buat dulu object customernya kemudian panggil functionnya
	customer := Customer{
		Name:  "Nurdin",
		Id:    1121,
		Email: "nurdin@mail.com",
	}

	// ! Memanggil methodnya akan langsung menjadi bagian dari struct nya
	customer.methodStruct("Hishasy")

}
