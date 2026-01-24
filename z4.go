// **Эмбеддинг как замена наследованию**
//
//	 Создайте структуру `Animal` с методом `Speak()`.
//	Создайте структуру `Dog`, эмбеддящую `Animal`, и переопределите `Speak()` так, чтобы он выводил `"Woof!"`.
package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Speak() {
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}
