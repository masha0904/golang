// ### Задача 9
// **Условие:**
// Горутина каждые 300 мс выводит `"Tick"`.
// Остановите её через 1.5 секунды с помощью канала сигнала завершения (`done := make(chan bool)`).
// После остановки выведите `"Stopped"`.
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Tick")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	done <- true
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Stopped")
}
