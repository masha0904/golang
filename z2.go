// **Указатели vs значения**
// Добавьте метод `Birthday()` к `Person`, увеличивающий возраст на 1. Реализуйте его дважды:
// один раз с приёмником-значением, другой — с приёмником-указателем. Объясните разницу в поведении.
package main

import "fmt"

func (p Person) Birthday() {
	p.Age += 1
	fmt.Println(p.Age)
}
func (p *Person) Birthday2() {
	p.Age += 1
	fmt.Println(p.Age)
}
