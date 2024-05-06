package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) // After ожидает истечения продолжительности d,
	// а затем отправляет текущее время по возвращаемому каналу.
}

func main() {
	sleep(3 * time.Second);
	fmt.Println("3 seconds after")
}