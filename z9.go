//  9. **Пустой интерфейс и type assertion**
//     Напишите функцию `Describe(v interface{})`, которая определяет тип переданного значения
//     (`int`, `string`, `Person`) и выводит соответствующее сообщение с использованием type switch.

package main

import "fmt"

func Describe(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case Person:
		fmt.Println("Person:", val.Name, val.Age)
	default:
		fmt.Println("неизвестный тип")
	}
}
