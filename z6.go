// **Конфликт имён при эмбеддинге**
//
//	Эмбеддьте две структуры с методами одинакового имени в третью. Что произойдёт? Как вызвать нужный метод явно?
//
// Если две встроенные структуры имеют методы с одинаковыми именами:
//
// # Компилятор не знает, какой метод вызывать
//
// Нужно указывать явно через имя типа
package main

import "fmt"

type Vegetable struct {
	name string
}
type Fruit struct {
	name string
}
type Food struct {
	Vegetable
	Fruit
}

func (v Vegetable) Message() {
	fmt.Println(v.name)
}

func (v Fruit) Message() {
	fmt.Println(v.name)
}
