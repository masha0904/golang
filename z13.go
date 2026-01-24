// **Интерфейс с контекстом**
//
//	Создайте интерфейс `Worker` с методом `Do(ctx context.Context) error`. Реализуйте его для `FileProcessor` и `NetworkFetcher`.
package main

import (
	"context"
	"fmt"
	"time"
)

type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct {
	FileName string
}

func (f *FileProcessor) Do(ctx context.Context) error {
	fmt.Printf("Обработка файла: %s\n", f.FileName)

	// Имитация работы с файлом
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("Файл %s обработан\n", f.FileName)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("обработка файла прервана: %v", ctx.Err())
	}
}

type NetworkFetcher struct {
	URL string
}

func (n *NetworkFetcher) Do(ctx context.Context) error {
	fmt.Printf("Загрузка с: %s\n", n.URL)

	// Имитация сетевого запроса
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("Данные с %s загружены\n", n.URL)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("загрузка прервана: %v", ctx.Err())
	}
}
