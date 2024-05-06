package main

import "fmt"

// Функция удаления элемента
func removeElement(arr []int, index int) []int {
	if len(arr) > 0 && index < len(arr) { // Обязательная проверка корректности введенных данных
		arr = append(arr[:index], arr[index+1:]...) 
	}
	return arr;
}

func main() {
	// Иниц. массива
	arr := []int{1, 2, 3, 4, 5};

	// Объявляем и присваиваем значение i
	fmt.Print("Enter i: ");
	var i int;
	_, err := fmt.Scan(&i);
	if err != nil {
		fmt.Println("Error")
		return;
	}

	fmt.Println(removeElement(arr, i));
}