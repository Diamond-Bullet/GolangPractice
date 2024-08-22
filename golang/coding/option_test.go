package coding

import "testing"

// Move optional configurations out of the constructor.
// reference: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

func TestInitializeObject(t *testing.T) {
	// Method 1: Functional Option
	_ = NewServer1(Name("server1"))

	// Method 2: separate Config struct
	_ = NewServer2(&Config{Name: "server2"})

	// Method 3: Builder Pattern
	_ = NewServer3().WithName("server3")
}

type Server struct {
	Name string
}

type Option func(*Server)

func Name(name string) Option {
	return func(server *Server) {
		server.Name = name
	}
}

func NewServer1(options ...Option) *Server {
	server := &Server{}
	for _, opt := range options {
		opt(server)
	}
	return server
}

type Config struct {
	Name string
}

func NewServer2(config *Config) *Server {
	return &Server{Name: config.Name}
}

func (s *Server) WithName(name string) *Server {
	s.Name = name
	return s
}

func NewServer3() *Server {
	return &Server{}
}
