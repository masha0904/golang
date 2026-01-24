// **Изменение полей через интерфейс**
//
//	Можно ли изменить внутреннее состояние объекта, если у вас только интерфейсный тип? Продемонстрируйте на примере.
package main

type BookManager interface {
	GetTitle() string
	GetAuthor() string
	UpdateTitle(newTitle string)
	UpdateAuthor(newAuthor string)
}

type Book struct {
	title  string
	author string
}

func (b Book) GetTitle() string {
	return b.title
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b *Book) UpdateTitle(newTitle string) {
	b.title = newTitle
}

func (b *Book) UpdateAuthor(newAuthor string) {
	b.author = newAuthor
}
