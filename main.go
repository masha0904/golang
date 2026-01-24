package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//task1
	//a := Person{
	//	Name: "Maria",
	//	Age:  20,
	//}
	//a.Greet()

	//task2
	//a := Person{
	//	Age: 20,
	//}
	//a.Birthday2()
	//fmt.Println(a.Age)

	//task3
	//a := bankAccount{balance: 100}
	//amount := 500.0
	//
	//fmt.Println("Начальный баланс:", a.GetBalance())
	//a.Deposit(amount)
	//fmt.Println("Конечный баланс:", a.GetBalance())

	//task4
	//d := Dog{
	//	Animal: Animal{name: "Rex"},
	//}
	//d.Speak()

	//task5
	//car := Car{Engine: Engine{"DA"}}
	//car.Drive()

	//task6
	//food := Food{Vegetable{name: "Potato"}, Fruit{"Banana"}}
	//food.Vegetable.Message()

	//task7
	//rect := Rectangle{Height: 14}
	//circle := Circle{Radius: 5}
	//fmt.Println("Площадь круга:", circle.Area())
	//fmt.Println("Площадь прямоугольника:", rect.Area())

	//task8
	//circle := Circle{Radius: 5}
	//rect := Rectangle{Height: 5}
	//
	//PrintArea(circle)
	//PrintArea(rect)

	//task9
	//Describe(100)
	//Describe("text")
	//Describe(Person{Name: "Maria", Age: 20})

	//task10
	//square := Square{Side: 5.0}
	//
	//fmt.Println("\n1.  Square как interface{}:")
	//Test(square)
	//
	//fmt.Println("\n2.  Square как Shape:")
	//var shape Shape = square
	//Test(shape)

	//task11
	//person1 := NewPerson("Alice", 25)
	//person2 := NewPerson("Bob", 30)
	//
	//fmt.Printf("%s: %d лет\n", person1.Name, person1.Age)
	//fmt.Printf("%s: %d лет\n", person2.Name, person2.Age)

	//task12
	//s := NewServer(ThePort(500))
	//fmt.Println(s)

	//task13
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	workers := []Worker{
		&FileProcessor{FileName: "data.txt"},
		&NetworkFetcher{URL: "https://example.com"},
	}

	for i, w := range workers {
		fmt.Printf("\nWorker %d:\n", i+1)
		if err := w.Do(ctx); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}

	//task14
	//service1 := Service{logger: RealLogger{}}
	//service1.Do()
	//t := &testing.T{}
	//TestService(t)

	//task15
	//var l *List = nil
	//fmt.Println(l.Len())

	//task16
	//book := &Book{
	//	title:  "Война и мир",
	//	author: "Лев Толстой",
	//}
	//// Работаем через интерфейс
	//var manager BookManager = book
	//
	//fmt.Println("ДО изменения:")
	//fmt.Println("Название:", manager.GetTitle())
	//fmt.Println("Автор:", manager.GetAuthor())
	//
	//manager.UpdateTitle("Анна Каренина")
	//manager.UpdateAuthor("Толстой Л.Н.")
	//
	//fmt.Println("\nПОСЛЕ изменения:")
	//fmt.Println("Название:", manager.GetTitle())
	//fmt.Println("Автор:", manager.GetAuthor())

	//task17
	//ints := []int{10, 20, 30}
	//ifaces := IntsToInterfaces(ints)
	//fmt.Println(ifaces)

	//task18
	//f := &File{name: "test"}
	//var rc ReadCloser = f
	//
	//rc.Read()
	//rc.Close()
}
