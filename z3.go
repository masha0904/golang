// **Приватность полей**
// Создайте структуру `bankAccount` с приватным полем `balance`. Напишите публичные методы `Deposit(amount float64)`
// и `GetBalance() float64`. Почему нельзя напрямую изменить `balance` извне пакета?
package main

import "fmt"

type bankAccount struct {
	balance float64
}

func (a *bankAccount) Deposit(amount float64) {
	a.balance += amount
	fmt.Println("Внесено:", amount)
}
func (a *bankAccount) GetBalance() float64 {
	return a.balance
}
