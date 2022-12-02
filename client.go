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

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	st := time.Now()
	b := make([]byte, 100)
	nn := len(b)
	n, err := conn.Write(b)
	fmt.Println("write", n, err)
	if err != nil {
		panic(err)
	}
	n, err = conn.Read(b)
	fmt.Println(n, err)
	d := time.Since(st)
	conn.Close()
	time.Sleep(1)
	fmt.Printf("Wrote %d bytes at %s\n", n*nn, formatSpeed(n*nn, d))
}
