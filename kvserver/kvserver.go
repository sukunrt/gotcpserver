package kvserver

import (
	"io"
	"net"
)

type Server struct {
	Host string
	Port string
}

func (s *Server) Start() {
	ss, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ss.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	conn.(*net.TCPConn).SetNoDelay(false)
	b := make([]byte, 10000)
	for {
		n, err := conn.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		_, err = conn.Write(b[:n])
		if err != nil {
			panic(err)
		}
	}
	conn.Close()
}
