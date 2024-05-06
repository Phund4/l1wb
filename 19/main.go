package main

import "fmt"

// Функция переворота строки
func reverseString(str string) string {
	// Создаем слайс рун
	res := []rune(str)

	// Переворачиваем строку
	for i := 0; i < len(str)/2; i++ {
		res[i], res[len(str)-i-1] = res[len(str)-i-1], res[i]
	}

	// Приводим массив к типу string
	return string(res)
}

func main() {
	// Изначальный массив
	str := "abcdefg"
	
	fmt.Println(reverseString(str))
}
