package coding_standard

// reference: https://coolshell.cn/articles/21146.html

type Server struct {
	Name string
}

// 方式1 Function Option

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

// 方式2

type Config struct {
	Name string
}

func NewServer2(config *Config) *Server {
	return &Server{Name: config.Name}
}

// 方式3

func (s *Server) WithName(name string) *Server {
	s.Name = name
	return s
}
