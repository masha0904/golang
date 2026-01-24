// **Фабрика объектов**
//
//	Напишите функцию `NewPerson(name string, age int) *Person`, инициализирующую и возвращающую указатель на `Person`.
package main

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
