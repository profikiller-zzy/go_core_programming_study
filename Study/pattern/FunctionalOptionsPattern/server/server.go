package server

// Server 假设有一个对象 Server，有很多配置项：
type Server struct {
	Host string
	Port int
	TLS  bool
}

type ServerOption func(*Server)

// WithHost 设置主机名
func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

// WithPort 设置端口
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

// WithTLS 是否启用 TLS
func WithTLS(enabled bool) ServerOption {
	return func(s *Server) {
		s.TLS = enabled
	}
}

// NewServer 构造函数，传参是函数
func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		Host: "localhost",
		Port: 80,
		TLS:  false,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
