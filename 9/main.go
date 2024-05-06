package main

import (
	"fmt"
	"sync"
)

// Функция читает слайс и записывает элементы в канал
func ReadArrayChannel(wg *sync.WaitGroup, numberList []int, firstCh chan<- int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for index := 0; index < len(numberList); index++ {
		firstCh <- numberList[index]
	}

	close(firstCh)
}

// Функция читает канал и записывает квадраты чисел из этого канала во второй канал
func SquaringChannel(wg *sync.WaitGroup, firstCh <-chan int, secondCh chan<- int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for number := range firstCh {
		secondCh <- number * number
	}

	close(secondCh)
}

// Функция читает канал и выводит сообщения из него
func ReadChannel(wg *sync.WaitGroup, secondCh <-chan int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for number := range secondCh {
		fmt.Println(number)
	}
}

func main() {
	// Создаём синхронизатор по указателю
	wg := new(sync.WaitGroup)

	// Создаем каналы
	firstCh := make(chan int)
	secondCh := make(chan int)

	// Исходный слайс
	numberList := []int{23, 321, 314, 52, 11}

	// Задаём счётчик горутин
	wg.Add(3)

	// Запускам горутины
	go ReadArrayChannel(wg, numberList, firstCh)
	go SquaringChannel(wg, firstCh, secondCh)
	go ReadChannel(wg, secondCh)

	// Ждем завершения всех горутин
	wg.Wait()
}