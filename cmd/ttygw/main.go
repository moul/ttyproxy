package main

import "time"

func main() {
	server := ttyproxy.NewTCPServer()
	tty := ttyproxy.NewTTY("/dev/tty")
	server.Proxy(tty, ":2001")
	for {
		time.Sleep(time.Second * 10)
	}
}
