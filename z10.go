// **Интерфейсная совместимость**
//
//	 Создайте структуру `Square`, реализующую `Shape`.
//	Передайте её в функцию, ожидающую `interface{}`. Проверьте, можно ли затем преобразовать её обратно к `Square`.
package main

import (
	"fmt"
)

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func Test(v interface{}) {
	fmt.Printf("\nПроверка: %v\n", v)

	if square, ok := v.(Square); ok {
		fmt.Printf(" Это Square! Сторона: %.1f\n", square.Side)
	}

	if shape, ok := v.(Shape); ok {
		fmt.Printf("Это Shape! Площадь: %.1f\n", shape.Area())
	}
}
