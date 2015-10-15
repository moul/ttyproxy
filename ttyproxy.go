package ttproxy

type TCPServer struct{}

func NewTCPServer() TCPServer {
	return TCPServer{}
}

func (s *TCPServer) Proxy(tty TTY, bind string) error {
	return nil
}

type TTY struct {
	Path string
}

func NewTTY(path string) TTY {
	return TTY{
		Path: path,
	}
}
