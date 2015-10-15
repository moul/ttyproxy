package main

import (
	"time"

	"github.com/moul/ttygw"
)

func main() {
	server := ttygw.NewTCPServer()
	tty := ttygw.NewTTY("/dev/tty")
	server.Proxy(tty, ":2001")
	for {
		time.Sleep(time.Second * 10)
	}
}
