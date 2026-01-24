// **Интерфейс ≠ конкретный тип**
//
//	Объясните, почему `[]T` не реализует `[]interface{}`. Напишите функцию, преобразующую `[]int` в `[]interface{}`.
package main

func IntsToInterfaces(nums []int) []interface{} {
	result := make([]interface{}, len(nums))
	for i, n := range nums {
		result[i] = n
	}
	return result
}
