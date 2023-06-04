package main

import "fmt"

func main() {

	var decimalNumber int

	fmt.Printf("Enter the decimal number: \n")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("\nBinary Number: %b\n", decimalNumber)
	fmt.Printf("Octal Number: %o\n", decimalNumber)
	fmt.Printf("Hexa Number: %x\n", decimalNumber)

}
