package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
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
	file := "output.go"

	strMap := make(map[int]string)
	createStruct(data, 1, strMap)
	// fmt.Println(strMap)

	fp, err := os.Create(file)
	if err == nil {
		fp.WriteString("package jsonToStruct\n")
		for _, v := range strMap {
			fp.WriteString(v)
			fp.WriteString("\n")
		}
	}
	ifError(err)
}

func ifError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createStruct(data Parser, i int, strMap map[int]string) {

	tmpl, _ := template.New("template").Funcs(template.FuncMap{
		"Title": getFieldName,
		"TypeOf": func(k string, v interface{}) string {
			if v == nil {
				return "string"
			}
			rType := reflect.TypeOf(v)
			if rType.Kind() == reflect.Map && rType.String() == "map[string]interface {}" {
				subData := Parser{Name: strings.Title(k), Fields: v.(map[string]interface{})}
				createStruct(subData, i+1, strMap)
				return getFieldName(k)
			} else if rType.Kind() == reflect.Slice {
				// fmt.Println(k, v)
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")

	var buf bytes.Buffer
	tmpl.ExecuteTemplate(&buf, "template.tpl", data)
	strMap[i] = buf.String()
	return
}

//getFieldName will return field name in Camel Case
func getFieldName(k string) (f string) {
	f = strcase.ToCamel(k)
	r := strings.NewReplacer(
		"Id", "ID",
	)
	f = r.Replace(f)
	return
}
