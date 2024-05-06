package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Writer воркер писателя данных в канал
func Writer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Writer completed the work")
			return
		default:
			ch <- rand.Intn(9999)              // записываем рандомные числа в канал
			time.Sleep(time.Millisecond * 400) // Ожидаем 400 миллисекунд
		}
	}
}

// Reader воркер читателя данных из канала
func Reader(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Reader completed the work")
			return
		default:
			fmt.Println("Worker read:", <-ch)
		}
	}
}

func main() {
	// Производим контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Создаём синхронизатор по указателю
	wg := new(sync.WaitGroup)

	// Создаем канал
	intCh := make(chan int)

	// Задаём счётчик горутин, будет запущено countWorker + 1 горутина на запись
	wg.Add(1)

	// Запускаем горутину писателя
	go Writer(ctx, intCh, wg)

	// Запускаем горутины читателей
	go Reader(ctx, intCh, wg)

	// Программа будет выполняться 10 секунд
	n := 10;
	time.Sleep(time.Second * time.Duration(n));

	// Уведомляем контексты об их завершении
	cancel()

	// Закрываем канал
	close(intCh)

	// Ждем завершения всех горутин
	wg.Wait()

}