package dom

import (
	"syscall/js"
)

var Document js.Value

func init() {
	Document = js.Global().Get("document")

}
func getElementById(elem string) js.Value {
	return Document.Call("getElementById", elem)
}
func GetElementValue(elem string, value string) js.Value {
	return getElementById(elem).Get(value)
}

func Hide(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "none")
}

func Show(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "block")
}
func AddClass(elem string, class string) {
	GetElementValue(elem, "classList").Call("add", class)
}

func RemoveClass(elem string, class string) {
	classList := GetElementValue(elem, "classList")
	if classList.Call("contains", class).Bool() {
		classList.Call("remove", class)
	}
}
func SetInner(elem string, value string) {
	getElementById(elem).Set("innerText", value)
}
