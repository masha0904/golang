// **"Наследование" через интерфейсы**
//
//	 Определите интерфейс `Reader` и `Closer`.
//	Создайте интерфейс `ReadCloser`, включающий оба. Реализуйте его для структуры `File`
package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type ReadCloser interface {
	Reader
	Closer
}

type File struct {
	name string
}

func (f *File) Read() {
	fmt.Println(f.name)
}

func (f *File) Close() {
	fmt.Println("Файл закрыт")
}
