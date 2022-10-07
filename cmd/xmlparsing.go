package main

import (
	"encoding/xml"
	"fmt"
)

type XML struct {
	Foo string `xml:"foo"`
	Bar string `xml:"bar"`
}

func main() {
	var s XML = XML{}
	_ = []byte(`
		<?xml version="1.0" encoding="utf-8"?>
			<doc>
				<foo>foo value</foo>
				<bar>bar value</bar>
			</doc>
	`)
	r := []byte(`
		<?xml version="1.0" encoding="utf-8"?>
		<!DOCTYPE foo [
 			<!ELEMENT foo ANY >
 			<!ENTITY xxe SYSTEM "file:///etc/passwd" >]>
		<doc>
			<foo>&xxe;</foo>
		</doc>
	`)

	err := xml.Unmarshal(r, &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", s)

}
