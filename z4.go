//###  Задача 4
//**Условие:**
//Перепишите **Задачу 3**, но вместо `time.Sleep` используйте `sync.WaitGroup`, чтобы дождаться завершения всех трёх горутин.
//Программа должна корректно завершаться **без задержек**.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3")
	}()
	wg.Wait()
}
