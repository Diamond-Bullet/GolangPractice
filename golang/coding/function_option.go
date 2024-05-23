package coding

// reference: https://coolshell.cn/articles/21146.html

type Server struct {
	Name string
}

// Method 1: Function Option

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

// Method 2

type Config struct {
	Name string
}

func NewServer2(config *Config) *Server {
	return &Server{Name: config.Name}
}

// Method 3

func (s *Server) WithName(name string) *Server {
	s.Name = name
	return s
}
