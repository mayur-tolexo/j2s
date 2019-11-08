package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	var (
		data Parser
		body map[string]interface{}
	)

	jsonFile, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &body)
	data.Name = "User"
	data.Fields = body

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
