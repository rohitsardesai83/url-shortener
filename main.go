package main

import (
	"runtime"
	"url-shortener/pkg/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server.New()

}
