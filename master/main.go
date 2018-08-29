package main

import (
	_ "runtime/pprof"
)

const (
	ADDR_1 = "localhost:9003"
)

func main() {
	go getMessage()
	go Server()
	Client()
}
