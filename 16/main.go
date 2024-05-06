package main;

import (
	"fmt"
)

// Быстрая сортировка
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr;
	}

	// Выбираем опорный элемент
	pivot := arr[len(arr) - 1];

	// Инициализация массивов
	var left, right, curr []int;

	// Сортировка по массивам
	for i := 0; i<len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i]) 
		} else if arr[i] > pivot {
			right = append(right, arr[i]);
		} else {
			curr = append(curr, arr[i])
		}
	}

	// Рекурсия
	return append(append(quickSort(left), curr...), quickSort(right)...)
}

func main() {
	// Исходный массив
	arr := []int{3, 5, 2, 8, 4, 3, 7}

	fmt.Println(quickSort(arr))
}