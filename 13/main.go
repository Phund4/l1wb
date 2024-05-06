package main;

import (
	"fmt"
)

func main() {
	num1, num2 := 1, 2;

	// Синтаксический сахар Go
	num1, num2 = num2, num1;
	fmt.Println(num1, num2);

	a, b := 1, 2

	// XOR
	a = a ^ b
	b = a ^ b // a ^ b ^ b = a
	a = a ^ b // a ^ a ^ b = b

	fmt.Println(a, b)
}