package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	if err := ShellBash("cd view/js &&  GOOS=js GOARCH=wasm go build -o main.wasm"); err != nil {
		if err != nil {
			fmt.Println(err)
		}
	}
	e.Static("/", "view")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = renderer

	// Named route "foobar"
	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "wasm_exec.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":8000"))
}

func ShellBash(s string) error {

	out, errout, err := Shellout(s)
	if err != nil {
		return err
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println(errout)
	}

	return nil

}
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
