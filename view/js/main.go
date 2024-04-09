package main

import (
	"syscall/js"
)

func main() {
	temp := `<p>dfsfdfdf</p>`
	tempform := `  <form action="/" method="post">
	<label for="fname">First name:</label><br>
	<input type="text" id="fname" name="fname" value="John"><br>
	<label for="lname">Last name:</label><br>
	<input type="text" id="lname" name="lname" value="Doe"><br><br>
	<input type="submit" value="Submit">
  </form> `

	js.Global().Call("CreateComponent", "my-component", temp)
	js.Global().Call("CreateFormComponent", "form-component", tempform)

	<-make(chan bool)
}
