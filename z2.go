// Задача 2
// Условие:
// Напишите функцию `greet(name string)`, которая выводит `"Hello, {name}!"`.
// Запустите эту функцию как горутину с именем `"Alice"`.
// Убедитесь, что сообщение выводится.
package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	name := "Alice"
	go greet(name)
	time.Sleep(100 * time.Millisecond)
}
