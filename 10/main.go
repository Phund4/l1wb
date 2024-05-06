package main;

import (
	"fmt"
)

func main() {
	// Инициализация массива и результирующей map
	temps := []float64{-25.0, -27.0, -21.0, 13.0, 19.0, 15.5, 0.7, 24.5, -5, -7, 7}
	mapTemps := make(map[int][]float64, len(temps));

	// Проходимся по элементам массива
	for _, el := range temps {

		// Ключ должен быть целым значением
		key := int(el / 10) * 10;

		// добавление значения по ключю
		mapTemps[int(key)] = append(mapTemps[int(key)], el);
	}

	fmt.Println(mapTemps);
}