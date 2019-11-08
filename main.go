package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//Parser model
type Parser struct {
	Name   string
	Fields map[string]interface{}
}

func main() {

	var data Parser

	data.Name = "User"
	data.Fields = map[string]interface{}{
		"hello": "World",
		"world": 1,
	}

	tmpl, _ := template.New("template").Funcs(template.FuncMap{
		"Title": strings.Title,
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")

	fp, err := os.Create("output.go")
	if err == nil {
		err = tmpl.ExecuteTemplate(fp, "template.tpl", data)
	}
	ifError(err)
}

func ifError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
