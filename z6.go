// ###  Задача 6
// **Условие:**
// Используйте канал **без передачи данных** (только как сигнал), чтобы:
// - Горутина выполнила свою работу (вывела `"Worker done"`),
// - `main` дождался этого и вывел `"Main continues"`.
package main

import "fmt"

func main() {
	a := make(chan struct{})
	go func() {
		fmt.Println("Worker done")
		a <- struct{}{}
	}()
	<-a
	fmt.Println("Main continues")
}
