package kvserver

import (
	"fmt"
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
	b := make([]byte, 4)
	rb := 0
	wb := 0
	for {
		n, err := conn.Read(b)
		rb += n
		if err != nil {
			fmt.Println("failed in read", "read", rb, "wrote", wb, err)
			break
		}
		n, err = conn.Write(b[:n])
		wb += n
		if err != nil {
			fmt.Println("failed in write", "read", rb, "wrote", wb, err)
			break
		}
	}
	conn.Close()
}
