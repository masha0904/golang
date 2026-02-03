// ### Задача 5
// **Условие:**
// Создайте небуферизованный канал типа `int`.
// В горутине отправьте в него число `42`.
// В `main` прочитайте это число и выведите: `"Received: 42"`.
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch

	fmt.Printf("Received: %d\n", value)
}
