// **Композиция вместо наследования**
//
//	 Создайте структуру `Engine` с методом `Start()`.
//	Создайте структуру `Car`, содержащую `Engine` как поле. Реализуйте метод `Drive()`, использующий `engine.Start()`.
package main

import "fmt"

type Engine struct {
	name string
}

func (e Engine) Start() {
	fmt.Println(e.name)
}

type Car struct {
	Engine Engine
}

func (c Car) Drive() {
	c.Engine.Start()
}
