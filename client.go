package main

import (
	"fmt"
	"net"
	"time"
)

func formatSpeed(numBytes int, d time.Duration) string {
	speed := float64(numBytes) / float64(d.Seconds())
	unit := "B/s"
	if speed > 1000000.0 {
		speed /= 1000000.0
		unit = "MB/s"
	} else if speed > 1000.0 {
		speed /= 1000.0
		unit = "KB/s"
	}
	return fmt.Sprintf("%0.2f %s", speed, unit)
}

func sendBytes(conn *net.TCPConn, total int, batchSize int) {
	b := make([]byte, batchSize)
	sent := 0
	received := 0
	for sent < total {
		conn.SetWriteDeadline(time.Now().Add(time.Duration(100 * time.Millisecond)))
		n, err := conn.Write(b)
		if err != nil && !err.(net.Error).Timeout() {
			panic(err)
		}
		sent += n
		n, err = conn.Read(b)
		if err != nil {
			panic(err)
		}
		received += n
	}
	for received < sent {
		n, err := conn.Read(b)
		if err != nil {
			panic(err)
		}
		received += n
	}
}

func main() {
	total := 10000000
	batchSize := 10000
	c, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	conn := c.(*net.TCPConn)
	conn.SetNoDelay(false)
	st := time.Now()
	sendBytes(conn, total, batchSize)
	d := time.Since(st)
	conn.Close()
	fmt.Printf("Wrote %d bytes at %s\n", total, formatSpeed(total, d))
}
