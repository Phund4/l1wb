package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка с помощью контекста
func first(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Context method end of work")
				return
			}
		}
	}
}

// Остановка при помощи канала
func second(stopCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			{
				fmt.Println("StopCh method end of work");
				return;
			}
		}
	}
}

// Стандартный метод завершения горутины
func third(wg *sync.WaitGroup) {
	defer wg.Done();

	fmt.Println("Standart method end of work");
}

func main() {
	// Производим контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Создаём синхронизатор по указателю
	wg := new(sync.WaitGroup)

	// Создаем канал
	chInt := make(chan int)

	// Задаём счётчик горутин, будет запущено countWorker + 1 горутина на запись
	wg.Add(3)

	// Запускаем способы
	go first(ctx, wg);
	go second(chInt, wg);
	go third(wg);

	// Останавливаем время на 3 секунды
	time.Sleep(3 * time.Second);

	// Завершаем контекст
	cancel();

	// Закрываем канал
	close(chInt);

	// Ожидаем выполнения всех горутин
	wg.Wait();	
}
