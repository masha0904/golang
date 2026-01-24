// **Конструктор с опциями (Functional Options)**
//
//	Реализуйте гибкий конструктор для структуры `Server` с полями `Port`, `Host`, `TLS`. Используйте паттерн функциональных опций.
package main

type Server struct {
	Host string
	Port int
	TLS  bool
}

type Option func(*Server)

func TheHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func ThePort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func TheTLS(tls bool) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(options ...Option) *Server {
	s := &Server{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}
	for _, option := range options {
		option(s)
	}
	return s
}
