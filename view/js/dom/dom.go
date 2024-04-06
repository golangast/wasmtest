package dom

import (
	"syscall/js"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}
func getElementById(elem string) js.Value {
	return document.Call("getElementById", elem)
}
func getElementValue(elem string, value string) js.Value {
	return getElementById(elem).Get(value)
}

func Hide(elem string) {
	getElementValue(elem, "style").Call("setProperty", "display", "none")
}

func Show(elem string) {
	getElementValue(elem, "style").Call("setProperty", "display", "block")
}
func AddClass(elem string, class string) {
	getElementValue(elem, "classList").Call("add", class)
}

func RemoveClass(elem string, class string) {
	classList := getElementValue(elem, "classList")
	if classList.Call("contains", class).Bool() {
		classList.Call("remove", class)
	}
}
func SetInner(elem string, value string) {
	getElementById(elem).Get(value).Set("innerText", value)
}
