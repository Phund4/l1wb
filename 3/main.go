package main

import (
	"fmt"
	"sync"
)

func pow(wg *sync.WaitGroup, number int, ch chan<- int) {
	// Уменьшаем счетчик, когда горутина завершается.
	defer wg.Done()

	// Записываем в канал квадрат числа, блокируя текущую горутину
	ch <- number * number
}

func main() {

	// Создаём синхронизатор по указателю, для того,
	// чтобы не запутаться и передать по указателю в функцию
	wg := new(sync.WaitGroup)

	// Создаем канал для записи квадратов чисел
	intCh := make(chan int)

	// Исходный массив
	numberList := [5]int{2, 4, 6, 8, 10}

	// Сумма квадратов чисел
	var sum int

	// Задаём счётчик горутин, будет запущено 5 горутин
	wg.Add(5)

	for i := 0; i < len(numberList); i++ {

		// Запускаем горутину
		go pow(wg, numberList[i], intCh)
	}

	// Ждем завершения всех горутин в отдельной горутине
	// Без неё основная горутина main не сможет производить чтение из канала
	go func() {
		wg.Wait()
		close(intCh) // закрытие канала
	}()

	// Цикл на чтение канала, автоматически завершится когда канал опустеет
	for number := range intCh {
		sum += number
	}

	// Вывод результата
	fmt.Println(sum)
}