package main

import (
	"syscall/js"
)

func main() {
	temp := `<p>dfsfdfdf</p>`

	js.Global().Call("CreateComponent", "my-component", temp)

	<-make(chan bool)
}
