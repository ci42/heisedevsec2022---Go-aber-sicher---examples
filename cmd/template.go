package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	var out bytes.Buffer
	t, err := template.New("foo").Parse(`Hello, {{.}}!`)
	if err != nil {
		log.Fatalln("Could not parse template.")
	}
	err = t.Execute(&out, "<script>alert('you have been pwned')</script>")
	fmt.Println(out.String())
}
