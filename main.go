package main

import (
	"text/template"
	"fmt"
	"flag"
	"bytes"
	"github.com/atotto/clipboard"
)

//
// Rubygen: generate ruby html tags.
//

type RubyChar struct {
	Symbol string
	Text string
}

type RubyElem struct {
	Chars []RubyChar
}


func main() {
	rubyTemplate, err := template.ParseFiles("ruby.tmpl")

	if err != nil {
		fmt.Println(err)
		return
	}

	flag.Parse()

	elem := RubyElem{}
	elem.Chars = make([]RubyChar,0)

	args := flag.Args()

	//Loop through arguments (in pairs)
	for index := 0; index < len(args); index += 2 {
		elem.Chars = append(elem.Chars, RubyChar{Symbol: args[index], Text: args[index + 1]})
	}

	output := &bytes.Buffer{}

	rubyTemplate.Execute(output, elem)

	clipboard.WriteAll(output.String())
	fmt.Print(output.String())
}
