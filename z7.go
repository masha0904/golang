// ### Задача 7
// **Условие:**
// Создайте буферизованный канал на 2 значения типа `string`.
// Запустите горутину, которая отправляет туда `"first"` и `"second"` **без блокировки**.
// В `main` прочитайте оба значения и выведите их.
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "first"
		ch <- "second"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
