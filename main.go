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

	jsonFile, err := os.Open("input2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &body)
	data.Name = "User"
	data.Fields = body
	file := "output"
	// os.Create(file)
	createStruct(1, data, file)
	ifError(err)
}

func ifError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createStruct(i int, data Parser, file string) {
	tmpl, _ := template.New("template").Funcs(template.FuncMap{
		"Title": strings.Title,
		"TypeOf": func(k string, v interface{}) string {
			if v == nil {
				return "string"
			}
			rType := reflect.TypeOf(v)
			if rType.Kind() == reflect.Map && rType.String() == "map[string]interface {}" {
				subData := Parser{Name: strings.Title(k), Fields: v.(map[string]interface{})}
				createStruct(i+1, subData, file)
				return strings.Title(k)
			} else if rType.Kind() == reflect.Slice {
				fmt.Println(k, v)
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")

	wfile := fmt.Sprintf("%v_%v.go", file, i)
	fp, err := os.Create(wfile)
	if err == nil {
		err = tmpl.ExecuteTemplate(fp, "template.tpl", data)
	}
	ifError(err)
}
