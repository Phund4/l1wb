package main

import (
	"fmt"
	"sync"
)

func squareNum(wg *sync.WaitGroup, num int, results chan int) {
	// Убираем горутину из группы
	defer wg.Done()

	// Отправляем в результат квадрат числа
	results <- num * num
}

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Результирующий канал
	results := make(chan int)

	// Экземпляр wg
	wg := new(sync.WaitGroup)

	// Добавляем в группу 5 горутин
	wg.Add(5)

	// Проходимся по массиву чисел
	for _, num := range numbers {

		// Запускаем горутину
		go squareNum(wg, num, results)
	}

	// Запуск горутины ожидания завершения всех горутин и закрытие канала
	go func() {
		wg.Wait()
		close(results)
	}()

	// Просмотр канала и вывод результатов, цикл завершится когда канал станет пустым
	for res := range results {
		fmt.Println(res)
	}
}
