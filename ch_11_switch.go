package main

import "fmt"

func main() {
	insertMonth(12)
	switchWithShortStatement("Nurdin")
	switchWithoutCondition(2)
}

func insertMonth(month int) {
	fmt.Println("SWITCH")
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Augustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Input invalid")
	}
}

func switchWithShortStatement(name string) {
	fmt.Println()
	fmt.Println("SWITCH WITH SHORT STATEMENT")
	switch length := len(name); length > 6 {
	case false:
		fmt.Println("Nama terlalu panjang")
	case true:
		fmt.Println("Nama benar")
	}
}

func switchWithoutCondition(number int) {
	fmt.Println()
	fmt.Println("SWITCH WITHOUT CONDITION")
	switch {
	case number == 1:
		fmt.Println("Satu")
	case number == 2:
		fmt.Println("Dua")
	case number == 3:
		fmt.Println("Tiga")
	case number == 4:
		fmt.Println("Empat")
	case number == 5:
		fmt.Println("Lima")
	}
}
