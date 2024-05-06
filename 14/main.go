package main

import "fmt"

// Определение типа переменной
func definitionType(p interface{}) {
	switch p.(type) {
	case int:
		{
			fmt.Println("int")
		}
	case string:
		{
			fmt.Println("string")
		}
	case bool:
		{
			fmt.Println("boolean")
		}
	case chan bool:
		{
			fmt.Println("chan bool")
		}
	default:
		{
			fmt.Println("unknown")
		}
	}
}

func main() {
	// Тип Interface{}
	var i interface{}
	definitionType(i); // unknown

	i = 123;
	definitionType(i); // int

	i = make(chan bool);
	definitionType(i); // chan bool

	i = "str";
	definitionType(i); // string
}
