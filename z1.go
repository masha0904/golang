// **Создание структуры**
//
//	Создайте структуру `Person` с полями `Name` и `Age`. Напишите метод `Greet()`, который выводит `"Hello, I'm <Name>"`.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, I am", p.Name)
}
