// **Методы на nil-приёмниках**
// Создайте структуру `List` с методом `Len() int`, возвращающим длину.
// Реализуйте его так, чтобы он корректно работал даже если приёмник — `nil`. Продемонстрируйте это.
package main

type List struct {
	len int
}

func (l *List) Len() int {
	if l == nil {
		return 0
	}
	return 1
}
