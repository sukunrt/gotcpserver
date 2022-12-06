package main

import (
	"flag"
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
	for sent < total {
		conn.SetWriteDeadline(time.Now().Add(time.Duration(100 * time.Millisecond)))
		n, err := conn.Write(b)
		if n != len(b) {
			fmt.Printf("Could only write: %d bytes\n", n)
		}
		if err != nil && !err.(net.Error).Timeout() {
			panic(err)
		}
		sent += n
	}
}

func main() {
	total := flag.Int("total", 1000, "total bytes to be sent over the network")
	batchSize := flag.Int("batch", 100, "batch size of a single transaction")
	flag.Parse()
	c, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	conn := c.(*net.TCPConn)
	conn.SetNoDelay(false)
	st := time.Now()
	sendBytes(conn, *total, *batchSize)
	d := time.Since(st)
	conn.Close()
	fmt.Printf("Wrote %d bytes at %s\n", *total, formatSpeed(*total, d))
}
