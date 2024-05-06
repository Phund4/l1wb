package main

import (
	"fmt"
	"strings"
)

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

// Функция переворота отдельных слов в строке
func reverseWords(str string) string {
	// Разбиваем строку на слова
	arr := strings.Split(str, " ");

	// используем ранее написанную функцию к каждому слову
	for i := 0; i<len(arr); i++ {
		arr[i] = reverseString(arr[i]);
	}

	// Соединяем массив
	return strings.Join(arr, " ");
}

func main() {
	// Изначальный массив
	str := "abcd ghtl gr ty";
	
	fmt.Println(reverseWords(str));
}