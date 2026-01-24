// **Мокирование через интерфейсы**
//
//	Напишите интерфейс `Logger` с методом `Log(msg string)`.
//
// Создайте реализацию `StdoutLogger` и мок `MockLogger` для тестов. Напишите простой юнит-тест.
package main

import "testing"

// 1. Интерфейс
type Logger interface {
	Log(msg string)
}

// 2. Реальная реализация
type RealLogger struct{}

func (r RealLogger) Log(msg string) {
	// в реальности пишет в лог
}

// 3. Мок
type MockLogger struct {
	called bool
}

func (m *MockLogger) Log(string) {
	m.called = true
}

// 4. Сервис
type Service struct{ logger Logger }

func (s Service) Do() {
	s.logger.Log("работаю")
}

// 5. Тест
func TestService(t *testing.T) {
	mock := &MockLogger{}
	service := Service{logger: mock}

	service.Do()

	if !mock.called {
		t.Error("логер не был вызван")
	}
}
