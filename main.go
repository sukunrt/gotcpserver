package main

import "kvserver/kvserver"

func main() {
	server := kvserver.Server{Host: "127.0.0.1", Port: "8888"}
	server.Start()
}
