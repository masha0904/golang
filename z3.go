// ###  Задача 3
// **Условие:**
// Запустите **3 горутины**. Каждая должна вывести своё число:
// - Первая: `"Goroutine 1"`
// - Вторая: `"Goroutine 2"`
// - Третья: `"Goroutine 3"`
// Используйте `time.Sleep` в `main`, чтобы все сообщения успели напечататься.
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Goroutine 1")
	}()
	go func() {
		fmt.Println("Goroutine 2")
	}()
	go func() {
		fmt.Println("Goroutine 3")
	}()
	time.Sleep(100 * time.Millisecond)
}
