package main

import (
	"fmt"
	"strings"
)

// Функция проверки строки на уникальность
func isUnique(str string) bool {
	m := make(map[rune]bool); // Храним встречающиеся символы в map
	str = strings.ToLower(str); // Приводим к нижнему регистру (регистронезависимая ф-ция)
	for _, el := range str {
		if _, ok := m[el]; !ok {
			m[el] = true;
		} else {
			return false; // Символ уже есть в строке
		}
	}
	return true; // Строка уникальна
}

func main() { 
	// Иниц. переменных
	str1 := "bdnjtulg";
	str2 := "adnfGFtrS";

	fmt.Println(isUnique(str1));
	fmt.Println(isUnique(str2));
}