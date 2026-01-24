// **Интерфейс как параметр функции**
//
//	Напишите функцию `PrintArea(s Shape)`, которая выводит площадь любой фигуры.
package main

import (
	"fmt"
)

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}
