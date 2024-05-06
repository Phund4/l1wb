package main

import "fmt"

// Бинарный поиск
func binarySearch(arr []int, num int) int {
	// Индекс правоого и левого указателей
	left := 0;
	right := len(arr) - 1;

	// Поиск элемента. Сдвигаем индексы в зависимости от результата
	for left <= right {
		index := (left + right) / 2;
		if arr[index] < num {
			left = right + 1;
		} else if arr[index] > num {
			right = left - 1;
		} else {
			return index;
		}
	}

	// Элемент не найден
	return -1;
}

func main() {
	arr := []int{1, 2, 2, 5, 6, 6, 7, 9}
	fmt.Println(binarySearch(arr, 5));
}