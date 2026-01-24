// **Базовое использование интерфейса**
//
//	Определите интерфейс `Shape` с методом `Area() float64`. Реализуйте его для структур `Rectangle` и `Circle`.
package main

import "math"

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Height float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Height
}
