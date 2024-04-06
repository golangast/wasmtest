package main

import (
	"github.com/golangast/wasmtest/dom"
)

func main() {
	dom.Hide("loading")
	<-make(chan bool)
}
