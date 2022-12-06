package main

import (
	"flag"
	"kvserver/kvserver"
)

func main() {
	batchSize := flag.Int("batch", 10, "batch size for the echo server")
	flag.Parse()
	server := kvserver.Server{Host: "localhost", Port: "8888", BatchSize: *batchSize}
	server.Start()
}
