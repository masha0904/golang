// ###  Задача 1
// **Условие:**
// Напишите программу, которая:
// - Выводит `"Main started"`.
// - Запускает горутину, которая выводит `"Goroutine finished"`.
// - Выводит `"Main finished"`.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main started")
	go func() {
		fmt.Println("Goroutine finished")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main finished")

}
