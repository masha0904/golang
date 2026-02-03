// Задача 8
// **Условие:**
// Напишите программу, в которой горутина выводит числа от 1 до 5 с паузой 200 мс между ними.
// Используйте `sync.WaitGroup`, чтобы `main` дождался завершения.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for num := range ch {
			fmt.Println(num)
		}
	}()

	wg.Wait()
}
