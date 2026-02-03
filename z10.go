// ### Задача 10
// **Условие:**
// Создайте программу, которая:
// - Запускает 2 горутины.
// - Первая отправляет в канал числа 1, 2, 3.
// - Вторая читает из этого канала и выводит их.
// - Используйте `sync.WaitGroup` или канал для синхронизации.
// - Программа завершается корректно, без зависаний.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		ch <- 1
		ch <- 2
		ch <- 3
	}()
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	wg.Wait()
}
