package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция добавления значения в map
func setValue(m *map[int]int, mut *sync.Mutex, key int, value int) {
	mut.Lock(); // Блокируем доступ к ресурсу
	defer mut.Unlock(); // Открываем доступ к ресурсу
	if _, ok := (*m)[key]; !ok {
		fmt.Printf("Goroutine add value %v on key %v\n", value, key); 
		(*m)[key] = value; // Добавляем значение по ключу
	} else {
		fmt.Printf("Map already has %v key\n", key); // В map уже есть значение по ключу
	}
}

// Функция получения значения из map
func getValue(m *map[int]int, mut *sync.Mutex, key int) {
	mut.Lock(); // Блокируем доступ к ресурсу
	defer mut.Unlock(); // Открываем доступ к ресурсу
	value, ok := (*m)[key];
	if !ok {
		fmt.Printf("Map hasn't value on %v key\n", key); // Map не имеет значения по ключу
	} else {
		fmt.Printf("Goroutine get %v\n", value) // Получаем элемент по ключу
	}
}

func main() {
	// Объявление map
	m := map[int]int{};

	// Объявление mutex
	mut := new(sync.Mutex);

	// Запуск горутин на запись в map
	for i := 0; i<10; i++ {
		go setValue(&m, mut, i, i + 100)
	}

	// Запуск горутин на чтение из map
	for i := 0; i<10; i++ {
		go getValue(&m, mut, i);
	}

	// Ожидаем завершения программы.
	time.Sleep(1 * time.Second);

}