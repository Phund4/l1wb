package main;

import (
	"fmt"
	"strconv"
)

func main() {
	// Иниц. числа
	var num int64 = 520;
	
	// Объявление и присваивание значения i
	fmt.Print("Enter index: ")
	var i int;
	fmt.Scan(&i);

	// Валидация индекса
	if i < 0 {
		fmt.Println("Index must be positive!")
		return;
	}

	// Преобразование числа в двоичный вид
	bin := strconv.FormatInt(num, 2);

	// Индекс бита в строке
	index := (len(bin) - i - 1);

	// Создание нового бита
	var resBit string;
	if bin[index] == '1' {
		resBit = "0";
	} else {
		resBit = "1";
	}

	// Создание новой строки
	bin = bin[:index] + resBit + bin[index+1:];

	fmt.Println(bin);

	// Конвертация бинарного представления числа в десятичную форму
	res, _ := strconv.ParseInt(bin, 2, 11)

	fmt.Println(res);
}