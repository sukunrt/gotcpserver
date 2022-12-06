package kvserver

import (
	"io"
	"net"
)

type Server struct {
	Host      string
	Port      string
	BatchSize int
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
		c := conn.(*net.TCPConn)
		ch := make(chan []byte, 100000)
		go s.readConnection(c, ch)
	}
}

func (s *Server) readConnection(conn *net.TCPConn, ch chan []byte) {
	conn.SetReadBuffer(10 * 1000 * 1000)
	b := make([]byte, s.BatchSize)
	totalRead := 0
	for {
		n, err := conn.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		totalRead += n
		if totalRead >= 1000*1000*10 {
			break
		}
	}
	conn.Close()
}

func (s *Server) writeConnection(conn *net.TCPConn, ch chan []byte) {
	conn.SetWriteBuffer(10 * 1000 * 1000)
	conn.SetNoDelay(false)
	for {
		b := <-ch
		if len(b) == 0 {
			break
		}
		_, err := conn.Write(b)
		if err != nil {
			panic(err)
		}
	}
}
